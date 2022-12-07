package main

import (
	"github.com/isyscore/isc-gobase/server"
	"github.com/kuchensheng/capc/service/handlers"
)

func main() {
	handlers.InitAllView()
	server.Run()
}
