package znet

import "zinx/ch04/ziface"

type Request struct {
	Connection ziface.IConnection
	Date       []byte
}

func (self *Request) GetConnection() ziface.IConnection {
	return self.Connection
}

func (self *Request) GetData() []byte {
	return self.Date
}
