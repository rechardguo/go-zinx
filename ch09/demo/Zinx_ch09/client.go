package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
	"zinx/ch09/znet"
)

func main() {
	fmt.Println("client0 is starting")

	conn, err := net.Dial("tcp4", "127.0.0.1:7777")
	if err != nil {
		panic(fmt.Sprintf("error in connect to server %s", err))
	}

	for {
		dp := znet.NewDataPack()
		msgByte, err := dp.Pack(znet.NewPackage(1, []byte("ping...ping...ping")))
		if err != nil {
			fmt.Println("error in pack message", err)
			return
		}
		//写消息到服务器端
		_, err = conn.Write(msgByte)
		if err != nil {
			fmt.Println("write to server error", err)
			return
		}

		//从服务器端读出消息
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
			println("recv server response , 消息index:", p.GetMsgId(), ",消息body:", string(p.GetData()))
		}
		time.Sleep(1 * time.Second)
	}

}
