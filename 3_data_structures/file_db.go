package data_structures

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
)

type fileDB struct {
	file     *os.File
	filepath string
}

func (s *fileDB) Close() error {
	return s.file.Close()
}

func (s *fileDB) Save(id ID, object interface{}) error {
	jsonObject, err := Serde{}.serialize(id, object)
	if err != nil {
		return err
	}
	if _, err := s.file.WriteString(jsonObject + "\n"); err != nil {
		return err
	}
	// TODO any file sync needed?
	return nil
}

func (s *fileDB) Find(id ID) (Person, error) {
	inFile, err := os.Open(s.filepath)
	if err != nil {
		fmt.Println(err.Error() + `: ` + s.filepath)
		return Person{}, err
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	var lastReadValue *Person
	serde := Serde{}
	for scanner.Scan() {
		text := scanner.Text()

		record, err := serde.deserialize(text)
		if err != nil {
			return Person{}, err
		}
		if record.deleted {
			lastReadValue = nil
		} else {
			person := record.Person
			lastReadValue = &person
		}
	}
	if lastReadValue == nil {
		return Person{}, NotFound{id: id}
	}
	return *lastReadValue, nil
}

func (s *fileDB) Delete(id ID) error {
	newRecord := Serde{}.markDeleted(id)
	if _, err := s.file.WriteString(newRecord); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func newFileDB(dir string) (KeyValueDB, error) {
	directory := getDirectory(dir)
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return nil, err
	}

	filename := "file_db.log"
	fileFullPath := path.Join(directory, filename)
	f, err := os.OpenFile(fileFullPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &fileDB{file: f, filepath: fileFullPath}, nil
}

func getDirectory(dir string) string {
	if dir == "" {
		uid, _ := uuid.NewRandom()
		return path.Join("/tmp/ddia/filedb/", uid.String())
	} else {
		return dir
	}
}
