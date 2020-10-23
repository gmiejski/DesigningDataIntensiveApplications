package data_structures

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

const TOMBSTONE = "DELETED"

type fileDB struct {
	file     *os.File
	filepath string
}

func (s *fileDB) Close() error {
	return s.file.Close()
}

func (s *fileDB) Save(id ID, object interface{}) error {
	jsonObject, err := json.Marshal(object)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, err := s.file.WriteString(strconv.Itoa(id) + "," + string(jsonObject) + "\n"); err != nil {
		log.Println(err)
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
	for scanner.Scan() {
		text := scanner.Text()

		index := strings.Index(text, ",")
		if index == -1 {
			return Person{}, errors.New("cannot find")
		}
		foundIDString := text[:index]
		foundID, err := strconv.Atoi(foundIDString)
		if err != nil {
			return Person{}, err
		}
		if foundID != id {
			return Person{}, fmt.Errorf("wrong id in the line: %s", foundIDString)
		}
		objectJson := text[index+1 : len(text)]
		if isDeleted(objectJson) {
			lastReadValue = nil
		} else {
			var person Person
			err = json.Unmarshal([]byte(objectJson), &person)
			if err != nil {
				return Person{}, err
			}
			lastReadValue = &person
		}
	}
	if lastReadValue == nil {
		return Person{}, NotFound{id: id}
	}
	return *lastReadValue, nil
}

func isDeleted(objectJson string) bool {
	return objectJson == TOMBSTONE
}

func (s *fileDB) Delete(id ID) error {
	newRecord := strconv.Itoa(id) + "," + TOMBSTONE + "\n"
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
