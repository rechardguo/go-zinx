package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/ch09/utils"
	"zinx/ch09/ziface"
)

type Connection struct {
	//connection对应的server
	TCPServer  ziface.IServer
	ConnId     uint32
	TCPConn    *net.TCPConn
	IsClosed   bool
	MsgHandler ziface.IMsgHandler
	ExitChan   chan bool
	MsgChan    chan []byte
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
	/*if _, err := self.TCPConn.Write(data); err != nil {
		return err
	}*/
	//将数据发到管道里
	self.MsgChan <- data
	return nil
}

//连接的读业务方法
func (self *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println("connId=", self.ConnId, "[conn Reader exit!], remote addr is ", self.RemoteAddr().String())
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
			return
		}

		msg, err := dp.UnPack(headBytes)
		if err != nil {
			fmt.Println("error in unpack ", err)
			return
		}
		//构造出消息体字节数组
		bodyBytes := make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(self.GetConn(), bodyBytes); err != nil {
			fmt.Println("error in read body", err)
			return
		}
		msg.SetData(bodyBytes)
		request := &Request{
			Connection: self,
			Message:    msg,
		}

		if utils.Config.TaskPoolSize > 0 {
			//将消息发送到一个任务队列里，由go channel来处理
			self.MsgHandler.SendMsgToTaskQueue(request)
		} else {
			go self.MsgHandler.Process(request)
		}

		/*go func(req ziface.IRequest) {
			self.Router.PreHandle(req)
			self.Router.Handle(req)
			self.Router.PostHandle(req)
		}(&request)*/
	}

}

func (self *Connection) Start() {
	fmt.Println("connection Start().. ConnId=", self.ConnId)
	//启动从当前连接读数据的业务
	go self.StartReader()
	//启动从当前连接写数据的业务
	go self.StartWriter()

}

//x, y := <-c, <-c这句会一直等待计算结果发送到channel中
func (self *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(self.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
		//MsgChan会阻塞等消息到来
		case data := <-self.MsgChan:
			if _, err := self.GetConn().Write(data); err != nil {
				fmt.Println("send data err", err)
				return
			}
		case <-self.ExitChan:
			//Stop()时候回发出
			//reader退出，writer也会退出
			return
		}
	}
}
func (self *Connection) Stop() {
	fmt.Println("connection Stop()...ConnID=", self.ConnId)
	if self.IsClosed {
		return
	}
	self.IsClosed = true
	//close socket
	self.TCPConn.Close()
	//连接关闭的hook函数
	self.TCPServer.CallConnClosed(self)
	//关闭管道
	close(self.ExitChan)
	close(self.MsgChan)
}

//handleFunc 就是回调函数
func NewConnection(server ziface.IServer, conn *net.TCPConn, cid uint32, handler ziface.IMsgHandler) ziface.IConnection {
	c := &Connection{
		TCPServer:  server,
		ConnId:     cid,
		TCPConn:    conn,
		IsClosed:   false,
		MsgHandler: handler,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
	}
	//连接连上的hook函数
	c.TCPServer.CallConnConnected(c)
	return c
}
