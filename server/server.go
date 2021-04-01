package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/golobby/config"
	"github.com/golobby/config/feeder"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type CrudServer struct {
	RpcClient     *rpc.Client
	EthClient     *ethclient.Client
	MongoDBClient *mongo.Client
	MongoDBName   string
}

func (crudServer *CrudServer) Init() {
	// 连接全节点
	crudServer.InitEthClient()
	crudServer.InitRpcClient()
	// 连接数据库
	crudServer.InitMongoDB()
}

// 初始化mongodb连接
func (crudServer *CrudServer) InitMongoDB() {
	// 由文件初始化数据
	c, err := config.New(config.Options{
		Feeder: feeder.Json{Path: "config.json"},
		Env:    ".env",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	host, err1 := c.GetString("mongodb.host")
	port, err2 := c.GetString("mongodb.port")
	user, err3 := c.GetString("mongodb.user")
	pass, err4 := c.GetString("mongodb.pass")
	name, err5 := c.GetString("mongodb.dbname")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		log.Fatal(err1, err2, err3, err4, err5)
	}
	var mongodbLocalUrl string
	if user == "" || pass == "" {
		mongodbLocalUrl = "mongodb://" + host + ":" + port + "/" + name
	} else {
		mongodbLocalUrl = "mongodb://" + user + ":" + pass + "@" + host + ":" + port + "/" + name
	}
	// 连接MongoDB
	// 设置客户端连接配置
	mongoDBClientOptions := options.Client().ApplyURI(mongodbLocalUrl)
	// 连接到MongoDB
	crudServer.MongoDBClient, err = mongo.Connect(context.TODO(), mongoDBClientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to the MongoDB client: %v", err)
	}
	crudServer.MongoDBName = name
	// 检查连接
	err = crudServer.MongoDBClient.Ping(context.TODO(), nil)
	if err != nil {
		err = errors.New("Failed to connect to the MongoDB client: %v" + err.Error())
	}
	log.Println("Connect to MongoDB successfully!")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Connect to MongoDB successfully!")
}

// 初始化mysql连接

// 初始化ethereum连接
func (crudServer *CrudServer) InitRpcClient() {
	// 由文件初始化数据
	c, err := config.New(config.Options{
		Feeder: feeder.Json{Path: "config.json"},
		Env:    ".env",
	})
	ethereumhost, err1 := c.GetString("ethereum.host")
	ethereumport, err2 := c.GetString("ethereum.port")
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}
	url := "http://" + ethereumhost + ":" + ethereumport
	// 连接全节点
	crudServer.RpcClient, err = rpc.Dial(url)
	// 检查连接
	if err != nil {
		err = errors.New("Failed to connect to the RPC client: %v" + err.Error())
		return
	}
	log.Println("Connect to RPC Client successfully!")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Connect to RPC Client successfully!")
	return
}
func (crudServer *CrudServer) InitEthClient() {
	// 由文件初始化数据
	c, err := config.New(config.Options{
		Feeder: feeder.Json{Path: "config.json"},
		Env:    ".env",
	})
	ethereumhost, err1 := c.GetString("ethereum.host")
	ethereumport, err2 := c.GetString("ethereum.port")
	if err1 != nil || err2 != nil {
		log.Fatal(err1, err2)
	}
	url := "http://" + ethereumhost + ":" + ethereumport
	// 连接全节点
	crudServer.EthClient, err = ethclient.Dial(url)
	// 检查连接
	if err != nil {
		err = errors.New("Failed to connect to the Ethereum client: %v" + err.Error())
		return
	}
	log.Println("Connect to Ethereum Client successfully!")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "Connect to Ethereum Client successfully!")
	return
}
