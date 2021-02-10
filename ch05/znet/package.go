package znet

type Package struct {
	DataLen   uint32
	DataIndex uint32
	Data      []byte
}

func (self *Package) GetDataLen() uint32 {
	return self.DataLen
}
func (self *Package) GetDataIndex() uint32 {
	return self.DataIndex
}
func (self *Package) GetData() []byte {
	return self.Data
}

func (self *Package) SetData(data []byte) {
	self.Data = data
}
