package jsonrpc

type BlockWithDetails struct {
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
	Transactions     []Transaction
	TransactionsRoot string
	Uncles           []string
}
