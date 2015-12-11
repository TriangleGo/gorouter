package router

import (
	"gorouter/handler/client"
	"gorouter/network/protocol"
)

// protocol handler
type Handler interface {
	Init()
	Handle(client *client.Client,data *protocol.Protocol) *client.Client
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