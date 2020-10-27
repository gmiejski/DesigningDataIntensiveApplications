package data_structures

import (
	"testing"
)

const hashIndexDir = "/tmp/ddia/tests/hash_index/"

func newTestHashIndex() KeyValueDB {
	db, err := newHashIndex(hashIndexDir)
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

func TestErrorDeletingNotExistingID_HashIndex(t *testing.T) {
	defer cleanup(hashIndexDir)
	// given
	db := newTestHashIndex()

	// expect
	testErrorDeletingNotExistingID(t, db)
}

func TestDeletedPersonIsNotRetrievable_HashIndex(t *testing.T) {
	defer cleanup(hashIndexDir)
	// given
	db := newTestHashIndex()

	// expect
	testDeletedPersonIsNotRetrievable(t, db)
}

func TestSavingValueAfterDeletingKey_HashIndex(t *testing.T) {
	defer cleanup(hashIndexDir)
	// given
	db := newTestHashIndex()

	// expect
	testSavingValueAfterDeletingKey(t, db)
}
