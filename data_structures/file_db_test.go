package data_structures

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
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

var paul = Person{
	Name:    "Paul",
	Surname: "Smith",
	Kids:    []string{},
}

const dir = "/tmp/ddia/filedb/d240d955-661b-4f6c-8e48-e85b8c14a9e4"

func TestErrorWhenNotFound(t *testing.T) {
	defer cleanup(dir)
	// given
	db, err := newFileDB(dir)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	assert.EqualValues(t, Person{}, value)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "not found id"))
}

func TestReadWriteValue(t *testing.T) {
	defer cleanup(dir)
	// given
	db, err := newFileDB(dir)
	require.NoError(t, err)
	err = db.Save(1, john)
	require.NoError(t, err)

	// when
	value, err := db.Find(1)

	// then
	require.NoError(t, err)
	require.Equal(t, john, value)
}

func TestReadingNewestValue(t *testing.T) {
	defer cleanup(dir)
	// given
	db, err := newFileDB(dir)
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

func cleanup(dir string) {
	os.RemoveAll(dir)
}
