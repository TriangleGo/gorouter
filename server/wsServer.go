package server

import (
	_"io"
	"net/http"
	"github.com/TriangleGo/gorouter/network/socket"	    	
	"github.com/TriangleGo/gorouter/network"	   
	"github.com/TriangleGo/gorouter/logger"
	//copy file golang.org/x/net/websocket
	"github.com/TriangleGo/gorouter/lib/websocket"
	"github.com/TriangleGo/gorouter/router"
)

type WebSocketServer struct {
	host string
}

// Echo the data received on the WebSocket.
func WsServerProc(ws *websocket.Conn) {
	logger.Info("wsserver connection \n")
	// conn starting
	network.GetConnectionManager().
		Produce(socket.NewBaseSocket(ws)).WSServe()
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

func (this *WebSocketServer) GetRouter() *router.Router{
	return router.GetRouter()
}
