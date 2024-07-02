package repository

import (
	"A-Koya/react-PTJcist/domain"
	"context"
	"fmt"
)

type AccountSQL struct {
	db SQL
}

func NewAccountSQL(db SQL) AccountSQL {
	return AccountSQL{db: db}
}

func (a AccountSQL) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	query := `
		INSERT INTO users (id,password)
		VALUES (?,?);
	`
	if err := a.db.ExecuteContext(ctx, query, account.ID, account.Password); err != nil {
		return domain.Account{}, fmt.Errorf("err:%e", err)
	}
	return account, nil
}
