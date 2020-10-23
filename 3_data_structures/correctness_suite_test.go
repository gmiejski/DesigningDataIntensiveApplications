package data_structures

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var john = Person{
	Name:    "John",
	Surname: "Smith",
	Kids:    []string{"Anabel", "Marco"},
}

var john2 = Person{
	Name:    "John",
	Surname: "Smith",
	Kids:    []string{"Anabel", "Marco", "Adelaide"},
}

func testErrorWhenNotFound(t *testing.T, db KeyValueDB) {
	// when
	value, err := db.Find(1)

	// then
	assert.EqualValues(t, Person{}, value)
	assert.Error(t, err)
	assert.IsType(t, NotFound{}, err)
}
func testReadWriteValue(t *testing.T, db KeyValueDB) {
	err := db.Save(1, john)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	require.NoError(t, err)
	require.Equal(t, john, value)
}

func testReadNewestValue(t *testing.T, db KeyValueDB) {
	// given
	err := db.Save(1, john)
	require.NoError(t, err)
	err = db.Save(1, john2)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	require.NoError(t, err)
	require.Equal(t, john2, value)
}

func testErrorDeletingNotExistingID(t *testing.T, db KeyValueDB) {
	// when
	err := db.Delete(1)

	// then
	require.Error(t, err)
	require.IsType(t, NotFound{}, err)
}

func testDeletedPersonIsNotRetrievable(t *testing.T, db KeyValueDB) {
	err := db.Save(1, john)
	require.NoError(t, err)

	// when
	err = db.Delete(1)

	// then
	require.NoError(t, err)
	_, err = db.Find(1)
	require.IsType(t, NotFound{}, err)
}

func testSavingValueAfterDeletingKey(t *testing.T, db KeyValueDB) {
	err := db.Save(1, john)
	require.NoError(t, err)
	err = db.Delete(1)
	require.NoError(t, err)

	// when
	err = db.Save(1, john)

	// then
	require.NoError(t, err)
	person, err = db.Find(1)
	require.NoError(t, err)
	require.Equal(t, john, person)
}
