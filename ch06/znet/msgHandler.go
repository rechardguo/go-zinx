package znet

import (
	"fmt"
	"zinx/ch06/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

//添加一个handler
func (self *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	//handler已经存在
	if _, ok := self.Apis[msgId]; ok {
		panic(fmt.Sprintf("message id=%d router exist!", msgId))
		return
	}
	self.Apis[msgId] = router
}

//处理消息
func (self *MsgHandler) Process(req ziface.IRequest) {
	if router, ok := self.Apis[req.GetMessage().GetMsgId()]; ok {
		router.PreHandle(req)
		router.Handle(req)
		router.PostHandle(req)
	} else {
		panic(fmt.Sprintf("message id=%d router is NOT FOUND! Please ADD", req.GetMessage().GetMsgId()))
	}
}
