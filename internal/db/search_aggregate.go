package db

import (
	"cmp"
	"context"
	"slices"
	"time"

	"github.com/RoaringBitmap/roaring/roaring64"
	"github.com/gernest/rbf/dsl/boolean"
	"github.com/gernest/rbf/dsl/bsi"
	"github.com/gernest/rbf/dsl/mutex"
	"github.com/gernest/rbf/dsl/tx"
	"github.com/gernest/rows"
	v1 "github.com/vinceanalytics/vince/gen/go/vince/v1"
	"github.com/vinceanalytics/vince/internal/defaults"
	"github.com/vinceanalytics/vince/internal/timeutil"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (db *DB) Aggregate(ctx context.Context, req *v1.Aggregate_Request) (*v1.Aggregate_Response, error) {
	defaults.Set(req)
	err := validate.Validate(req)
	if err != nil {
		return nil, err
	}
	m := dupe(req.Metrics)
	a := newAggregate(m)
	from, to := periodToRange(req.Period, req.Date)
	props := append(req.Filters,
		&v1.Filter{Property: v1.Property_domain, Op: v1.Filter_equal, Value: req.SiteId},
	)
	ts := bsi.Filter("timestamp", bsi.RANGE, from.UnixMilli(), to.UnixMilli())
	fs := filterProperties(props...)

	r, err := db.db.Reader()
	if err != nil {
		return nil, err
	}
	defer r.Release()
	for _, shard := range r.Range(from, to) {
		err := r.View(shard, func(txn *tx.Tx) error {
			f, err := ts.Apply(txn, nil)
			if err != nil {
				return err
			}
			if f.IsEmpty() {
				return nil
			}
			r, err := fs.Apply(txn, f)
			if err != nil {
				return err
			}
			if r.IsEmpty() {
				return nil
			}
			return a.Apply(txn, r)
		})
		if err != nil {
			return nil, err
		}
	}
	res := &v1.Aggregate_Response{
		Results: make(map[string]float64),
	}
	for _, mx := range m {
		res.Results[mx.String()] = a.Result(mx)
	}
	return res, nil
}

func dupe[T cmp.Ordered](a []T) []T {
	m := make(map[T]struct{})
	for i := range a {
		m[a[i]] = struct{}{}
	}
	o := make([]T, 0, len(m))
	for k := range m {
		o = append(o, k)
	}
	slices.Sort(o)
	return o
}

type aggregateQuery struct {
	a applyList
}

var _ Query = (*aggregateQuery)(nil)

func (a *aggregateQuery) View(_ time.Time) View {
	return a.a
}

type aggregate struct {
	visitors    roaring64.Bitmap
	visits      int64
	views       int64
	bounceTrue  int64
	bounceFalse int64
	events      int64
	duration    int64
	cache       applyList
}

func (a *aggregate) or(b *aggregate) {
	a.visitors.Or(&b.visitors)
}

func newAggregate(metrics []v1.Metric) *aggregate {
	a := &aggregate{}
	a.newApplyList(metrics)
	return a
}

func (a *aggregate) Apply(tx *tx.Tx, columns *rows.Row) error {
	return a.cache.Apply(tx, columns)
}

func (a *aggregate) Result(m v1.Metric) float64 {
	switch m {
	case v1.Metric_visitors:
		return float64(a.Visitors())
	case v1.Metric_visits:
		return float64(a.Visits())
	case v1.Metric_pageviews:
		return float64(a.Views())
	case v1.Metric_views_per_visit:
		return a.ViewsPerVisit()
	case v1.Metric_bounce_rate:
		return a.BounceRate()
	case v1.Metric_visit_duration:
		d := a.Duration()
		// convert to seconds
		return (time.Duration(d) * time.Millisecond).Seconds()
	case v1.Metric_events:
		return float64(a.Events())
	default:
		return 0
	}
}

func (a *aggregate) newApplyList(m []v1.Metric) applyList {
	if len(a.cache) > 0 {
		return a.cache
	}
	o := make(map[string]struct{})
	for i := range m {
		switch m[i] {
		case v1.Metric_visitors:
			o["visitors"] = struct{}{}
		case v1.Metric_visits:
			o["visits"] = struct{}{}
		case v1.Metric_pageviews:
			o["views"] = struct{}{}
		case v1.Metric_views_per_visit:
			o["views"] = struct{}{}
			o["visits"] = struct{}{}
		case v1.Metric_bounce_rate:
			o["bounce"] = struct{}{}
			o["visits"] = struct{}{}
		case v1.Metric_visit_duration:
			o["duration"] = struct{}{}
		case v1.Metric_events:
			o["events"] = struct{}{}
		}
	}
	ls := make(applyList, 0, len(o))
	for k := range o {
		switch k {
		case "visitors":
			ls = append(ls, a.applyVisitors)
		case "visits":
			ls = append(ls, a.applyVisits)
		case "views":
			ls = append(ls, a.applyViews)
		case "bounce":
			ls = append(ls, a.applyBounce)
		case "duration":
			ls = append(ls, a.applyDuration)
		case "events":
			ls = append(ls, a.applyEvents)
		}
	}
	a.cache = ls
	return ls
}

type applyList []func(*tx.Tx, *rows.Row) error

var _ View = (*applyList)(nil)

func (ls applyList) Apply(tx *tx.Tx, columns *rows.Row) error {
	for i := range ls {
		err := ls[i](tx, columns)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *aggregate) Reset() {
	a.visitors.Clear()
	a.visits = 0
	a.views = 0
	a.events = 0
	a.bounceTrue = 0
	a.bounceFalse = 0
	a.duration = 0
}

func (a *aggregate) BounceRate() float64 {
	bounce := a.Bounce()
	visits := a.Visits()
	if visits != 0 {
		return float64(bounce) / float64(visits)
	}
	return 0
}

func (a *aggregate) ViewsPerVisit() float64 {
	views := a.Views()
	visits := a.Visits()
	if visits != 0 {
		return float64(views) / float64(visits)
	}
	return 0
}

func (a *aggregate) Events() uint64 {
	return a.Events()
}

func (a *aggregate) Visitors() uint64 {
	return a.visitors.GetCardinality()
}

func (a *aggregate) Visits() uint64 {
	return uint64(a.visits)
}

func (a *aggregate) Views() uint64 {
	return uint64(a.views)
}

func (a *aggregate) Bounce() uint64 {
	if a.bounceTrue < a.bounceFalse {
		return uint64(a.bounceTrue - a.bounceFalse)
	}
	return 0
}

func (a *aggregate) Duration() uint64 {
	return uint64(a.duration)
}

func (a *aggregate) applyEvents(tx *tx.Tx, columns *rows.Row) error {
	a.events += int64(columns.Count())
	return nil
}

func (a *aggregate) applyVisitors(tx *tx.Tx, columns *rows.Row) error {
	return mutex.Distinct(tx, "id", &a.visitors, columns)
}

func (a *aggregate) applyDuration(tx *tx.Tx, columns *rows.Row) error {
	_, sum, err := bsi.SumCount(tx, "duration", nil, columns)
	if err != nil {
		return err
	}
	a.duration += sum
	return nil
}

func (a *aggregate) applyVisits(tx *tx.Tx, columns *rows.Row) error {
	count, err := boolean.Count(tx, "session", true, columns)
	a.visits += count
	return err
}

func (a *aggregate) applyViews(tx *tx.Tx, columns *rows.Row) error {
	count, err := boolean.Count(tx, "view", true, columns)
	a.visits += count
	return err
}

func (a *aggregate) applyBounce(tx *tx.Tx, columns *rows.Row) error {
	count, err := boolean.Count(tx, "bounce", true, columns)
	if err != nil {
		return err
	}
	a.bounceTrue += count
	count, err = boolean.Count(tx, "bounce", true, columns)
	if err != nil {
		return err
	}
	a.bounceFalse += count
	return nil
}

func periodToRange(period *v1.TimePeriod, tsDate *timestamppb.Timestamp) (start, end time.Time) {
	date := tsDate.AsTime()
	switch e := period.Value.(type) {
	case *v1.TimePeriod_Base_:
		switch e.Base {
		case v1.TimePeriod_day:
			end = date
			start = timeutil.BeginDay(end)
		case v1.TimePeriod__7d:
			end = date
			start = end.AddDate(0, 0, -6)
		case v1.TimePeriod__30d:
			end = date
			start = end.AddDate(0, 0, -30)
		case v1.TimePeriod_mo:
			end = date
			start = timeutil.BeginMonth(end)
			end = timeutil.EndMonth(end)
		case v1.TimePeriod__6mo:
			end = timeutil.EndMonth(date)
			start = timeutil.BeginMonth(end.AddDate(0, -5, 0))
		case v1.TimePeriod__12mo:
			end = timeutil.EndMonth(date)
			start = timeutil.BeginMonth(end.AddDate(0, -11, 0))
		case v1.TimePeriod_year:
			end = timeutil.EndYear(date)
			start = timeutil.BeginYear(end)
		}

	case *v1.TimePeriod_Custom_:
		end = e.Custom.End.AsTime()
		start = e.Custom.Start.AsTime()
	}
	return
}

type nowFunc func() time.Time

func parseDate(now nowFunc, ts *timestamppb.Timestamp) time.Time {
	if ts != nil {
		return ts.AsTime()
	}
	return timeutil.EndDay(now())
}

func ValidByPeriod(period *v1.TimePeriod, i v1.Interval) bool {
	switch e := period.Value.(type) {
	case *v1.TimePeriod_Base_:
		switch e.Base {
		case v1.TimePeriod_day:
			return i == v1.Interval_minute || i == v1.Interval_hour
		case v1.TimePeriod__7d:
			return i == v1.Interval_hour || i == v1.Interval_date
		case v1.TimePeriod_mo, v1.TimePeriod__30d:
			return i == v1.Interval_date || i == v1.Interval_week
		case v1.TimePeriod__6mo, v1.TimePeriod__12mo, v1.TimePeriod_year:
			return i == v1.Interval_date || i == v1.Interval_week || i == v1.Interval_month
		default:
			return false
		}
	case *v1.TimePeriod_Custom_:
		return i == v1.Interval_date || i == v1.Interval_week || i == v1.Interval_month
	default:
		return false
	}
}