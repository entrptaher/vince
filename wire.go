package len64

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"regexp"

	"github.com/RoaringBitmap/roaring/v2/roaring64"
	"github.com/cespare/xxhash/v2"
	"github.com/cockroachdb/pebble"
)

var (
	shardPrefix     = byte(0x0)
	timeRangePrefix = byte(0x1)
	trPrefix        = byte(0x2)
	seqKey          = []byte{0x3}
	bsiPrefix       = []byte{shardPrefix, 0x1}
	trKeyPrefix     = []byte{trPrefix, 0x2}
	trIDPrefix      = []byte{trPrefix, 0x3}
	sep             = []byte{'='}
)

const (
	shardWidth = 1048576
)

type Batch struct {
	seq      uint64
	shard    uint64
	db       *pebble.DB
	strings  map[uint64]string
	hash     xxhash.Digest
	b        bytes.Buffer
	m        map[string]*roaring64.BSI
	min, max uint64
}

func newBatch(db *pebble.DB) (*Batch, error) {
	return &Batch{
		seq:     ReadSeq(db),
		shard:   zero,
		db:      db,
		strings: map[uint64]string{},
		m:       make(map[string]*roaring64.BSI),
	}, nil
}

const zero = ^uint64(0)

func (i *Batch) Write(ts uint64, f func(idx Index)) error {
	i.seq++
	shard := i.seq / shardWidth
	if shard != i.shard {
		if i.shard != zero {
			err := i.emit()
			if err != nil {
				return err
			}
		}
		i.shard = shard
	}
	f(i)
	if i.min == 0 {
		i.min = ts
	} else {
		i.min = min(i.min, ts)
	}
	i.max = max(i.max, ts)
	return nil
}

func (i *Batch) Release() error {
	defer func() {
		WriteSeq(i.db, i.seq)
		i.db = nil
	}()
	return i.emit()
}

func (i *Batch) Flush() error {
	defer func() {
		i.shard = zero
	}()
	return i.emit()
}

func (i *Batch) emit() error {
	defer func() {
		clear(i.strings)
		clear(i.m)
	}()

	b := i.db.NewBatch()

	err := WriteBSI(b, i.shard, i.m)
	if err != nil {
		return err
	}

	err = WriteTimeRange(b, i.shard, i.min, i.max)
	if err != nil {
		return err
	}

	err = WriteString(b, i.shard, i.strings)
	if err != nil {
		return err
	}

	return i.db.Apply(b, nil)
}

type Index interface {
	Int64(field string, value int64)
	String(field string, value string)
	Bool(field string, value bool)
}

func (i *Batch) Int64(field string, value int64) {
	i.get(field).SetValue(i.seq, value)
}

func (i *Batch) Bool(field string, value bool) {
	n := int64(0)
	if value {
		n = 1
	}
	i.get(field).SetValue(i.seq, n)
}

func (i *Batch) String(field string, value string) {
	i.b.Reset()
	i.b.WriteString(field)
	i.b.Write(sep)
	i.b.WriteString(value)

	i.hash.Reset()
	i.hash.Write(i.b.Bytes())
	sum := i.hash.Sum64()
	i.strings[sum] = i.b.String()
	i.get(field).SetValue(i.seq, int64(sum))
}

func (i *Batch) get(name string) *roaring64.BSI {
	b, ok := i.m[name]
	if !ok {
		b = roaring64.NewDefaultBSI()
		i.m[name] = b
	}
	return b
}

func WriteSeq(db *pebble.DB, seq uint64) error {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], seq)
	return db.Set(seqKey, b[:], nil)
}

func ReadSeq(db *pebble.DB) uint64 {
	value, done, err := db.Get(seqKey)
	if err != nil {
		return 0
	}
	seq := binary.BigEndian.Uint64(value)
	done.Close()
	return seq
}

func WriteBSI(b *pebble.Batch, shard uint64, m map[string]*roaring64.BSI) error {
	if len(m) == 0 {
		return nil
	}
	key := make([]byte, 1<<10)
	copy(key, bsiPrefix)
	binary.BigEndian.PutUint64(key[2:], shard)

	var buf bytes.Buffer

	for name, v := range m {
		key = append(key[:10], []byte(name)...)
		buf.Reset()
		v.RunOptimize()
		_, err := v.WriteTo(&buf)
		if err != nil {
			return fmt.Errorf("writing bsi %d:%s%w", shard, name, err)
		}

		err = b.Merge(key, buf.Bytes(), nil)
		if err != nil {
			return fmt.Errorf("merge bsi %d:%s%w", shard, name, err)
		}
	}
	return nil
}

