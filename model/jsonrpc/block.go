package jsonrpc

import "C"
import (
	"context"
	"go-ethereum-development/server"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Block struct {
	Miner            string
	ExtraData        string
	Difficulty       string
	GasLimit         string
	GasUsed          string
	Hash             string
	LogsBloom        string
	MixHash          string
	Nonce            string
	Number           string
	ParentHash       string
	ReceiptsRoot     string
	Sha3Uncles       string
	Size             string
	StateRoot        string
	Timestamp        string
	TotalDifficulty  string
	Transactions     []string
	TransactionsRoot string
	Uncles           []string
}

// Create
func (block *Block) Insert(crudServer server.CrudServer) (err error) {
	collection := crudServer.MongoDBClient.Database("explorer").Collection("block")
	var insertResult *mongo.InsertOneResult
	insertResult, err = collection.InsertOne(context.TODO(), *block)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Inserted a single document in collection block: ", insertResult.InsertedID)
	return
}

// Query
func (block *Block) QueryByNumber(crudServer server.CrudServer) (err error) {
	collection := crudServer.MongoDBClient.Database("explorer").Collection("block")
	filter := bson.D{{"number", block.Number}}
	err = collection.FindOne(context.TODO(), filter).Decode(&block)
	if err != nil {
		log.Println(err)
		return
	}
	return
}

// Update
func (block *Block) Update() (err error) {
	return
}

// Delete
func (block *Block) Delete() (err error) {
	return
}
