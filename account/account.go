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
	login    string
	password string
	url      string
}

// accountWithTimeStamp - расширенная структура учетной записи с временными
// метками содержит встроенный Account и дополнительные поля времени создания
// обновления
type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

// OutputPassword выводит логин в цвете циан
func (acc *Account) OutputPassword() {
	color.Cyan(acc.login)
}

// generatePassword генерирует случайный пароль заданной длины
// использует символы из letterRunes
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// NewAccountWithTimeStamp создает новую учетную запись с временными метками
// Параметры:
//   - login: строка логина (не может быть пустой)
//   - password: строка пароля (если пустая, будет сгенерирован)
//   - urlString: строка URL (должна быть валидным URI)
//
// Возвращает:
//   - *accountWithTimeStamp: указатель на созданную учетную запись
//   - error: ошибку, если входные данные некорректны
func NewAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			url:      urlString,
			login:    login,
			password: password,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
