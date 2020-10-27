package data_structures

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

func (s Serde) deserialize(line string) (ID, *Person, error) {
	index := strings.Index(line, ",")
	if index == -1 {
		return 0, nil, errors.New("cannot deserialize")
	}
	foundIDString := line[:index]
	foundID, err := strconv.Atoi(foundIDString)
	if err != nil {
		return 0, nil, err
	}

	objectJson := line[index+1 : len(line)]
	var person Person
	err = json.Unmarshal([]byte(objectJson), &person)
	if err != nil {
		return 0, nil, err
	}
	return foundID, &person, nil
}
