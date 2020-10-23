package data_structures

const compactingTestPath = "/tmp/ddia/tests/compactingHashIndex/"

func newTestCompactingIndex() KeyValueDB {
	index, err := newCompactingHashIndex(compactingTestPath)
	if err != nil {
		panic(err.Error())
	}
	return index
}

//func TestErrorWhenNotFound_CompactingHashIndex(t *testing.T) {
//	defer cleanup(compactingTestPath)
//	// given
//	db := newTestCompactingIndex()
//
//	// expect
//	testErrorWhenNotFound(t, db)
//}
//
//func TestReadWriteValue_CompactingHashIndex(t *testing.T) {
//	defer cleanup(compactingTestPath)
//	// given
//	db := newTestCompactingIndex()
//
//	// expect
//	testReadWriteValue(t, db)
//}
//
//func TestReadingNewestValue_CompactingHashIndex(t *testing.T) {
//	defer cleanup(compactingTestPath)
//	// given
//	db := newTestCompactingIndex()
//
//	// expect
//	testReadNewestValue(t, db)
//}
