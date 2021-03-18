package jsonrpc

type Peer struct {
	Caps      []string
	Id        string
	Name      string
	Network   Network
	Protocols Protocols
}
type Network struct {
	LocalAddress  string
	RemoteAddress string
}
type Protocols struct {
	Eth Eth
}
type Eth struct {
	Difficulty int64  `json:"difficulty"`
	Head       string `json:"head"`
	Version    int    `json:"version"`
	Genesis    string `json:"genesis"`
	Network    int    `json:"network"`
}
