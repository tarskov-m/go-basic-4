// Package files
package files

import (
	"demo/password/output"
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
			output.PrintError("Файл не найден")
		} else {
			output.PrintError(err)
		}
		return nil, err
	}
	return data, nil
}

func (db JSONDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		output.PrintError(err)
		return
	}
	fmt.Println("Запись успешна")
}
