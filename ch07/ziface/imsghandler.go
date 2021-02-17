package ziface

type IMsgHandler interface {
	//添加一个handler
	AddRouter(uint32, IRouter)
	//处理消息
	Process(request IRequest)
}
