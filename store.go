package main

import (
	"github.com/peterbourgon/diskv/v3"
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
}

type DiskDB struct {
	db *diskv.Diskv
}

func NewDiskDB(path string) *DiskDB {
	return &DiskDB{
		db: diskv.New(diskv.Options{
			BasePath: path,
			Transform: func(s string) []string {
				return []string{} // put all keys in base dir
			},
		}),
	}
}

// Save a given date
func (d *DiskDB) Save(date string, office bool) error {
	var b byte = 0
	if office {
		b = 1
	}
	return d.db.Write(date, []byte{b})
}

// Lookup a given date
func (d *DiskDB) Lookup(date string) (bool, error) {
	if !d.db.Has(date) {
		return false, nil
	}
	res, err := d.db.Read(date)
	if err != nil {
		return false, err
	}
	return res[0] == 1, nil
}

// Load all dates
func (d *DiskDB) LoadAll() (map[string]bool, error) {
	res := make(map[string]bool)
	for key := range d.db.Keys(make(<-chan struct{})) {
		val, err := d.db.Read(key)
		if err != nil {
			return nil, err
		}
		res[key] = val[0] == 1
	}
	return res, nil
}

// Clear all data
func (d *DiskDB) Clear() error {
	return d.db.EraseAll()
}
