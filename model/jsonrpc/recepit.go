package jsonrpc

type Recepit struct {
	TransactionHash   string
	TransactionIndex  string
	BlockNumber       string
	BlockHash         string
	CumulativeGasUsed string
	GasUsed           string
	ContractAddress   string
	Logs              []interface{}
	LogsBloom         string
	Status            string
}
