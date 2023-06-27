package timeseries

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"sync"
	"time"

	"github.com/cespare/xxhash/v2"
	"github.com/dgraph-io/badger/v4"
	"github.com/vinceanalytics/vince/internal/core"
	"github.com/vinceanalytics/vince/pkg/log"
)

type mergeFunction func(context.Context, uint64, *kvTs, *mergeStats) error

func Merge(ctx context.Context) {
	stats := storeForever(ctx, forever)
	log.Get().Debug().
		Int("visited", stats.keys.visited).
		Int("skipped", stats.keys.skipped).
		Int("accepted", stats.keys.accepted).
		Int("processed", stats.keys.processed).
		Msgf("merged in %s", stats.elapsed)
}

type mergeStats struct {
	elapsed time.Duration
	keys    struct {
		visited, skipped, accepted, processed int
	}
}

func forever(ctx context.Context, ts uint64, kv *kvTs, stats *mergeStats) error {
	db := Get(ctx)
	txn := db.NewTransactionAt(ts, true)
	s := newSlice()
	for _, b := range kv.b {
		if err := store(txn, s, b.k.Bytes(), b.b); err != nil {
			if !errors.Is(err, badger.ErrTxnTooBig) {
				s.release()
				txn.Discard()
				return err
			}
			err = txn.CommitAt(ts, nil)
			if err != nil {
				return err
			}
			txn = db.NewTransactionAt(ts, true)
		}
	}
	err := txn.CommitAt(ts, nil)
	if err != nil {
		s.release()
		txn.Discard()
		return err
	}
	s.release()
	stats.keys.processed++
	return nil
}

func store(txn *badger.Txn, sl *slice, key []byte, value uint64) error {
	g, err := txn.Get(key)
	if err != nil {
		if errors.Is(err, badger.ErrKeyNotFound) {
			return txn.Set(key, sl.u64(value))
		}
		return err
	}
	return g.Value(func(val []byte) error {
		return txn.Set(key, sl.u64(binary.BigEndian.Uint64(val)+value))
	})
}

func storeForever(ctx context.Context, mergeFn mergeFunction) (stats mergeStats) {
	start := core.Now(ctx)

	ts := uint64(start.UnixMilli())
	txn := GetMike(ctx).NewTransactionAt(ts, true)
	o := badger.DefaultIteratorOptions
	it := txn.NewIterator(o)
	ls := newTxnBufferList()
	defer ls.release()

	m := mergePool.Get().(*merge)
	defer m.release()
	defer func() {
		ls.release()
		m.release()
		stats.elapsed = core.Now(ctx).Sub(start)
	}()

	for it.Rewind(); it.Valid(); it.Next() {
		stats.keys.visited++
		x := it.Item()
		if x.IsDeletedOrExpired() {
			stats.keys.skipped++
			continue
		}
		key := x.Key()
		x.Value(func(val []byte) error {
			m.add(key, val)
			return nil
		})
		k := ls.Get()
		k.Write(key)
		txn.Delete(k.Bytes())
		stats.keys.accepted++
	}
	it.Close()
	err := txn.CommitAt(ts, nil)
	if err != nil {
		log.Get().Err(err).Msg("failed to commit merge transaction")
		return
	}
	err = m.do(ctx, mergeFn, &stats)
	if err != nil {
		log.Get().Err(err).Msg("failed merge operation")
	}
	return
}

type merge struct {
	ts map[uint64]*kvTs
	m  map[uint64]*kvBuf
	h  *xxhash.Digest
}

func (m *merge) release() {
	for k, v := range m.m {
		v.reset()
		delete(m.m, k)
	}
	for k, v := range m.ts {
		v.reset()
		delete(m.ts, k)
	}
	m.h.Reset()
	mergePool.Put(m)
}

var mergePool = &sync.Pool{
	New: func() any {
		return &merge{
			ts: make(map[uint64]*kvTs),
			m:  make(map[uint64]*kvBuf),
			h:  xxhash.New(),
		}
	},
}

func (m *merge) hash(b []byte) uint64 {
	m.h.Reset()
	m.h.Write(b)
	return m.h.Sum64()
}

func (m *merge) add(key, value []byte) {
	m.addInternal(key, value)
	if key[propOffset] == byte(Base) {
		// Store global stats. Global stats are grouped into
		//  Per Site : This is covered by the base property. It aggregates metrics per site.
		//  Per User:  This is covered by setting site to zero
		//  Per Instance : Both user and site are set to zero.
		// We also want to be able to chart or compute diffs between global stats so
		// we provide variations of per user and per instance stats with timestamp
		// appended.
		b := get()
		b.Write(key)
		g := b.Bytes()

		// per user global  stats with timestamp
		copy(g[siteOffset:], zero)
		m.addInternal(g, value)

		// per vince instance global stats with timestamp
		copy(g[userOffset:], zero)
		m.addInternal(g, value)

		b.Reset()
		b.Write(key)
		g = b.Bytes()

		ts := binary.BigEndian.Uint64(g[len(g)-8:])
		// spot stats. Think of this as a global counter. Not tied to specific
		// timestamp
		copy(g[len(g)-8:], zero)

		copy(g[siteOffset:], zero)
		m.addInternal(g, value, ts)

		copy(g[userOffset:], zero)
		m.addInternal(g, value, ts)
		put(b)
	}
}

func (m *merge) addInternal(key, value []byte, stamp ...uint64) {
	m.h.Reset()
	baseKey := key[:len(key)-8]
	baseTs := key[len(key)-8:]
	keyHash := m.hash(baseKey)
	var ts uint64
	if len(stamp) > 0 {
		ts = stamp[0]
	} else {
		ts = binary.BigEndian.Uint64(baseTs)
	}
	m.h.Reset()
	v := binary.BigEndian.Uint16(value)
	b, ok := m.m[keyHash]
	if ok {
		// existing key
		b.b += uint64(v)
	} else {
		b = kvBufPool.Get().(*kvBuf)
		b.k = get()
		b.k.Write(key)
		b.b = uint64(v)
		m.m[keyHash] = b
		t, ok := m.ts[ts]
		if !ok {
			t = kvTsPool.Get().(*kvTs)
			m.ts[ts] = t
		}
		t.b = append(t.b, b)
	}
}

func (m *merge) do(ctx context.Context, f mergeFunction, stats *mergeStats) error {
	for k, kt := range m.ts {
		err := f(ctx, k, kt, stats)
		if err != nil {
			return err
		}
	}
	return nil
}

type kvTs struct {
	b []*kvBuf
}

func (k *kvTs) reset() {
	k.b = k.b[:0]
	kvTsPool.Put(k)
}

var kvTsPool = &sync.Pool{
	New: func() any {
		return &kvTs{
			b: make([]*kvBuf, 0, 1<<10),
		}
	},
}

type kvBuf struct {
	k *bytes.Buffer
	b uint64
}

func (k *kvBuf) reset() {
	put(k.k)
	kvBufPool.Put(k)
}

var kvBufPool = &sync.Pool{
	New: func() any {
		return &kvBuf{}
	},
}
