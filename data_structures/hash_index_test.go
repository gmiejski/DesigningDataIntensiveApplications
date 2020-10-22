package data_structures

import (
	"testing"
)

const hashIndexDir = "/tmp/ddia/tests/hash_index/"

func newTestHashIndex() KeyValueDB {
	db, err := newFileDB(fileDBPath)
	if err != nil {
		panic(err)
	}
	return db
}

func TestErrorWhenNotFound_HashIndex(t *testing.T) {
	defer cleanup(hashIndexDir)
	// given
	db := newTestHashIndex()

	// expect
	testErrorWhenNotFound(t, db)
}

func TestReadWriteValue_HashIndex(t *testing.T) {
	defer cleanup(hashIndexDir)
	// given
	db := newTestHashIndex()

	// expect
	testReadWriteValue(t, db)
}

func TestReadingNewestValue_HashIndex(t *testing.T) {
	defer cleanup(hashIndexDir)
	// given
	db := newTestHashIndex()

	// expect
	testReadNewestValue(t, db)
}
