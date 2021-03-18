package main

import (
	"fmt"
	"go-ethereum-development/server"
)

var Server server.CrudServer

func main() {
	Server.Init()
	fmt.Println("haha")
}
