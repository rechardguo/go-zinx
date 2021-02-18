package ziface

/**
connection的管理器
*/
type IConnManager interface {
	//加入Connection
	AddConn(connection IConnection)
	//删除Connection
	RemoveConn(connection IConnection)
	//根据ConnId删除Connection
	RemoveConnByConnId(connId uint32)
	//根据connId找到Connection
	GetConn(connId uint32) IConnection
	//清除所有的Connection
	CleanAllConn()
	//得到Connection的连接数
	ConnNums() int
}
