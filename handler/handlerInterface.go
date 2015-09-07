package handler

import (
	"gorouter/types"
)

// protocol handler
type Handler interface {
	Init()
	Handle(client *types.Client) *types.Client
}

//some one disconnected
type DisconnectHandler interface {
	Handle(client *types.Client)
}

//some one connected
type ConnectHandler interface {
	Handle(client *types.Client, ch chan []byte) *types.Client
}
