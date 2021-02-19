package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	//返回Connection的id
	GetConnId() uint32
	//返回net.Conn
	GetConn() *net.TCPConn

	RemoteAddr() net.Addr

	Send(uint32, []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
