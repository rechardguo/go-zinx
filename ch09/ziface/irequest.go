package ziface

type IRequest interface {
	GetConnection() IConnection
	GetMessage() IMessage
}
