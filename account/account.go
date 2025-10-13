// Package account
package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

// letterRunes содержит символы для генерации случайных паролей
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-*!")

// Account представляет учетную запись с логином, паролем и URL
type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// OutputPassword выводит логин в цвете циан
func (acc *Account) OutputPassword() {
	color.Cyan(acc.Login)
}

// generatePassword генерирует случайный пароль заданной длины
// использует символы из letterRunes
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

// NewAccount создает новую учетную запись с временными метками
// Параметры:
//   - login: строка логина (не может быть пустой)
//   - password: строка пароля (если пустая, будет сгенерирован)
//   - urlString: строка URL (должна быть валидным URI)
//
// Возвращает:
//   - *Account: указатель на созданную учетную запись
//   - error: ошибку, если входные данные некорректны
func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Url:       urlString,
		Login:     login,
		Password:  password,
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
