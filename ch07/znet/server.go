package znet

import (
	"fmt"
	"net"
	"zinx/ch07/utils"
	. "zinx/ch07/ziface"
)

type Server struct {
	name       string
	IP         string
	Port       uint32
	TCPVersion string
	MsgHandler IMsgHandler
}

func (self *Server) AddRouter(msgId uint32, router IRouter) {
	self.MsgHandler.AddRouter(msgId, router)
}

func (self *Server) Start() {
	fmt.Printf("Server %s  version : %s at Ip :%s, Port: %d  is starting \n", self.name, utils.Config.Version, self.IP, self.Port)
	fmt.Printf("maxConnection : %d, maxPackageSize : %d \n", utils.Config.MaxConn, utils.Config.MaxPackageSize)

	//获取一个TCP的addr
	addr, err := net.ResolveTCPAddr(self.TCPVersion, fmt.Sprintf("%s:%d", self.IP, self.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
	}
	listener, err := net.ListenTCP(self.TCPVersion, addr)
	if err != nil {
		fmt.Println("listen error", err)
	}
	fmt.Println("start Zinx server success")
	go func() {
		var cid uint32 = 0
		for {
			// accept是阻塞的,等待客户端的连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}
			dealConn := NewConnection(conn, cid, self.MsgHandler)
			cid++
			if dealConn != nil {
				//每拿到一个conn后就开启一个新的线程
				go dealConn.Start()
			}

		}

	}()
}

func NewServer(name string) IServer {
	return &Server{
		name:       utils.Config.Name,
		IP:         utils.Config.IP,
		Port:       utils.Config.Port,
		TCPVersion: "tcp4",
		MsgHandler: NewMsgHandler(),
	}
}

func (self *Server) Stop() {
	//todo
}

func (self *Server) Serve() {
	self.Start()
	//todo  做一些启动服务器之外的额外业务
	select {}
}
