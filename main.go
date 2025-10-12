package main

import (
	"fmt"

	"demo/password/account"
	"demo/password/files"
)

/*
*
main - основная функция программы, которая:
1. Записывает тестовый текст в файл
2. Запрашивает у пользователя логин, пароль и URL
3. Создаёт учётную запись с временной меткой
4. Обрабатывает возможные ошибки ввода
5. Выводит сгенерированный пароль и информацию об аккаунте
*/
func main() {
	files.WriteFile("Привет, я файл!", "test.txt")
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	myAccount.OutputPassword()
	fmt.Println(myAccount)
}

/*
*
promptData - вспомогательная функция для получения данных от пользователя
@param prompt - текст приглашения к вводу
@return string - введённое значение
*/
func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
