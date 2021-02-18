package ziface

type IMsgHandler interface {
	//添加一个handler
	AddRouter(uint32, IRouter)
	//处理消息
	Process(request IRequest)
	//将消息交给任务队列来处理
	SendMsgToTaskQueue(request IRequest)
	//初始化taskPool
	InitTaskPool()
}
