package main

import (
	"A-Koya/react-PTJcist/infrastructure"
)

func main() {
	app := infrastructure.NewConfig().Mysql().WebServer()
	app.Start()
}
