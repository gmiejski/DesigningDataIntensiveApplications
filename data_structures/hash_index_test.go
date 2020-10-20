package data_structures

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const hashIndexDir = "/tmp/ddia/hash_index/d240d955-661b-4f6c-8e48-e85b8c14a9e4"

func TestHashIndexErrorWhenNotFound(t *testing.T) {
	defer cleanupHashIndex()
	// given
	db, err := newHashIndex(hashIndexDir)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	assert.EqualValues(t, Person{}, value)
	assert.Error(t, err)
}

func TestHashIndexReadWriteValue(t *testing.T) {
	defer cleanupHashIndex()
	// given
	db, err := newHashIndex(hashIndexDir)
	require.NoError(t, err)
	err = db.Save(1, john)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	require.NoError(t, err)
	require.Equal(t, john, value)
}

func TestHashIndexReadingNewestValue(t *testing.T) {
	defer cleanupHashIndex()
	// given
	db, err := newHashIndex(hashIndexDir)
	require.NoError(t, err)
	err = db.Save(1, john)
	err = db.Save(1, john2)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	require.NoError(t, err)
	require.Equal(t, john2, value)
}

func cleanupHashIndex() {
	os.RemoveAll(hashIndexDir)
}
