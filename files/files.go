package files

import (
	"fmt"
	"os"
)

// функция для чтения файла.
func ReadFile() {
	data, err := os.ReadFile("test.txt")
	if err != nil {
		if err == os.ErrNotExist {
			fmt.Println("Файл не найден")
		} else {
			fmt.Println(err)
		}
		return
	}
	fmt.Println(string(data))
}

// функция для записи в файл.
func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}
