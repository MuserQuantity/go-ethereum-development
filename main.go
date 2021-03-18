package main

import (
	"fmt"
	"github.com/MuserQuantity/go-ethereum-development/server"
)

var Server server.CrudServer

func main() {
	Server.Init()
	fmt.Println("haha")
}
