package data_structures

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"strconv"
)

type ID = int

type hashIndex struct {
	file             *os.File
	filepath         string
	index            map[ID]int64
	nextIndexPointer int64
	readPointer      *os.File
}

func (h *hashIndex) Save(id ID, person interface{}) error {
	record, err := Serde{}.serialize(id, person)
	if err != nil {
		return err
	}
	nextLine := record + "\n"
	if _, err := h.file.WriteString(nextLine); err != nil {
		log.Println(err)
		return err
	}
	h.index[id] = h.nextIndexPointer
	h.nextIndexPointer += int64(len(nextLine))
	return nil
}

func (h *hashIndex) Find(id ID) (Person, error) {
	if h.readPointer == nil {
		pointer, err := os.OpenFile(h.filepath, os.O_RDONLY|os.O_CREATE, 0644)
		if err != nil {
			return Person{}, err
		}
		h.readPointer = pointer
	}
	text, err := h.findIndexLine(id)
	if err != nil {
		return Person{}, NotFound{id: id}
	}
	return h.parseLine(id, text)

}

func (h *hashIndex) findIndexLine(id ID) (string, error) {
	pointer, ok := h.index[id]
	if !ok {
		return "", fmt.Errorf("not found: %d", id)
	}
	_, err := h.readPointer.Seek(pointer, 0)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(h.readPointer)
	scanner.Scan()
	return scanner.Text(), nil
}

func (h *hashIndex) parseLine(id ID, text string) (Person, error) {
	foundID, person, err := Serde{}.deserialize(text)
	if err != nil {
		return Person{}, nil
	}
	if foundID != id {
		return Person{}, fmt.Errorf("wrong id in the line: %s", text)
	}
	return *person, nil
}

func (h *hashIndex) Delete(id ID) error {
	if _, ok := h.index[id]; !ok {
		return NotFound{id: id}
	}
	newRecord := strconv.Itoa(id) + ",DELETED\n"
	if _, err := h.file.WriteString(newRecord); err != nil {
		log.Println(err)
		return err
	}
	delete(h.index, id)
	h.nextIndexPointer += int64(len(newRecord))
	return nil
}

func getDirectoryHashIndex(dir string) string {
	if dir == "" {
		uid, _ := uuid.NewRandom()
		return path.Join("/tmp/ddia/hash_index/", uid.String())
	} else {
		return dir
	}
}

func newHashIndex(dir string) (KeyValueDB, error) {
	// TODO open existing DB from a file
	directory := getDirectoryHashIndex(dir)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return nil, err
	}

	filename := "hash_index.log"
	fileFullPath := path.Join(directory, filename)
	f, err := os.OpenFile(fileFullPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &hashIndex{
		file:             f,
		filepath:         fileFullPath,
		index:            make(map[ID]int64),
		nextIndexPointer: 0,
	}, nil
}

func (h *hashIndex) Close() error {
	h.readPointer.Close()
	h.file.Close()
	return nil // TODO erroring
}
