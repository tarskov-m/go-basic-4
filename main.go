package main

import (
	"fmt"
	"strings"

	"demo/password/account"
	"demo/password/encrypter"
	"demo/password/files"
	"demo/password/output"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDB){
	"1": createAccount,
	"2": findAccountByURL,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	fmt.Println("___ Менеджер паролей ___")
	err := godotenv.Load()
	if err != nil {
		output.PrintError("Ошибка при загрузке файла .env")
	}
	vault := account.NewVault(files.NewJSONDB("data.json"), *encrypter.NewEncrypter())
Menu:
	for {
		variant := promptData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по логину",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите вариант",
		)
		funcMenu := menu[variant]
		if funcMenu == nil {
			break Menu
		}
		funcMenu(vault)
	}
}

func findAccountByURL(vault *account.VaultWithDB) {
	url := promptData("Введите URL")
	accaunts := vault.FindAccounts(url, func(account account.Account, str string) bool {
		return strings.Contains(account.URL, str)
	})

	outputResult(&accaunts)
}

func findAccountByLogin(vault *account.VaultWithDB) {
	login := promptData("Введите логин")
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
	url := promptData("Введите URL")
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
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")
	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Неверный формат URL или Логин")
		return
	}
	vault.AddAccount(*myAccount)
}

func promptData(prompt ...any) string {
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
