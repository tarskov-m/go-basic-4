package main

import (
	"fmt"

	"demo/password/account"
	"demo/password/files"
)

func main() {
	fmt.Println("___ Менеджер паролей ___")
Menu:
	for {
		variant := GetMenu()
		switch variant {
		case 1:
			createAccount()
		case 2:
			findAccount()
		case 3:
			deleteAccount()
		default:
			break Menu
		}
	}
}

func GetMenu() int {
	fmt.Println("Выберите вариант:")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	var variant int
	fmt.Scan(&variant)
	return variant
}

func findAccount() {
}

func deleteAccount() {
}

// Создает новый аккаунт на основе вводных данных пользователя
//
// Функция последовательно запрашивает у пользователя логин, пароль и URL,
// создает объект Account, выполняет его сериализацию в JSON и сохраняет
// результат в файл "data.json".
//
// - Использует функцию promptData для получения входных данных
// - Проверяет корректность формата URL и логина через метод NewAccount
// - В случае ошибки выводит диагностическое сообщение и завершает работу
// - Результирующий JSON записывается в файл с помощью метода WriteFile
func createAccount() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	file, err := myAccount.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
