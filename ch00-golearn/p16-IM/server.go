package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Server struct {
	IP         string
	Port       int
	Message    chan string
	OnlineUser map[string]*User
	Lock       sync.RWMutex
}

func NewServer(ip string, port int) *Server {
	return &Server{
		IP:         ip,
		Port:       port,
		Message:    make(chan string),
		OnlineUser: make(map[string]*User, 5),
		Lock:       sync.RWMutex{}, //接口类型，只能&
	}
}

func (s *Server) ListenMessage() {

	for {
		msg := <-s.Message
		s.Lock.Lock()
		for _, user := range s.OnlineUser {
			//往 user.C 写入消息
			user.C <- msg
		}
		s.Lock.Unlock()
	}
}

func (s *Server) Broadcast(user *User, msg string) {
	s.Message <- "[" + user.Name + "]" + msg
}

func (s *Server) Start() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("error in listen", err)
		return
	}

	defer listen.Close()

	//启动go程监听消息
	go s.ListenMessage()

	for {
		conn, err := listen.Accept()

		if err != nil {
			fmt.Println("error in accept", err)
			continue
		}

		go s.Handler(conn)
	}
}

func (s *Server) Handler(conn net.Conn) {
	user := NewUser(conn, s)

	user.Online()

	isaLive := make(chan bool)
	//这里会阻塞住，
	//接收到用户消息并发给所有的
	go func() {
		for {
			buf := make([]byte, 4096)

			c, err := conn.Read(buf)
			if c == 0 {
				//客户下线
				user.Offline()
				return
			}
			if err != nil {
				fmt.Println("read client message error", err)
				return
			}
			// server进行的消息广播
			//s.Broadcast(user, string(buf[:c]))

			user.DoMessage(string(buf[:c-1]))
			isaLive <- true
		}
	}()

	/**
	for是一直循环，case里会阻塞
	time.After 会重建一个chan, 并且10秒后才会往你放东西
	*/
	for {
		select {
		case <-isaLive:
			//如果isaLive一直不活跃，则下面的case在10s后就会触发
			//如果isaLive活跃，则下面的case会重新重建
		case <-time.After(120 * time.Second):
			// 下线用户
			s.Lock.Lock()
			delete(s.OnlineUser, user.Name)
			user.Sendmsg("you are kick out!")
			user.Conn.Close()

			s.Lock.Unlock()
			//return 会退出外面的for
			return
		}
	}

}

func main() {
	server := NewServer("127.0.0.1", 9999)
	server.Start()
}
