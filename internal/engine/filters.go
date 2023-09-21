package engine

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"regexp"
	"slices"
	"sort"

	"github.com/apache/arrow/go/v14/arrow"
	"github.com/apache/arrow/go/v14/arrow/array"
	"github.com/apache/arrow/go/v14/arrow/compute"
	"github.com/apache/arrow/go/v14/arrow/scalar"
	"github.com/bits-and-blooms/bitset"
	blocksv1 "github.com/vinceanalytics/vince/gen/proto/go/vince/blocks/v1"
	v1 "github.com/vinceanalytics/vince/gen/proto/go/vince/store/v1"
	"github.com/vinceanalytics/vince/internal/entry"
)

type Op uint

const (
	Eq Op = 1 + iota
	Ne
	Gt
	GtEg
	Lt
	LtEq
	ReEq
	ReNe
)

func (o Op) String() string {
	switch o {
	case Eq:
		return "equal"
	case Ne:
		return "not_equal"
	case Gt:
		return "greater"
	case GtEg:
		return "greater_equal"
	case Lt:
		return "less"
	case LtEq:
		return "less_equal"
	case ReEq:
		return "regex_equal"
	case ReNe:
		return "regex_not_equal"
	default:
		panic(fmt.Sprintf("unknown operation %d", o))
	}
}

var ErrNoFilter = errors.New("no filters")

type Filters struct {
	Index []IndexFilter
	Value []ValueFilter
}

// IndexFilter is an interface for choosing row groups and row group pages to
// read based on a column index.
//
// This filter is first applied before values are read into arrow.Record. Only
// row groups and row group pages that matches across all IndexFilter are
// selected for reads
type IndexFilter interface {
	Column() v1.Column
	FilterIndex(ctx context.Context, idx *blocksv1.ColumnIndex) (*RowGroups, error)
}

type RowGroups struct {
	groups bitset.BitSet
	pages  map[uint]*bitset.BitSet
}

func (g *RowGroups) Set(index uint, pages []uint) {
	g.groups.Set(index)
	if g.pages == nil {
		g.pages = make(map[uint]*bitset.BitSet)
	}
	p := new(bitset.BitSet)
	for i := range pages {
		p.Set(pages[i])
	}
}

func (g *RowGroups) SelectGroup(groups *bitset.BitSet, f func(uint, *bitset.BitSet)) {
	for k, v := range g.pages {
		if groups.Test(k) {
			f(k, v)
		}
	}
}

type FilterIndex func(ctx context.Context, idx *blocksv1.ColumnIndex) (*RowGroups, error)

type ValueFilter interface {
	Column() v1.Column
	FilterValue(ctx context.Context, value arrow.Array) (arrow.Array, error)
}

type FilterValue func(ctx context.Context, b *array.BooleanBuilder, value arrow.Array) (arrow.Array, error)

type IndexFilterResult struct {
	RowGroups []uint
	Pages     []*bitset.BitSet
}

func BuildIndexFilter(
	ctx context.Context,
	f []IndexFilter, source func(v1.Column) *blocksv1.ColumnIndex) (*IndexFilterResult, error) {
	if len(f) == 0 {
		return nil, ErrNoFilter
	}
	// Make sure timestamp is processed first
	sort.Slice(f, func(i, j int) bool {
		return f[i].Column() == v1.Column_timestamp
	})
	groups := make([]*RowGroups, len(f))
	for i := range f {
		g, err := f[i].FilterIndex(ctx, source(f[i].Column()))
		if err != nil {
			return nil, err
		}
		groups[i] = g
	}
	g := &groups[0].groups
	for i := range groups {
		if i == 0 {
			continue
		}
		g = groups[i].groups.Intersection(g)
	}
	o := make([]uint, g.Count())
	_, all := g.NextSetMany(0, o)
	slices.Sort(o)
	r := &IndexFilterResult{
		RowGroups: make([]uint, len(all)),
		Pages:     make([]*bitset.BitSet, len(all)),
	}
	pages := make(map[uint]*bitset.BitSet)
	for _, a := range all {
		pages[a] = new(bitset.BitSet)
	}
	for i := range groups {
		groups[i].SelectGroup(g, func(u uint, bs *bitset.BitSet) {
			if i == 0 {
				pages[u] = bs
			} else {
				pages[u] = pages[u].Intersection(bs)
			}
		})
	}
	for i := range all {
		r.RowGroups[i] = all[i]
		r.Pages[i] = pages[all[i]]
	}
	return r, nil
}