func ReadBSI(db *pebble.Snapshot, shard uint64, name string) (*roaring64.BSI, error) {
	key := make([]byte, 1<<10)
	copy(key, bsiPrefix)
	binary.BigEndian.PutUint64(key[2:], shard)
	key = append(key[:10], []byte(name)...)

	value, done, err := db.Get(key)
	if err != nil {
		return nil, err
	}
	defer done.Close()
	r := roaring64.NewDefaultBSI()
	_, err = r.ReadFrom(bytes.NewReader(value))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func bsiFrom(value []byte) (*roaring64.BSI, error) {
	r := roaring64.NewDefaultBSI()
	_, err := r.ReadFrom(bytes.NewReader(value))
	if err != nil {
		return nil, err
	}
	return r, nil
}

func WriteString(b *pebble.Batch, shard uint64, m map[uint64]string) error {
	if len(m) == 0 {
		return nil
	}
	key := make([]byte, 1<<10)
	copy(key, trKeyPrefix)

	value := make([]byte, 2+8)
	copy(value, trIDPrefix)

	for id, v := range m {
		key = append(key[:2], []byte(v)...)
		err := b.Set(key, []byte{}, nil)
		if err != nil {
			return fmt.Errorf("write string key %d:%s%w", shard, v, err)
		}

		binary.BigEndian.PutUint64(value[2:], id)
		err = b.Set(value, []byte(v), nil)
		if err != nil {
			return fmt.Errorf("write string id %d:%s%w", shard, v, err)
		}
	}
	return nil
}

func Search(db *pebble.DB, shard uint64, k, v string) (uint64, bool) {
	var b bytes.Buffer
	b.Write(trKeyPrefix)
	var x [8]byte
	binary.BigEndian.PutUint64(x[:], shard)
	b.Write(x[:])
	b.WriteString(k)
	b.WriteByte('=')
	b.WriteString(v)
	full := b.Bytes()
	_, done, err := db.Get(full)
	if err != nil {
		return 0, false
	}
	done.Close()
	return xxhash.Sum64(full[10:]), true
}

func SearchRegex(db *pebble.Snapshot, shard uint64, k, v string) ([]uint64, error) {
	var b bytes.Buffer
	b.Write(trKeyPrefix)
	var x [8]byte
	binary.BigEndian.PutUint64(x[:], shard)
	b.Write(x[:])
	b.WriteString(k)
	b.WriteByte('=')

	re, err := regexp.Compile(v)
	if err != nil {
		return nil, err
	}
	prefix := b.Bytes()

	it, err := db.NewIter(nil)
	if err != nil {
		return nil, err
	}
	result := make([]uint64, 0, 4)
	h := new(xxhash.Digest)
	for it.SeekGE(prefix); bytes.HasPrefix(it.Key(), prefix); it.Next() {
		full := it.Key()
		str := full[10:]
		_, value, _ := bytes.Cut(str, sep)
		if re.Match(value) {
			h.Reset()
			h.Write(str)
			result = append(result, h.Sum64())
		}
	}
	return result, nil
}

func WriteTimeRange(b *pebble.Batch, shard uint64, min, max uint64) error {
	key := make([]byte, 1+8+8)
	key[0] = timeRangePrefix
	binary.BigEndian.PutUint64(key[1:], min)
	binary.BigEndian.PutUint64(key[1+8:], shard)

	err := b.Set(key, []byte{}, nil)
	if err != nil {
		return fmt.Errorf("write time range  %d:%d%w", shard, min, err)
	}
	binary.BigEndian.PutUint64(key[1:], max)
	err = b.Set(key, []byte{}, nil)
	if err != nil {
		return fmt.Errorf("write time range  %d:%d%w", shard, max, err)
	}
	return nil
}

func ReadTimeRange(db *pebble.Snapshot, start, end uint64, b *roaring64.Bitmap) error {
	key := make([]byte, 1+8+8)
	key[0] = timeRangePrefix
	binary.BigEndian.PutUint64(key[1:], start)
	binary.BigEndian.PutUint64(key[1+8:], 0)

	from := bytes.Clone(key)
	binary.BigEndian.PutUint64(key[1:], end)
	binary.BigEndian.PutUint64(key[1+8:], math.MaxUint64)

	it, err := db.NewIter(&pebble.IterOptions{
		LowerBound: from,
		UpperBound: key,
	})
	if err != nil {
		return err
	}
	defer it.Close()
	for it.First(); it.Valid(); it.Next() {
		b.Add(
			binary.BigEndian.Uint64(it.Key()[1+8:]),
		)
	}
	return nil
}
