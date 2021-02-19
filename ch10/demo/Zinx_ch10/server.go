package main

import (
	"fmt"
	"zinx/ch10/ziface"
	"zinx/ch10/znet"
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

func OnConnConnected(conn ziface.IConnection) {
	fmt.Println("connId=", conn.GetConnId(), "is connected")
	resp := fmt.Sprintf("[zinx server] welcome connId=%d", conn.GetConnId())
	conn.Send(200, []byte(resp))
	conn.AddProperty("myname", "rechard")
	conn.AddProperty("age", 35)
	conn.AddProperty("aboutme", "i eager for success")
}

func OnConnClosed(conn ziface.IConnection) {
	fmt.Println("connId=", conn.GetConnId(), "is closed")
	if val, err := conn.GetProperty("myname"); err == nil {
		fmt.Println("myname=", val)
	}
	if val, err := conn.GetProperty("age"); err == nil {
		fmt.Println("age=", val)
	}
	if val, err := conn.GetProperty("aboutme"); err == nil {
		fmt.Println("aboutme=", val)
	}
}

func main() {
	server := znet.NewServer("zinx_ch10")
	server.AddRouter(1, &PingRouter{})
	server.AddRouter(2, &GreetingRouter{})

	server.OnConnConnected(OnConnConnected)
	server.OnConnClosed(OnConnClosed)

	server.Serve()
}
