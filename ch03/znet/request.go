package znet

import "zinx/ch03/ziface"

type Request struct {
	Connection ziface.IConnection
	Data       []byte
}

func (self *Request) GetConnection() ziface.IConnection {
	return self.Connection
}

func (self *Request) GetData() []byte {
	return self.Data
}
