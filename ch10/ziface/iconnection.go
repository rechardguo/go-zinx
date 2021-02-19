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

	//添加连接属性
	AddProperty(string, interface{})
	//移除连接属性
	RemoveProperty(string)
	//获取连接属性
	GetProperty(string) (interface{}, error)
}

type HandleFunc func(*net.TCPConn, []byte, int) error
