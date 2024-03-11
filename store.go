package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const DateFormat = "20060102"

type Store interface {
	// Flag today as an office day
	Flag() error
	// Load all office days currently in the store
	Load() (map[time.Time]bool, error)
	// Clear the currently saved office days
	Clear() error
}

type FileStore struct {
	path string
}

func NewFileStore(path string) *FileStore {
	return &FileStore{
		path: path,
	}
}

// Flag today as an office day, saving the date to the data file, unless it's already there
func (f *FileStore) Flag() error {
	date := time.Now().Truncate(24 * time.Hour)
	days, err := f.Load()
	if err != nil {
		return err
	}
	if days[date] {
		return nil // already flagged, nothing to do
	}
	file, err := os.OpenFile(f.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s\n", date.Format(DateFormat)))
	return err
}

// Load all office days currently in the store, using the data file
func (f *FileStore) Load() (map[time.Time]bool, error) {
	result := make(map[time.Time]bool)
	if _, err := os.Stat(f.path); errors.Is(err, os.ErrNotExist) {
		return result, nil
	}
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" {
			continue // skip empty lines
		}
		t, err := time.Parse(DateFormat, line)
		if err != nil {
			return nil, err
		}
		result[t.Local()] = true
	}
	return result, nil
}

// Clear the data file
func (f *FileStore) Clear() error {
	return os.Remove(f.path)
}
