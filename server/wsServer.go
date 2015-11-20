package server

import (
	_"io"
	"net/http"
	"gorouter/network/socket"	    	
	"gorouter/network"	   
	"gorouter/logger"
	"golang.org/x/net/websocket"
)

type WebSocketServer struct {
	host string
}

// Echo the data received on the WebSocket.
func WsServerProc(ws *websocket.Conn) {
	logger.Info("wsserver connection \n")
	// conn starting
	network.GetConnectionManager().
		Produce(socket.NewBaseSocket(ws)).SyncServe()
}

// This example demonstrates a trivial echo server.
func NewWSServer(hostname string) *WebSocketServer{
	return &WebSocketServer{host : hostname}
	
}

func (this *WebSocketServer) Run() {
	http.Handle("/ws", websocket.Handler(WsServerProc))
    	go func() {
		err := http.ListenAndServe(this.host, nil)
	    	if err != nil {
			panic("ListenAndServe: " + err.Error())
	    	}
	}()
}
