package router

import (
	"A-Koya/react-PTJcist/adapter/repository"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaMux struct {
	router *mux.Router
	db     repository.SQL
}

func newGorillaMux(db repository.SQL) *gorillaMux {
	return &gorillaMux{router: mux.NewRouter(), db: db}
}

func (g gorillaMux) Listen() {
	g.setAppHandler(g.router)
	//g.setMiddleware(g.router)

	fmt.Println("start server")
	if err := http.ListenAndServe(":8080", g.router); err != nil {
		log.Fatal("server stoped", err)
	}
}

func (g gorillaMux) setAppHandler(router *mux.Router) {
	//routing
}
