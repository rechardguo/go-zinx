package znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"zinx/ch08/ziface"
)

type DataPack struct {
}

func NewDataPack() ziface.IDataPack {
	return &DataPack{}
}

//body字节数    u4
//消息的index   u4
//body         body字节数个byte
//因此包头固定为8
func (self *DataPack) GetHeadLen() uint32 {
	return 8
}

func (self *DataPack) Pack(pack ziface.IMessage) ([]byte, error) {
	//构建出一个bytebuffer
	//b:=make([]byte,512)
	buf := bytes.NewBuffer([]byte{})
	//buf := new(bytes.Buffer)

	msg := pack.(*Message)

	//将消息的index写入
	if err := binary.Write(buf, binary.LittleEndian, msg.DataIndex); err != nil {
		fmt.Println("error in pack package", err)
		return nil, err
	}
	//将消息体的长度写入
	if err := binary.Write(buf, binary.LittleEndian, msg.DataLen); err != nil {
		fmt.Println("error in pack package", err)
		return nil, err
	}

	//将消息的body写入
	if err := binary.Write(buf, binary.LittleEndian, msg.Data); err != nil {
		fmt.Println("error in pack package", err)
		return nil, err
	}

	return buf.Bytes(), nil
}

func (self *DataPack) UnPack(dataBytes []byte) (ziface.IMessage, error) {
	pack := &Message{}
	buf := bytes.NewReader(dataBytes)
	//buf := bytes.NewBuffer(dataBytes)
	//消息index
	if err := binary.Read(buf, binary.LittleEndian, &pack.DataIndex); err != nil {
		fmt.Println("unpack recv data error", err)
		return nil, err
	}
	//读取的消息字节数
	if err := binary.Read(buf, binary.LittleEndian, &pack.DataLen); err != nil {
		fmt.Println("unpack recv data error", err)
		return nil, err
	}
	/*pack.SetData(make([]byte, pack.GetDataLen()))
	if err := binary.Read(buf, binary.LittleEndian, &pack.Data); err != nil {
		fmt.Println("unpack recv data error", err)
		return nil, err
	}*/

	return pack, nil
}
