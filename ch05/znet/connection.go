package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/ch05/ziface"
)

type Connection struct {
	ConnId   uint32
	TCPConn  *net.TCPConn
	IsClosed bool
	Router   ziface.IRouter
}

func (self *Connection) GetConnId() uint32 {
	return self.ConnId
}

func (self *Connection) GetConn() *net.TCPConn {
	return self.TCPConn
}

func (self *Connection) RemoteAddr() net.Addr {
	return self.TCPConn.RemoteAddr()
}

func (self *Connection) Send(msgId uint32, data []byte) error {
	dp := NewDataPack()
	data, err := dp.Pack(NewPackage(msgId, data))
	if err != nil {
		return err
	}
	if _, err := self.TCPConn.Write(data); err != nil {
		return err
	}
	return nil
}

//连接的读业务方法
func (self *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connId=", self.ConnId, " Reader exist, remote addr is ", self.RemoteAddr().String())
	defer self.Stop()

	for {
		//buf := make([]byte, utils.Config.MaxPackageSize)
		//_, err := self.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}
		dp := NewDataPack()
		//先读取出消息头字节数组
		headBytes := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(self.GetConn(), headBytes); err != nil {
			fmt.Println("error in read head", err)
			continue
		}

		msg, err := dp.UnPack(headBytes)
		if err != nil {
			fmt.Println("error in unpack ", err)
			continue
		}
		//构造出消息体字节数组
		bodyBytes := make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(self.GetConn(), bodyBytes); err != nil {
			fmt.Println("error in read body", err)
			continue
		}
		msg.SetData(bodyBytes)
		request := Request{
			Connection: self,
			Message:    msg,
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
	self.TCPConn.Close()
}
func (self *Connection) Handler() {

}

//handleFunc 就是回调函数
func NewConnection(conn *net.TCPConn, cid uint32, router ziface.IRouter) ziface.IConnection {
	return &Connection{
		ConnId:   cid,
		TCPConn:  conn,
		IsClosed: false,
		Router:   router,
	}
}
