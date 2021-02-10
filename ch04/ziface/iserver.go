package ziface

type IServer interface {
	Start()
	Serve()
	Stop()
	AddRouter(router IRouter)
}
