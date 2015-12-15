package network

import (
	_"net"
	"gorouter/network/socket"
	"sync"
)

type ConnectionManager struct {
	MapConnections  map[string]*Connection
	Mtx *sync.Mutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		MapConnections:make(map[string]*Connection),
		Mtx : new(sync.Mutex),
	}
}

func (this *ConnectionManager) Produce(s *socket.BaseSocket) *Connection {	
	connection := NewConnection(s)
	key := connection.GetHash()
	//lock
	this.Mtx.Lock()
	this.MapConnections[key] = connection
	this.Mtx.Unlock()
	return connection
}

func (this *ConnectionManager) Release(c *Connection)  {
	this.Mtx.Lock()
	delete(this.MapConnections,c.GetHash())
	this.Mtx.Unlock()
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


