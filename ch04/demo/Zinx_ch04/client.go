package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client is starting")

	conn, err := net.Dial("tcp4", "127.0.0.1:7777")
	if err != nil {
		panic(fmt.Sprintf("error in connect to server %s", err))
	}

	for {
		msg := []byte("i love you")
		//写消息到服务器端
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("write to server error", err)
			return
		}

		resp := make([]byte, 512)
		len, err := conn.Read(resp)
		if err != nil {
			fmt.Printf("error in read server response message", err)
			return
		}

		fmt.Printf("receive from server: %s \n", resp[:len])

		time.Sleep(1 * time.Second)
	}

}
