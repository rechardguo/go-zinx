package ziface

type IDataPack interface {
	GetHeadLen() uint32
	Pack(IPackage) ([]byte, error)
	UnPack([]byte) (IPackage, error)
}
