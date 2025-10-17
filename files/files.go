// Package files
package files

import (
	"fmt"
	"os"
)

type JSONDB struct {
	filename string
}

func NewJSONDB(name string) *JSONDB {
	return &JSONDB{
		filename: name,
	}
}

func (db JSONDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		if err == os.ErrNotExist {
			fmt.Println("Файл не найден")
		} else {
			fmt.Println(err)
		}
		return nil, err
	}
	return data, nil
}

func (db JSONDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
