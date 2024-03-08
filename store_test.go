package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	impl := []Store{
		&MemStore{},
	}
	for _, i := range impl {
		testStoreImpl(t, i)
	}
}

func testStoreImpl(t *testing.T, s Store) {

	item1 := &Workday{Date: time.Now(), Location: Home}
	item2 := &Workday{Date: time.Now(), Location: Office}

	t.Run("store is initially empty", func(t *testing.T) {
		data, err := s.LoadAll()
		assert.NoError(t, err)
		assert.Len(t, data, 0)
	})

	t.Run("add items to store", func(t *testing.T) {
		assert.NoError(t, s.Save(item1))
		assert.NoError(t, s.Save(item2))
	})

	t.Run("load items in order", func(t *testing.T) {
		data, err := s.LoadAll()
		assert.NoError(t, err)
		assert.Equal(t, item1, data[0])
		assert.Equal(t, item2, data[1])
	})

}
