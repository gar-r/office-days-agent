package main

import (
	"github.com/dgraph-io/badger/v4"
)

type DB interface {
	// Save a given date
	Save(date string, office bool) error
	// Lookup a given date
	Lookup(date string) (bool, error)
	// Load all dates
	LoadAll() (map[string]bool, error)
	// Clear all data
	Clear() error
	// Close the db
	Close()
}

type Badger struct {
	db *badger.DB
}

func NewBadger(path string) (*Badger, error) {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		return nil, err
	}
	return &Badger{db}, nil
}

// Save a given date
func (b *Badger) Save(date string, office bool) error {
	return b.db.Update(func(txn *badger.Txn) error {
		var b byte
		if office {
			b = 1
		}
		return txn.Set([]byte(date), []byte{b})
	})
}

// Lookup a given date
func (b *Badger) Lookup(date string) (bool, error) {
	res := false
	err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(date))
		if err == badger.ErrKeyNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			if len(val) > 0 && val[0] == 1 {
				res = true
			}
			return nil
		})
	})
	return res, err
}

// Load all dates
func (b *Badger) LoadAll() (map[string]bool, error) {
	res := make(map[string]bool)
	err := b.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			date := string(item.Key())
			return item.Value(func(val []byte) error {
				res[date] = len(val) > 0 && val[0] == 1
				return nil
			})
		}
		return nil
	})
	return res, err
}

// Clear all data
func (b *Badger) Clear() error {
	return b.db.DropAll()
}

// Close the db
func (b *Badger) Close() {
	b.db.Close()
}
