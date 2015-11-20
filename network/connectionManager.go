package network

import (
	_"net"
	"gorouter/network/socket"
)

type ConnectionManager struct {
	
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{}
}

func (this *ConnectionManager) Produce(s *socket.BaseSocket) *Connection {
	connection := NewConnection(s)
	return connection
}

/*
*	static model
*/
var connMgr *ConnectionManager

func GetConnectionManager() *ConnectionManager{
	//fmt.Printf("static Manager call \n")
	if connMgr == nil {
		connMgr = NewConnectionManager() 
	} 
	return connMgr
}


