package utils

import (
	"context"
	"errors"
	"github.com/MuserQuantity/go-ethereum-development/model/jsonrpc"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
	"strconv"
)

type RpcCall struct {
	EthereumClient *rpc.Client
}

//
func (rpcCall *RpcCall) Init(client *rpc.Client) {
	rpcCall.EthereumClient = client
}
func (rpcCall *RpcCall) GetBlockNumber() (blockNumber uint64, err error) {
	var result string
	err = rpcCall.EthereumClient.CallContext(context.Background(), &result, "eth_blockNumber")
	if err != nil {
		log.Println("RPC CALL ERROR: eth_blockNumber", err.Error())
		return
	}
	if result == "" {
		err = errors.New("SEARCH FAILED: eth_blockNumber in TestEthereum")
		return
	}
	blockNumber, err = strconv.ParseUint(result, 0, 64)
	if err != nil {
		log.Println("BLOCKNUMBER CONVERSION ERROR:", err.Error())
		return
	}
	return
}

// 获取balance
func (rpcCall *RpcCall) GetBalance(address string) (balance *big.Int, err error) {
	var result string
	err = rpcCall.EthereumClient.CallContext(context.Background(), &result, "eth_getBalance", address, "latest")
	if err != nil {
		log.Println("RPC CALL ERROR: eth_getBalance", err.Error())
		return
	}
	value, err := hexutil.DecodeBig(result)
	if err != nil {
		log.Println("RPC CALL ERROR: eth_getBalance", err.Error())
		return
	}
	balance = FromWei(value)
	return
}
func (rpcCall *RpcCall) ListAccount() (accounts []string, err error) {
	err = rpcCall.EthereumClient.CallContext(context.Background(), &accounts, "eth_accounts")
	if err != nil {
		log.Println("RPC CALL ERROR: eth_getAccounts", err.Error())
		return
	}
	return
}
func (rpcCall *RpcCall) GetBlockByHash(hash string) (block jsonrpc.Block, err error) {
	return
}
func (rpcCall *RpcCall) GetBlockByNumber(number uint64) (block jsonrpc.Block, err error) {
	return
}
func (rpcCall *RpcCall) GetTransactionByHash(hash string) (transaction jsonrpc.Transaction, err error) {
	return
}
func (rpcCall *RpcCall) GetTransactionsByBlockNumber(blockNumber uint64) (transactions []jsonrpc.Transaction, err error) {
	return
}
func (rpcCall *RpcCall) GetTransactionsByBlockHash(blockHash string) (transactions []jsonrpc.Transaction, err error) {
	return
}
func (rpcCall *RpcCall) GetNodeList(blockHash string) (nodes []jsonrpc.NodeInfo, err error) {
	return
}
