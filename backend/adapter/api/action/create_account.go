package action

import (
	"A-Koya/react-PTJcist/adapter/api/response"
	"A-Koya/react-PTJcist/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type CreateAccountAction struct {
	uc usecase.CreateAccountUseCase
}

func NewCreateAccountAction(uc usecase.CreateAccountUseCase) CreateAccountAction {
	return CreateAccountAction{uc: uc}
}

func (a CreateAccountAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateAccountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("can't decode user infomation")
		response.NewErrorMessage("jsonのデコードに失敗しました", http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()
	//TODO make validater
	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		log.Println(err)
		response.NewErrorMessage("SQL実行に失敗しました", http.StatusBadRequest).Send(w)
		return
	}
	response.NewSuccessMessage(output, http.StatusAccepted).Send(w)
}
