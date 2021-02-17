package main

import (
	"fmt"
	"zinx/ch05/ziface"
	"zinx/ch05/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (self *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router handler")
	fmt.Println("recv from client message index:", request.GetMessage().GetDataIndex(),
		"message:", string(request.GetMessage().GetData()))
	request.GetConnection().Send(1, []byte("ping...ping...ping \n"))
}

func main() {
	server := znet.NewServer("zinx_ch05")
	pingRouter := &PingRouter{}
	server.AddRouter(pingRouter)
	server.Serve()
}
