package jsonrpc

type NodeInfo struct {
	Enode      string
	Id         string
	Ip         string
	ListenAddr string
	Name       string
	Ports      Ports
	Protocols  Protocols
}
type Ports struct {
	Discovery int
	Listener  int
}
