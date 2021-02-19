package znet

import (
	"fmt"
	"net"
	"zinx/ch09/utils"
	. "zinx/ch09/ziface"
)

type Server struct {
	name        string
	IP          string
	Port        uint32
	TCPVersion  string
	MsgHandler  IMsgHandler
	ConnManager IConnManager
	//连接连上的hook函数
	ConnConnectedHook func(IConnection)
	//连接断开的hook函数
	ConnClosedHook func(IConnection)
}

func (self *Server) GetConnManager() IConnManager {
	return self.ConnManager
}

func (self *Server) OnConnConnected(hookfn func(IConnection)) {
	self.ConnConnectedHook = hookfn
}

func (self *Server) OnConnClosed(hookfn func(IConnection)) {
	self.ConnClosedHook = hookfn
}

func (self *Server) CallConnConnected(connection IConnection) {
	if self.ConnConnectedHook != nil {
		self.ConnConnectedHook(connection)
	}
}

func (self *Server) CallConnClosed(connection IConnection) {
	if self.ConnClosedHook != nil {
		self.ConnClosedHook(connection)
	}
}

func (self *Server) AddRouter(msgId uint32, router IRouter) {
	self.MsgHandler.AddRouter(msgId, router)
}

func (self *Server) Start() {
	fmt.Printf("Server %s  version : %s at Ip :%s, Port: %d  is starting \n", self.name, utils.Config.Version, self.IP, self.Port)
	fmt.Printf("maxConne : %d, maxPackageSize : %d \n", utils.Config.MaxConn, utils.Config.MaxPackageSize)

	//启动任务处理池
	self.MsgHandler.InitTaskPool()

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
			//判断是否超过最大连接数
			if self.ConnManager.ConnNums() >= int(utils.Config.MaxConn) {
				fmt.Println("===========>too many connection, max conn=", utils.Config.MaxConn)
				conn.Close()
				continue
			}

			dealConn := NewConnection(self, conn, cid, self.MsgHandler)
			self.ConnManager.AddConn(dealConn)
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
		name:        utils.Config.Name,
		IP:          utils.Config.IP,
		Port:        utils.Config.Port,
		TCPVersion:  "tcp4",
		MsgHandler:  NewMsgHandler(),
		ConnManager: NewConnManager(),
	}

}

func (self *Server) Stop() {
	//做一些资源的释放
	self.ConnManager.CleanAllConn()
}

func (self *Server) Serve() {
	self.Start()
	//todo  做一些启动服务器之外的额外业务
	select {}
}
