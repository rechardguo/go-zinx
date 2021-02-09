package main

import (
	"zinx/ch03/ziface"
	"zinx/ch03/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (self *PingRouter) PreHandle(request ziface.IRequest) {
	request.GetConnection().GetConn().Write([]byte("PreHandle of PingRouter \n"))
}
func (self *PingRouter) Handle(request ziface.IRequest) {
	request.GetConnection().GetConn().Write([]byte("ping...ping...ping \n"))
}
func (self *PingRouter) PostHandle(request ziface.IRequest) {
	request.GetConnection().GetConn().Write([]byte("PostHandle of PingRouter \n"))
}

func main() {
	server := znet.NewServer("zinx_ch03")
	pingRouter := &PingRouter{}
	server.AddRouter(pingRouter)
	server.Serve()
}
