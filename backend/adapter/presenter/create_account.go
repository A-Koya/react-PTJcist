package presenter

import (
	"A-Koya/react-PTJcist/domain"
	"A-Koya/react-PTJcist/usecase"
)

type createAccountPresenter struct{}

func NewCreateAccountInteractor() createAccountPresenter {
	return createAccountPresenter{}
}

func (a createAccountPresenter) Output(account domain.Account) usecase.CreateAccountOutput {
	return usecase.CreateAccountOutput{
		ID:       account.ID,
		Password: account.Password,
	}
}
