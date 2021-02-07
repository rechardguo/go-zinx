package znet

import (
	"fmt"
	"net"
	"zinx/ch01/ziface"
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
		fmt.Println("listenr error", err)
	}
	fmt.Println("start Zinx server success")
	go func() {
		for {
			// accept是阻塞的,等待客户端的连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}
			go func() {
				for {
					//这样写时干什么用的？ 开启一个异步的立即执行的方法
					//相当于java的thread
					buf := make([]byte, 512)
					//从Connection里将数据读入到buf里
					//count为读取的字节数
					count, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recieve buf error", err)
						continue
					}
					fmt.Printf("receive message from client:%s \n", buf[:count])
					//回显功能
					if _, err := conn.Write(buf[:count]); err != nil {
						fmt.Println("write buf error", err)
						continue
					}
				}
			}()

		}

	}()
}

func (self *Server) Stop() {
	//todo
}

func (self *Server) Serve() {
	self.Start()

	//todo  做一些启动服务器之外的额外业务
	select {}
}

func NewServer(name string) ziface.IServer {
	return &Server{
		name:       name,
		IP:         "127.0.0.1",
		Port:       9000,
		TCPVersion: "tcp4",
	}
}
