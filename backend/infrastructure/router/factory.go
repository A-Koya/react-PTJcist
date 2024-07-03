package router

import "A-Koya/react-PTJcist/adapter/repository"

type Server interface {
	Listen()
}

func NewServerFactory(db repository.SQL) (Server, error) {
	return newGorillaMux(db), nil
}
