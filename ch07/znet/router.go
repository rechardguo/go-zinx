package znet

import "zinx/ch07/ziface"

type BaseRouter struct {
}

func (self *BaseRouter) PreHandle(request ziface.IRequest) {
}
func (self *BaseRouter) Handle(request ziface.IRequest) {
}
func (self *BaseRouter) PostHandle(request ziface.IRequest) {
}
