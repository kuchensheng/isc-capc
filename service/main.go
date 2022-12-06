package main

import (
	"github.com/isyscore/isc-gobase/server"
	"github.com/kuchensheng/capc/infrastructure/connetor"
	"github.com/kuchensheng/capc/service/handlers"
)

func main() {
	connetor.InitMysqlConnector()
	handlers.InitAllView()
	server.Run()
}
