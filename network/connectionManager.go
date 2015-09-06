package network

import (
	"fmt"
	"net"
)

type ConnectionManager struct {
	
}

func Push() {
	fmt.Printf("test")
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{}
}

func (this *ConnectionManager) Produce(c *net.Conn) *Connection {
	connection := NewConnection(*c)
	return connection
}
