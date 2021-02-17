package znet

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"testing"
)

func TestPackagePackUnpack(t *testing.T) {

	//写个服务器端的连接
	//不断的读取消息
	listener, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("error in listen ", err)
		return
	}
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept in listen ", err)
			return
		}
		for {
			dp := NewDataPack()

			//先读取出消息头字节数组
			headBytes := make([]byte, dp.GetHeadLen())
			io.ReadFull(conn, headBytes)
			//消息字节数组放入解析
			p, err := dp.UnPack(headBytes)
			if err != nil {
				fmt.Println("error unpack bytes", err)
				continue
			}

			if p.GetDataLen() > 0 {
				p.SetData(make([]byte, p.GetDataLen()))
				if err := binary.Read(conn, binary.LittleEndian, p.GetData()); err != nil {
					fmt.Println("unpack recv data error", err)
					continue
				}
				println("消息index:", p.GetDataIndex(), ",消息body:", string(p.GetData()))
			}

		}

	}()

	//客户端连上服务器并发送2条粘在一起的消息
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	//构造2条消息发出去
	msg1 := []byte("how are you")
	pack1 := NewPackage(1, msg1)
	msg2 := []byte("hello world")
	pack2 := NewPackage(2, msg2)
	dp := NewDataPack()
	pack1byte, err := dp.Pack(pack1)
	if err != nil {
		fmt.Println("error in pack", err)
		return
	}
	pack2byte, err := dp.Pack(pack2)
	if err != nil {
		fmt.Println("error in pack", err)
		return
	}

	pack1byte = append(pack1byte, pack2byte...)

	conn.Write(pack1byte)
	//阻塞
	select {}
}
