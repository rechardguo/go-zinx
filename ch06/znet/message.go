package znet

import "zinx/ch06/ziface"

type Message struct {
	DataIndex uint32
	DataLen   uint32
	Data      []byte
}

func NewPackage(msgId uint32, data []byte) ziface.IMessage {
	//buf:=bytes.NewBuffer([]byte{})
	p := &Message{
		DataIndex: msgId,
		DataLen:   uint32(len(data)),
		Data:      data,
	}
	return p
}

func (self *Message) GetDataLen() uint32 {
	return self.DataLen
}
func (self *Message) GetMsgId() uint32 {
	return self.DataIndex
}
func (self *Message) GetData() []byte {
	return self.Data
}

func (self *Message) SetData(data []byte) {
	self.Data = data
}
