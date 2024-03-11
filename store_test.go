package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFileStore(t *testing.T) {

	path := "testfile.txt"

	fs := NewFileStore(path)
	defer os.Remove(path)

	// flag today a couple of times
	require.NoError(t, fs.Flag())
	require.NoError(t, fs.Flag())
	require.NoError(t, fs.Flag())

	// the data should contain today's date
	data, err := fs.Load()
	require.NoError(t, err)
	assert.True(t, data[time.Now().Truncate(24*time.Hour)])

	// clear the data file and verify that it no longer contains anything
	require.NoError(t, fs.Clear())
	data, err = fs.Load()
	require.NoError(t, err)
	assert.Len(t, data, 0)
}
