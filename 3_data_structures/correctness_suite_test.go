package data_structures

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
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
	assert.True(t, strings.Contains(err.Error(), "not found id"))
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
