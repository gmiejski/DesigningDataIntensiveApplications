package data_structures

import (
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

func TestReadingNewestValue_FileDb(t *testing.T) {
	defer cleanup(fileDBPath)
	// given
	db := newTestFileDB()

	// expect
	testReadNewestValue(t, db)
}
func cleanup(dir string) {
	os.RemoveAll(dir)
}
