package znet

import (
	"net"
	"zinx/ch02/ziface"
)

type Connection struct {
	Uid  uint32
	Conn *net.TCPConn
}

func (self *Connection) GetConnId() uint32 {
	return self.Uid
}

func (self *Connection) GetConn() *net.TCPConn {
	return self.Conn
}

func (self *Connection) RemoteAddr() net.Addr {
	return self.Conn.RemoteAddr()
}

func (self *Connection) Start() {

}
func (self *Connection) Stop() {

}
func (self *Connection) Handler() {

}

func NewConnection(conn *net.TCPConn) ziface.IConnection {
	return &Connection{
		Uid:  1,
		Conn: conn,
	}
}
