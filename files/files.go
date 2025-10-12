package files

import (
	"fmt"
	"os"
)

/*
*
ReadFile читает содержимое файла "test.txt" и выводит его на экран.
В случае ошибки проверяет, существует ли файл, и выводит соответствующее сообщение.
*/
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

/*
*
WriteFile записывает переданное содержимое в файл с указанным именем.
Создаёт новый файл или перезаписывает существующий.
После записи файл автоматически закрывается.

Parameters:
  - content string - данные для записи в файл
  - name string - имя целевого файла
*/
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
