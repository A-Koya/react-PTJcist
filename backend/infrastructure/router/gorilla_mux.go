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
	g.router.Use(corsMiddleware)

	fmt.Println("start server")
	if err := http.ListenAndServe(":8080", g.router); err != nil {
		log.Fatal("server stoped", err)
	}
}

func (g gorillaMux) setAppHandler(router *mux.Router) {
	//routing
	router.Handle("/account/create", g.buildCreateAccountAction()).Methods(http.MethodPost)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, application/json")

		// Preflightリクエストに対応するための設定
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
