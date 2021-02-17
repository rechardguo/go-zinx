package ziface

type IDataPack interface {
	GetHeadLen() uint32
	Pack(IMessage) ([]byte, error)
	//byte只是消息头的长度
	UnPack([]byte) (IMessage, error)
}
