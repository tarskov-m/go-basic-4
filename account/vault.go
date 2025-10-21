// Package account
package account

import (
	"encoding/json"
	"strings"
	"time"

	"demo/password/output"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type DB interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDB struct {
	Vault
	DB DB
}

func NewVault(db DB) *VaultWithDB {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			DB: db,
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файд data.json")
		return &VaultWithDB{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			DB: db,
		}
	}
	return &VaultWithDB{
		Vault: vault,
		DB:    db,
	}
}

func (vault *VaultWithDB) FindAccountsByURL(url string) []Account {
	var accounts []Account
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
		if isMatched {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func (vault *VaultWithDB) DeleteAccountByURL(url string) bool {
	var accounts []Account
	isDeleted := false
	for _, account := range vault.Accounts {
		isMatched := strings.Contains(account.URL, url)
		if !isMatched {
			accounts = append(accounts, account)
		} else {
			isDeleted = true
		}
	}
	vault.Accounts = accounts
	vault.save()
	return isDeleted
}

func (vault *VaultWithDB) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	} else {
		return file, nil
	}
}

func (vault *VaultWithDB) save() {
	vault.UpdatedAt = time.Now()
	data, err := vault.Vault.ToBytes()
	if err != nil {
		output.PrintError(err)
	}
	vault.DB.Write(data)
}
