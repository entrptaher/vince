package shards

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"sync"
	"time"

	"github.com/cockroachdb/pebble"
	"github.com/gernest/roaring"
	"github.com/vinceanalytics/vince/internal/compute"
	"github.com/vinceanalytics/vince/internal/encoding"
	"github.com/vinceanalytics/vince/internal/models"
	"github.com/vinceanalytics/vince/internal/ro2"
	"github.com/vinceanalytics/vince/internal/timeseries/cursor"
	"github.com/vinceanalytics/vince/internal/util/assert"
	"github.com/vinceanalytics/vince/internal/util/data"
)

func Format(shard uint64) string {
	return fmt.Sprintf("%006d", shard)
}

func Parse(str string) uint64 {
	a, err := strconv.ParseUint(str, 10, 64)
	assert.Nil(err)
	return a
}

type Views [encoding.Month + 1]*ro2.Bitmap

func (v *Views) Init() {
	for i := range v {
		v[i] = ro2.NewBitmap()
	}
}

func (v *Views) Reset() {
	for i := range v {
		v[i].Containers.Reset()
	}
}

type DB struct {
	ops    *pebble.DB
	base   string
	shards struct {
		sync.RWMutex
		data [1 << 10]*Shard
		max  uint64
	}
}

func New(base string) (*DB, error) {
	ops, err := data.Open(filepath.Join(base, "ops"), nil)
	if err != nil {
		return nil, err
	}
	db := &DB{ops: ops, base: base}
	shardsBase := filepath.Join(base, "shards")
	entries, _ := os.ReadDir(shardsBase)
	shards := make([]uint64, 0, len(entries))
	for _, e := range entries {
		shards = append(shards, Parse(e.Name()))
	}
	slices.Sort(shards)
	for _, sh := range shards {
		db.shards.data[sh] = NewShard(shardsBase, sh)
		db.shards.max = max(db.shards.max, sh)
	}
	return db, nil
}

func (db *DB) Get() *pebble.DB {
	return db.ops
}

func (db *DB) Iter(
	domainId uint64,
	re encoding.Resolution,
	start, end time.Time,
	filters []models.Field,
	f func(cu *cursor.Cursor, shard, view uint64, match *ro2.Bitmap, exists map[models.Field]*ro2.Bitmap) error) error {
	db.shards.RLock()
	defer db.shards.RUnlock()

	views := slices.Collect(compute.Range(re, start, end))
	if len(views) == 0 {
		return nil
	}

	cu := new(cursor.Cursor)
	defer cu.Release()

	for i := uint64(0); i <= db.shards.max; i++ {
		sh := db.shards.data[i]
		if sh == nil {
			continue
		}
		it, err := sh.DB.NewIter(nil)
		if err != nil {
			return err
		}
		cu.SetIter(it)

		for _, view := range views {
			if !cu.ResetData(re, models.Field_domain, view) {
				continue
			}
			m := ro2.Row(cu, sh.ID, domainId)
			if !m.Any() {
				continue
			}
			exists := readExistence(cu, re, filters, sh.ID, view)
			err := f(cu, sh.ID, view, m, exists)
			if err != nil {
				it.Close()
				return err
			}
		}
		it.Close()
	}
	return nil
}

func readExistence(cu *cursor.Cursor, re encoding.Resolution, fields []models.Field, shard, view uint64) (m map[models.Field]*ro2.Bitmap) {
	m = make(map[models.Field]*roaring.Bitmap)
	for _, f := range fields {
		if !cu.ResetExistence(re, f, view) {
			continue
		}
		m[f] = ro2.Existence(cu, shard)
	}
	return
}

func (db *DB) Shard(shard uint64) *Shard {
	db.shards.RLock()
	sh := db.shards.data[shard]
	db.shards.RUnlock()
	if sh != nil {
		return sh
	}
	db.shards.Lock()
	defer db.shards.Unlock()
	sh = NewShard(filepath.Join(db.base, "shards"), shard)
	db.shards.data[shard] = sh
	db.shards.max = max(db.shards.max, shard)
	return sh
}

func (db *DB) Close() error {
	var errs []error
	errs = append(errs, db.ops.Close())
	db.shards.Lock()
	defer db.shards.Unlock()

	for _, sh := range db.shards.data {
		if sh != nil {
			errs = append(errs, sh.DB.Close())
		}
	}
	return errors.Join(errs...)
}

type Shard struct {
	ID uint64
	DB *pebble.DB
}

func NewShard(base string, shard uint64) *Shard {
	path := filepath.Join(base, Format(shard))
	db, err := data.Open(path, &pebble.Options{
		Merger: ro2.Merge,
	})
	assert.Nil(err, fmt.Sprintf("opening database shard path=%q", path))
	s := &Shard{ID: shard, DB: db}
	return s
}

func rangeBits(b *ro2.Bitmap, start, end uint64) *ro2.Bitmap {
	itr := b.Iterator()
	itr.Seek(start)
	o := ro2.NewBitmap()
	for v, eof := itr.Next(); !eof && v <= end; v, eof = itr.Next() {
		o.Add(v)
	}
	return o
}
