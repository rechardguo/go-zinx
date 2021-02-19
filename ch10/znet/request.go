package znet

import "zinx/ch10/ziface"

type Request struct {
	Connection ziface.IConnection
	Message    ziface.IMessage
}

func (self *Request) GetConnection() ziface.IConnection {
	return self.Connection
}

func (self *Request) GetMessage() ziface.IMessage {
	return self.Message
}
