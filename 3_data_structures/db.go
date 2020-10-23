package data_structures

import "fmt"

type Person struct {
	Name    string
	Surname string
	Kids    []string
}

type KeyValueDB interface {
	Save(id ID, person interface{}) error
	Find(id ID) (Person, error)
	Delete(id ID) error
	Close() error
}

type NotFound struct {
	id ID
}

func (err NotFound) Error() string {
	return fmt.Sprintf("person not found %d", err.id)
}
