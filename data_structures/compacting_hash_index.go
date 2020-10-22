package data_structures

type compactingHashIndex struct {
	hashIndex
}

func (c *compactingHashIndex) Save(id ID, person interface{}) error {
	panic("implement me")
}

func (c *compactingHashIndex) Find(id ID) (Person, error) {
	panic("implement me")
}

func (c *compactingHashIndex) Close() error {
	return c.hashIndex.Close()
}

func newCompactingHashIndex(dir string) (KeyValueDB, error) {
	h, err := newHashIndex(dir)
	if err != nil {
		return nil, err
	}
	db := h.(*hashIndex)
	return &compactingHashIndex{hashIndex: *db}, nil
}
