package ziface

type IMessage interface {
	GetDataLen() uint32
	GetDataIndex() uint32
	GetData() []byte
	SetData([]byte)
}
