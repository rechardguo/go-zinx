package znet

import (
	"fmt"
	"zinx/ch10/utils"
	"zinx/ch10/ziface"
)

type MsgHandler struct {
	Apis     map[uint32]ziface.IRouter
	TaskPool []Task
}

type Task struct {
	TaskQueue chan ziface.IRequest
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:     make(map[uint32]ziface.IRouter),
		TaskPool: make([]Task, utils.Config.TaskPoolSize),
	}
}

//添加一个router
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

//将消息发给任务队列来处理
func (self *MsgHandler) SendMsgToTaskQueue(req ziface.IRequest) {
	//从taskpool里找出一个task
	connId := req.GetConnection().GetConnId()
	taskId := int(connId) % len(self.TaskPool)
	//将req放到TaskQueue
	self.TaskPool[taskId].TaskQueue <- req
}

//初始化taskPool
func (self *MsgHandler) InitTaskPool() {
	for i := 0; i < int(utils.Config.TaskPoolSize); i++ {
		self.TaskPool[i] = Task{make(chan ziface.IRequest, utils.Config.TaskQueueMaxLen)}
		go self.StartTask(i, self.TaskPool[i].TaskQueue)
	}
}
func (self *MsgHandler) StartTask(taskId int, task chan ziface.IRequest) {
	for {
		select {
		case req := <-task:
			//如果任务过来了
			fmt.Println("taskId=", taskId, "is process ConnId=", req.GetConnection().GetConnId())
			self.Process(req)
		}
	}
}
