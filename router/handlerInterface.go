package router

import (
	"github.com/TriangleGo/gorouter/client"
	"github.com/TriangleGo/gorouter/network/protocol"
)

// protocol handler
type Handler interface {
	Init()
	Handle(client *client.Client,data *protocol.Protocol) *client.Client
}

// websocket protocol handler
type WSHandler interface {
	Init()
	Handle(client *client.Client,command string,data string) *client.Client
}

// protocol handler
type IpcHandler interface {
	Init()
	Handle(client *client.Client,data interface{}) *client.Client
}

//some one disconnected
type DisconnectHandler interface {
	Handle(client *client.Client)
}

//some one connected
type ConnectHandler interface {
	Handle(client *client.Client, ch chan []byte) *client.Client
}
