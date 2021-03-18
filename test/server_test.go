package test

import (
	"context"
	"errors"
	"fmt"
	"go-ethereum-development/server"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
	"testing"
)

var Server server.CrudServer

func TestMongo(t *testing.T) {
	collectionBlock := Server.MongoDBClient.Database("explorer").Collection("block")
	var filterBlock interface{}
	err := bson.Unmarshal([]byte(`{$match: {$and:[{"timestamp":{$lte:"0x60182580"}},{"timestamp":{$gte:"0x6016d400"}}]}},{$unwind: "$transactions"},{$project: {count:{$add:1}}},{$group: {_id: null, number: {$sum: "$count" }}}`), &filterBlock)
	if err != nil {
		log.Println(err.Error())
	} else {
		collectionBlock.Aggregate(context.Background(), filterBlock)
	}
}
func TestEthereum(t *testing.T) {
	var result string
	err := Server.EthereumClient.CallContext(context.Background(), &result, "eth_blockNumber")
	if err != nil {
		log.Println("RPC CALL ERROR: eth_blockNumber", err.Error())
		return
	}
	if result == "" {
		err = errors.New("SEARCH FAILED: eth_blockNumber in TestEthereum")
		return
	}
	blockNumber, err := strconv.ParseUint(result, 0, 64)
	if err != nil {
		log.Println("BLOCKNUMBER CONVERSION ERROR:", err.Error())
		return
	}
	fmt.Println(blockNumber)
}
func TestMain(m *testing.M) {
	Server.Init()
	m.Run()
}
