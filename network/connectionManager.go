package network

import (
	"time"
	_"net"
	"github.com/TriangleGo/gorouter/network/socket"
	"sync"
	"runtime"
	"runtime/debug"
)

type ConnectionManager struct {
	MapConnections  map[string]*Connection
	Mtx *sync.Mutex
	LastGCTime int64
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		MapConnections:make(map[string]*Connection),
		Mtx : new(sync.Mutex),
		LastGCTime : time.Now().Unix(),
	}
}

func (this *ConnectionManager) Produce(s *socket.BaseSocket) *Connection {	
	connection := NewConnection(s)
	key := connection.Conn.MakeAddrToHash()
	//lock
	this.Mtx.Lock()
	this.MapConnections[key] = connection
	this.Mtx.Unlock()
	return connection
}

func (this *ConnectionManager) GetConnection(s *socket.BaseSocket) *Connection {	
	hashKey := s.MakeAddrToHash()
	return this.MapConnections[hashKey]
}

//delete the connection in the Map 
func (this *ConnectionManager) Release(c *Connection)  {
	this.Mtx.Lock()
	//delete the connection in the Map by conn's Hash value
	delete(this.MapConnections,c.Conn.MakeAddrToHash())
	this.Mtx.Unlock()
	
	Now := time.Now().Unix()
	if Now > this.LastGCTime + 60 * 10 { // 10 min go gc
		runtime.Gosched()
		runtime.GC()
		debug.FreeOSMemory()
		this.LastGCTime = Now
	}
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


