package znet

import (
	"fmt"
	"net"
	"zinx/ch02/ziface"
)

type Connection struct {
	ConnId    uint32
	Conn      *net.TCPConn
	IsClosed  bool
	handleApi ziface.HandleFunc
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

//连接的读业务方法
func (self *Connection) StartReader() {
	fmt.Printf("Reader Goroutine is running...")
	defer fmt.Println("connId=", self.ConnId, " Reader exist, remote addr is ", self.RemoteAddr().String())
	defer self.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := self.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}

		if err := self.handleApi(self.Conn, buf, cnt); err != nil {
			fmt.Println("connId=", self.ConnId, "handler is error", err)
			continue
		}
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

//handleFunc 就是回调函数
func NewConnection(conn *net.TCPConn, cid uint32, handleFunc ziface.HandleFunc) ziface.IConnection {
	return &Connection{
		ConnId:    cid,
		Conn:      conn,
		IsClosed:  false,
		handleApi: handleFunc,
	}
}
