package store

import (
	"errors"
	"sync"

	"filippo.io/age"
	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/ristretto/z"
	"github.com/vinceanalytics/vince/internal/encoding"
	"github.com/vinceanalytics/vince/internal/keys"
	"github.com/vinceanalytics/vince/internal/models"
	"github.com/vinceanalytics/vince/internal/roaring"
	"github.com/vinceanalytics/vince/internal/util/trie"
)

type Store struct {
	db *badger.DB

	// saves db lookups when querying
	mu   sync.RWMutex
	trie *trie.Trie

	// below fields are used for batching which only occurs in a single goroutine.
	// They are not thread safe.
	ranges   [models.TranslatedFieldsSize]*badger.Sequence
	keys     [models.TranslatedFieldsSize][][]byte
	values   [models.TranslatedFieldsSize][]uint64
	tree     *z.Tree
	mutex    [models.TranslatedFieldsSize]map[uint64]*roaring.Bitmap
	bsi      [models.BSIFieldsSize]map[uint64]*roaring.Bitmap
	enc      encoding.Encoding
	id       uint64
	time     uint64
	shard    uint64
	events   uint64
	txnCount int
}

func Open(path string) (*Store, error) {
	return open(path)
}

func open(path string) (*Store, error) {
	return newDB(path)
}

func (o *Store) Web() (secret *age.X25519Identity, err error) {
	err = o.Update(func(tx *Tx) error {
		key := keys.Cookie
		it, err := tx.tx.Get(key)
		if err != nil {
			if !errors.Is(err, badger.ErrKeyNotFound) {
				return err
			}
			secret, err = age.GenerateX25519Identity()
			if err != nil {
				return err
			}
			return tx.tx.Set(key, []byte(secret.String()))
		}
		return it.Value(func(val []byte) error {
			secret, err = age.ParseX25519Identity(string(val))
			return err
		})
	})
	return
}
