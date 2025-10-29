// Package encrypter
package encrypter

import "os"

type Encrypter struct {
	Key string
}

func NewEncrypter() *Encrypter {
	key := os.Getenv("KEY")
	if key == "" {
		panic("Не передан параметр KEY в переменные окружения")
	}
	return &Encrypter{
		Key: key,
	}
}

func (e *Encrypter) encrypt(plainStr string) string {
	return ""
}

func (e *Encrypter) decrypt(encryptedStr string) string {
	return ""
}
