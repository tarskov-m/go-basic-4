package files

import (
	"fmt"
	"os"
)

// функция для чтения файла.
func ReadFile() {

}

// функция для записи в файл.
func WriteFile(content string, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = file.WriteString(content)
	if err != nil {
		file.Close()
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
	file.Close()
}
