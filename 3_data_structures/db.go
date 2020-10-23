package data_structures

type Person struct {
	Name    string
	Surname string
	Kids    []string
}

type KeyValueDB interface {
	Save(id ID, person interface{}) error
	Find(id ID) (Person, error)
	Close() error
}
