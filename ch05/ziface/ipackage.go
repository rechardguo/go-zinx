package ziface

type IPackage interface {
	GetDataLen() uint32
	GetDataIndex() uint32
	GetData() []byte
	SetData([]byte)
}
