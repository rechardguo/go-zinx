package znet

import (
	"errors"
	"fmt"
	"net"
	. "zinx/ch02/ziface"
)

type Server struct {
	name       string
	IP         string
	Port       int
	TCPVersion string
}

func (self *Server) Start() {
	fmt.Printf("Server %s at Ip :%s, Port:%d  is starting \n", self.name, self.IP, self.Port)

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
			dealConn := NewConnection(conn, cid, CallBackToClient)
			cid++
			if dealConn != nil {
				//每拿到一个conn后就开启一个新的线程
				go dealConn.Start()
			}

		}

	}()
}

//回调函数
//目前写死后期会由客户端来自定义
func CallBackToClient(conn *net.TCPConn, buf []byte, cnt int) error {
	//回显功能
	if _, err := conn.Write(buf[:cnt]); err != nil {
		fmt.Println("write buf error", err)
		return errors.New(fmt.Sprintf("write buf error:%s", err))
	}
	return nil
}

func NewServer(name string) IServer {
	return &Server{
		name:       name,
		IP:         "127.0.0.1",
		Port:       9000,
		TCPVersion: "tcp4",
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
