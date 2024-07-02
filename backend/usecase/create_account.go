package usecase

import (
	"A-Koya/react-PTJcist/domain"
	"context"
)

type (
	createAccountInteractor struct {
		repo      domain.AccountRepository
		presenter CreateAccountPresenter
	}

	CreateAccountUseCase interface {
		Execute(context.Context, CreateAccountInput) (CreateAccountOutput, error)
	}

	CreateAccountPresenter interface {
		Output(domain.Account) CreateAccountOutput
	}

	CreateAccountInput struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}

	CreateAccountOutput struct {
		ID       string `json:"id"`
		Password string `json:"password"`
	}
)

func NewCreateAccountInteractor(repo domain.AccountRepository, presenter CreateAccountPresenter) CreateAccountUseCase {
	return createAccountInteractor{
		repo:      repo,
		presenter: presenter,
	}
}

func (a createAccountInteractor) Execute(ctx context.Context, input CreateAccountInput) (CreateAccountOutput, error) {
	account := domain.NewAccount(input.ID, input.Password)
	account, err := a.repo.Create(ctx, account)
	if err != nil {
		return a.presenter.Output(domain.Account{}), err
	}
	return a.presenter.Output(account), nil
}
