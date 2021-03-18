package main

import (
	"fmt"
	"github.com/MuserQuantity/go-ethereum-development/server"
	"github.com/MuserQuantity/go-ethereum-development/utils"
	"log"
)

var Server server.CrudServer

func main() {
	Server.InitRpcClient()
	var rpcCall utils.RpcCall
	rpcCall.Init(Server.RpcClient)
	blockNumber, err := rpcCall.GetBlockNumber()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(blockNumber)
}
