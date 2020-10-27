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
	id, person, err := serde.deserialize(serialized)
	require.NoError(t, err)
	require.Equal(t, 1, id)
	require.Equal(t, john, *person)
}

func TestErrorDeserializingWithoutID(t *testing.T) {
	// given
	serde := Serde{}
	jsonObject := toJson(john)
	// when
	_, p, err := serde.deserialize(jsonObject)

	// then
	require.Error(t, err)
	require.Nil(t, p)
}

func TestErrorWhenCannotReadID(t *testing.T) {
	// given
	serde := Serde{}
	jsonObject := toJson(john)
	// when
	_, p, err := serde.deserialize("adosib," + jsonObject)

	// then
	require.Error(t, err)
	require.Nil(t, p)
}

func toJson(person Person) string {
	jsonObject, err := json.Marshal(person)
	if err != nil {
		panic(err.Error())
	}
	return string(jsonObject)
}
