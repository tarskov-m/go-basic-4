package account

import (
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts:  []Account{},
		UpdatedAt: time.Now(),
	}
}

func (vault *Vault) AddAccount(account Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	} else {
		return file, nil
	}
}