func BuildValueFilter(ctx context.Context,
	f []ValueFilter,
	source func(v1.Column) arrow.Array) (arrow.Array, error) {
	// Make sure timestamp is processed first
	sort.Slice(f, func(i, j int) bool {
		return f[i].Column() == v1.Column_timestamp
	})
	var filter arrow.Array
	for i := range f {
		g, err := f[i].FilterValue(ctx, source(f[i].Column()))
		if err != nil {
			return nil, err
		}
		if i == 0 {
			filter = g
		} else {
			// we and all filters
			filter, err = call("and", nil, filter, g, filter.Release, g.Release)
			if err != nil {
				return nil, err
			}
		}
	}

	return filter, nil
}

func call(name string, o compute.FunctionOptions, a any, b any, fn ...func()) (arrow.Array, error) {
	ad := compute.NewDatum(a)
	bd := compute.NewDatum(b)
	defer ad.Release()
	defer bd.Release()
	defer func() {
		for _, f := range fn {
			f()
		}
	}()
	out, err := compute.CallFunction(context.TODO(), name, o, ad, bd)
	if err != nil {
		return nil, err
	}
	return out.(*compute.ArrayDatum).MakeArray(), nil
}

type IndexMatchFuncs struct {
	Col             v1.Column
	FilterIndexFunc func(ctx context.Context, idx *blocksv1.ColumnIndex) (*RowGroups, error)
}

var _ IndexFilter = (*IndexMatchFuncs)(nil)

func (i *IndexMatchFuncs) Column() v1.Column {
	return i.Col
}

func (i *IndexMatchFuncs) FilterIndex(ctx context.Context, idx *blocksv1.ColumnIndex) (*RowGroups, error) {
	return i.FilterIndexFunc(ctx, idx)
}

type ValueMatchFuncs struct {
	Col             v1.Column
	FilterValueFunc func(ctx context.Context, value arrow.Array) (arrow.Array, error)
}

var _ ValueFilter = (*ValueMatchFuncs)(nil)

func (i *ValueMatchFuncs) Column() v1.Column {
	return i.Col
}

func (i *ValueMatchFuncs) FilterValue(ctx context.Context, value arrow.Array) (arrow.Array, error) {
	return i.FilterValueFunc(ctx, value)
}

type Value interface {
	~int64 | // support arrow.Timestamp
		float64 | string
}

func Match[T Value](col v1.Column, matchValue T, op Op) ValueFilter {
	return &ValueMatchFuncs{
		Col: col,
		FilterValueFunc: func(ctx context.Context, value arrow.Array) (arrow.Array, error) {
			switch op {
			case ReEq:
				m, ok := any(matchValue).(string)
				if !ok {
					slog.Warn("using regex for not string columns is not supported",
						"column", col.String(),
						"value", any(matchValue),
					)
					return nil, ErrNoFilter
				}
				return boolExpr(value.(*array.String), reMatch(m))
			case ReNe:
				m, ok := any(matchValue).(string)
				if !ok {
					slog.Warn("using regex for not string columns is not supported",
						"column", col.String(),
						"value", any(matchValue),
					)
					return nil, ErrNoFilter
				}
				return boolExpr(value.(*array.String), not(reMatch(m)))
			default:
				return call(op.String(), nil, value, &compute.ScalarDatum{
					Value: scalar.MakeScalar(matchValue),
				})
			}
		},
	}
}

func not(m func(string) bool) func(string) bool {
	return func(s string) bool {
		return !m(s)
	}
}

func reMatch(r string) func(s string) bool {
	x := regexp.MustCompile(r)
	return x.MatchString
}

func boolExpr(s *array.String, f func(string) bool) (arrow.Array, error) {
	b := array.NewBooleanBuilder(entry.Pool)
	defer b.Release()
	b.Reserve(s.Len())
	for i := 0; i < s.Len(); i++ {
		b.UnsafeAppend(f(s.Value(i)))
	}
	return b.NewBooleanArray(), nil
}