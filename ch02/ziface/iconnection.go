package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	//放回Connection的id
	GetConnId() uint32
	//返回net.Conn
	GetConn() *net.TCPConn
	Handler()
	RemoteAddr() net.Addr
}
