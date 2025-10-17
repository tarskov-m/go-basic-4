package main

import (
	"fmt"

	"demo/password/account"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___ Менеджер паролей ___")
	vault := account.NewVault()
Menu:
	for {
		variant := GetMenu()
		switch variant {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
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

func findAccount(vault *account.Vault) {
	url := promptData("Введите URL")
	accaunts := vault.FindAccountsByURL(url)
	if len(accaunts) == 0 {
		color.Red("Не найдено")
	}
	for _, account := range accaunts {
		account.Output()
	}
}

func deleteAccount(vault *account.Vault) {
	url := promptData("Введите URL")
	if vault.DeleteAccountByURL(url) {
		color.Green("Удалено")
	} else {
		color.Red("Не найдено")
	}
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
func createAccount(vault *account.Vault) {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или Логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
