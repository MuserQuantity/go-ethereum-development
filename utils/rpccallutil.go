package utils

import "go-ethereum-development/model/jsonrpc"

func GetBlockNumber() (blocknumber uint64, err error) {
	return
}
func GetBlockByHash(hash string) (block jsonrpc.Block, err error) {
	return
}
func GetBlockByNumber(number uint64) (block jsonrpc.Block, err error) {
	return
}
func GetTransactionByHash(hash string) (transaction jsonrpc.Transaction, err error) {
	return
}
func GetTransactionsByBlockNumber(blockNumber uint64) (transactions []jsonrpc.Transaction, err error) {
	return
}
func GetTransactionsByBlockHash(blockHash string) (transactions []jsonrpc.Transaction, err error) {
	return
}
func GetNodeList(blockHash string) (nodes []jsonrpc.NodeInfo, err error) {
	return
}
