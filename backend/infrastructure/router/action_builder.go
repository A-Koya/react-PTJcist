package router

import (
	"A-Koya/react-PTJcist/adapter/api/action"
	"A-Koya/react-PTJcist/adapter/presenter"
	"A-Koya/react-PTJcist/adapter/repository"
	"A-Koya/react-PTJcist/usecase"
	"net/http"
)

func (g gorillaMux) buildCreateAccountAction() http.HandlerFunc {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		var (
			uc  = usecase.NewCreateAccountInteractor(repository.NewAccountSQL(g.db), presenter.NewCreateAccountInteractor())
			act = action.NewCreateAccountAction(uc)
		)
		act.Execute(w, r)
	}
	return handler
}
