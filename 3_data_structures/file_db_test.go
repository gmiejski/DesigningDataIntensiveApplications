package data_structures

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

const fileDBPath = "/tmp/ddia/tests/fileDB/"

func newTestFileDB() KeyValueDB {
	db, err := newFileDB(fileDBPath)
	if err != nil {
		panic(err)
	}
	return db
}

func TestErrorWhenNotFound_FileDb(t *testing.T) {
	defer cleanup(fileDBPath)
	// given
	db := newTestFileDB()

	// expect
	testErrorWhenNotFound(t, db)
}

func TestReadWriteValue_FileDb(t *testing.T) {
	defer cleanup(fileDBPath)
	// given
	db := newTestFileDB()

	// expect
	testReadWriteValue(t, db)
}

func TestNoErrorWhenDeletingNotExistingKey(t *testing.T) {
	defer cleanup(fileDBPath)
	// given
	db := newTestFileDB()

	// when
	err := db.Delete(1)

	// then
	require.NoError(t, err)
}

func TestDeletedPersonIsNotRetrievable_FileDb(t *testing.T) {
	defer cleanup(fileDBPath)
	// given
	db := newTestHashIndex()

	// expect
	testDeletedPersonIsNotRetrievable(t, db)
}

func TestSavingValueAfterDeletingKey_FileDb(t *testing.T) {
	defer cleanup(fileDBPath)
	// given
	db := newTestHashIndex()

	// expect
	testSavingValueAfterDeletingKey(t, db)
}

func cleanup(dir string) {
	os.RemoveAll(dir)
}
