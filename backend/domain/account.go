package domain

import "context"

type (
	AccountRepository interface {
		Create(context.Context, Account) (Account, error)
	}

	Account struct {
		ID       string
		Password string
	}
)

func NewAccount(id, password string) Account {
	return Account{
		ID:       id,
		Password: password,
	}
}
