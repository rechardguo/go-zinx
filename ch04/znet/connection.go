package znet

import (
	"fmt"
	"net"
	"zinx/ch04/utils"
	"zinx/ch04/ziface"
)

type Connection struct {
	ConnId   uint32
	Conn     *net.TCPConn
	IsClosed bool
	Router   ziface.IRouter
}

func (self *Connection) GetConnId() uint32 {
	return self.ConnId
}

func (self *Connection) GetConn() *net.TCPConn {
	return self.Conn
}

func (self *Connection) RemoteAddr() net.Addr {
	return self.Conn.RemoteAddr()
}

// 连接的读业务方法
func (self *Connection) StartReader() {
	fmt.Printf("Reader Goroutine is running...")
	defer fmt.Println("connId=", self.ConnId, " Reader exist, remote addr is ", self.RemoteAddr().String())
	defer self.Stop()

	for {
		buf := make([]byte, utils.Config.MaxPackageSize)

		_, err := self.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		/*request:=&Request{
			Connection: self,
			Data: buf,
		}
		self.Router.PreHandle(request)
		self.Router.Handle(request)
		self.Router.PostHandle(request)*/
		//为啥要写成下面的样子
		request := Request{
			Connection: self,
			Date:       buf,
		}
		go func(req ziface.IRequest) {
			self.Router.PreHandle(req)
			self.Router.Handle(req)
			self.Router.PostHandle(req)
		}(&request)
	}

}

func (self *Connection) Start() {
	fmt.Println("connection Start().. ConnID=", self.ConnId)
	//启动从当前连接读数据的业务
	go self.StartReader()
	//启动从当前连接写数据的业务

}
func (self *Connection) Stop() {
	fmt.Println("connection Stop().. ConnID=", self.ConnId)
	if self.IsClosed {
		return
	}
	self.IsClosed = true
	//close socket
	self.Conn.Close()
}
func (self *Connection) Handler() {

}

// handleFunc 就是回调函数
func NewConnection(conn *net.TCPConn, cid uint32, router ziface.IRouter) ziface.IConnection {
	return &Connection{
		ConnId:   cid,
		Conn:     conn,
		IsClosed: false,
		Router:   router,
	}
}
