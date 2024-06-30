package infrastructure

import (
	"A-Koya/react-PTJcist/adapter/repository"
	"A-Koya/react-PTJcist/infrastructure/mysql"
	"A-Koya/react-PTJcist/infrastructure/router"
	"fmt"
	"log"
)

type config struct {
	db        repository.SQL
	webServer router.Server
}

func NewConfig() *config {
	return &config{}
}

func (c *config) Mysql() *config {
	db, err := mysql.NewMysqlHandler()
	if err != nil {
		log.Fatal(err, "Could not make a connection to the database")
	}
	fmt.Println("db connection success")
	c.db = db
	return c
}

func (c *config) WebServer() *config {
	webServer, err := router.NewServerFactory(c.db)
	if err != nil {
		log.Fatal(err, "webServer factory failed")
	}
	c.webServer = webServer
	return c
}

func (c *config) Start() {
	c.webServer.Listen()
}
