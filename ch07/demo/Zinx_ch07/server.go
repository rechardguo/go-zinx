package main

import (
	"fmt"
	"zinx/ch07/ziface"
	"zinx/ch07/znet"
)

type GreetingRouter struct {
	znet.BaseRouter
}

func (self *GreetingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call GreetingRouter handler")
	fmt.Println("recv from client message Id:", request.GetMessage().GetMsgId(),
		"message:", string(request.GetMessage().GetData()))
	request.GetConnection().Send(200, []byte("welcome to zinx server \n"))
}

type PingRouter struct {
	znet.BaseRouter
}

func (self *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call PingRouter handler")
	fmt.Println("recv from client message Id:", request.GetMessage().GetMsgId(),
		"message:", string(request.GetMessage().GetData()))
	request.GetConnection().Send(201, []byte("pong..pong...pong \n"))
}

func main() {
	server := znet.NewServer("zinx_ch07")
	server.AddRouter(1, &PingRouter{})
	server.AddRouter(2, &GreetingRouter{})
	server.Serve()
}
