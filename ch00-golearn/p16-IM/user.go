package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)

type User struct {
	Addr   string
	Name   string
	Conn   net.Conn
	C      chan string
	Server *Server
}

func (u *User) ListenMsg() {

	for {
		msg := <-u.C
		u.Sendmsg(msg)
	}
}

func (u *User) Online() {
	//往onlineMap里加入用户
	u.Server.Lock.Lock()
	u.Server.OnlineUser[u.Name] = u
	u.Server.Lock.Unlock()
	//广播消息
	u.Server.Broadcast(u, "online")
}

func (u *User) Offline() {
	//往onlineMap里加入用户
	u.Server.Lock.Lock()
	delete(u.Server.OnlineUser, u.Name)
	u.Server.Lock.Unlock()
	//广播消息
	u.Server.Broadcast(u, "offline")
}

func (u *User) Sendmsg(msg string) {
	u.Conn.Write([]byte(msg + "\n"))
}
func (u *User) DoMessage(msg string) {

	if msg == "who" {
		var response bytes.Buffer
		u.Server.Lock.Lock()
		for _, u := range u.Server.OnlineUser {
			response.WriteString(fmt.Sprintf("[%s] %s online \n", u.Addr, u.Name))
		}
		u.Server.Lock.Unlock()
		// 不能用println，这个表示是server本身打印
		// println()
		// u.Conn.Write([]byte(response))
		u.Sendmsg(response.String())
	} else if len(msg) > 5 && msg[:5] == "user|" {
		reqArr := strings.Split(msg, "|")

		reqNum := reqArr[1]
		userName := reqArr[2]

		//查询用户是否存在
		u.Server.Lock.Lock()

		if user := u.Server.OnlineUser[userName]; user != nil {
			u.Conn.Write([]byte(reqNum + "|1"))
		} else {
			u.Conn.Write([]byte(reqNum + "|0"))
		}
		u.Server.Lock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//改名功能
		strs := strings.Split(msg, "|")
		userName := strs[1]
		//名字不存在则修改
		if u.Server.OnlineUser[userName] != nil {
			u.Sendmsg("name " + userName + " exist, please change one")
		} else {
			u.Server.Lock.Lock()
			delete(u.Server.OnlineUser, u.Name)
			u.Name = userName
			u.Server.OnlineUser[userName] = u
			u.Server.Lock.Unlock()
			u.Sendmsg("you change name to :" + userName)
		}
	} else if len(msg) > 3 && msg[:3] == "to|" {
		//私聊功能
		remoteName := strings.Split(msg, "|")[1]

		if remoteName == "" {
			u.Sendmsg("send message incorrect, usage: to|<username>|<content>")
			return
		}
		remoteUser := u.Server.OnlineUser[remoteName]
		if remoteUser == nil {
			u.Sendmsg("username=" + remoteName + " not exist")
			return
		}
		content := strings.Split(msg, "|")[2]

		if content == "" {
			u.Sendmsg("you can not send empty message")
			return
		}

		remoteUser.Sendmsg(u.Name + " say: " + content)

	} else {
		u.Server.Broadcast(u, msg)
	}
}

func NewUser(c net.Conn, server *Server) *User {
	user := &User{
		c.RemoteAddr().String(),
		c.RemoteAddr().String(),
		c,
		make(chan string),
		server,
	}

	go user.ListenMsg()

	return user
}
