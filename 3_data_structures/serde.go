package data_structures

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const tombstone = "DELETED"

type Record struct {
	Person  Person
	ID      ID
	deleted bool
}

type Serde struct {
}

func (s Serde) serialize(id ID, person interface{}) (string, error) {
	jsonObject, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return strconv.Itoa(id) + "," + string(jsonObject), nil
}

func (s Serde) deserialize(line string) (Record, error) {
	index := strings.Index(line, ",")
	if index == -1 {
		return Record{}, errors.New("cannot deserialize")
	}
	foundIDString := line[:index]
	foundID, err := strconv.Atoi(foundIDString)
	if err != nil {
		return Record{}, err
	}

	serializedObject := line[index+1 : len(line)]
	if strings.Trim(serializedObject, "\n") == tombstone {
		return Record{
			ID:      foundID,
			deleted: true,
		}, nil
	}
	var person Person
	err = json.Unmarshal([]byte(serializedObject), &person)
	if err != nil {
		return Record{}, err
	}
	return Record{
		Person:  person,
		ID:      foundID,
		deleted: false,
	}, nil
}

func (s Serde) markDeleted(id int) string {
	return fmt.Sprintf("%d,%s\n", id, tombstone)
}
