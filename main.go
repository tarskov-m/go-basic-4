package main

import (
	"fmt"
	"strings"

	"demo/password/account"
	"demo/password/files"
	"demo/password/output"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("___ Менеджер паролей ___")
	vault := account.NewVault(files.NewJSONDB("data.json"))
Menu:
	for {
		variant := promptData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		})
		funcMenu := menu[variant]
		if funcMenu == nil {
			break Menu
		}
		funcMenu(vault)
		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}

func findAccountByURL(vault *account.VaultWithDB) {
	url := promptData([]string{"Введите URL"})
	accaunts := vault.FindAccounts(url, func(account account.Account, str string) bool {
		return strings.Contains(account.URL, str)
	})

	outputResult(&accaunts)
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData([]string{"Введите логин"})
	accaunts := vault.FindAccounts(login, func(account account.Account, str string) bool {
		return strings.Contains(account.Login, str)
	})

	outputResult(&accaunts)
}

func outputResult(accaunts *[]account.Account) {
	if len(*accaunts) == 0 {
		output.PrintError("Не найдено")
	}
	for _, account := range *accaunts {
		account.Output()
	}
}

func deleteAccount(vault *account.VaultWithDB) {
	url := promptData([]string{"Введите URL"})
	if vault.DeleteAccountByURL(url) {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
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
func createAccount(vault *account.VaultWithDB) {
	login := promptData([]string{"Введите логин"})
	password := promptData([]string{"Введите пароль"})
	url := promptData([]string{"Введите URL"})
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData[T any](prompt []T) string {
	for i, line := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", line)
		} else {
			fmt.Println(line)
		}
	}

	var res string
	fmt.Scanln(&res)
	return res
}
