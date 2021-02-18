package znet

import (
	"fmt"
	"sync"
	"zinx/ch09/ziface"
)

type ConnManager struct {
	conns map[uint32]ziface.IConnection
	lock  sync.RWMutex
}

func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		conns: make(map[uint32]ziface.IConnection),
	}
}

//加入Connection
func (self *ConnManager) AddConn(connection ziface.IConnection) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.conns[connection.GetConnId()] = connection
	fmt.Printf("connId=", connection.GetConnId(), " add successful, conn nums=", self.ConnNums())
}

//根据ConnId删除Connection
func (self *ConnManager) RemoveConnByConnId(connId uint32) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.conns[connId] = nil
	fmt.Printf("connId=", connId, " remove successful, conn nums=", self.ConnNums())
}

//删除Connection
func (self *ConnManager) RemoveConn(connection ziface.IConnection) {
	self.lock.Lock()
	defer self.lock.Unlock()
	self.conns[connection.GetConnId()] = nil
	fmt.Printf("connId=", connection.GetConnId(), " remove successful, conn nums=", self.ConnNums())
}

//根据connId找到Connection
func (self *ConnManager) GetConn(connId uint32) ziface.IConnection {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return self.conns[connId]
}

//清除所有的Connection
func (self *ConnManager) CleanAllConn() {
	self.lock.Lock()
	defer self.lock.Unlock()
	for connId := range self.conns {
		self.GetConn(connId).Stop()
		self.RemoveConnByConnId(connId)
	}
}

func (self *ConnManager) ConnNums() int {
	return len(self.conns)
}
