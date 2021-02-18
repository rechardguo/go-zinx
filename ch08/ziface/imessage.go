package ziface

type IMessage interface {
	GetDataLen() uint32
	GetMsgId() uint32
	GetData() []byte
	SetData([]byte)
}
