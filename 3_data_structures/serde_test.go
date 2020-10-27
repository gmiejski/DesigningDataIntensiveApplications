package data_structures

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSerializeDeserializeValue(t *testing.T) {
	// given
	serde := Serde{}
	// when
	serialized, err := serde.serialize(1, john)

	// then
	require.NoError(t, err)

	// when
	record, err := serde.deserialize(serialized)

	require.NoError(t, err)
	require.Equal(t, 1, record.ID)
	require.Equal(t, john, record.Person)
}

func TestErrorDeserializingWithoutID(t *testing.T) {
	// given
	serde := Serde{}
	jsonObject := toJson(john)
	// when
	record, err := serde.deserialize(jsonObject)

	// then
	require.Error(t, err)
	require.EqualValues(t, Record{}, record)
}

func TestErrorWhenCannotReadID(t *testing.T) {
	// given
	serde := Serde{}
	jsonObject := toJson(john)
	// when
	record, err := serde.deserialize("adosib," + jsonObject)

	// then
	require.Error(t, err)
	require.EqualValues(t, Record{}, record)
}

func TestDeletedRecord(t *testing.T) {
	// given
	serde := Serde{}
	deleted := serde.markDeleted(1)

	// when
	record, err := serde.deserialize(deleted)

	// then
	require.NoError(t, err)
	require.Equal(t, Person{}, record.Person)
	require.True(t, record.deleted)
}

func toJson(person Person) string {
	jsonObject, err := json.Marshal(person)
	if err != nil {
		panic(err.Error())
	}
	return string(jsonObject)
}
