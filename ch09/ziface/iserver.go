package ziface

type IServer interface {
	Start()
	Serve()
	Stop()
	AddRouter(msgId uint32, router IRouter)
	//connection 连上的hook函数
	OnConnConnected(func(IConnection))
	//connection 断开的hook函数
	OnConnClosed(func(IConnection))
	//调用connection 连上的hook函数
	CallConnConnected(IConnection)
	//调用connection 断开的hook函数
	CallConnClosed(IConnection)
	GetConnManager() IConnManager
}
