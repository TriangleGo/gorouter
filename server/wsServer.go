package server

import (
	"fmt"
    "io"
    "net/http"

    "golang.org/x/net/websocket"
)

type WebSocketServer struct {
	host string
}

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	for {
		buf := make([]byte,4096)
		n,err := ws.Read(buf)
		if err != nil {
			fmt.Printf("err %v \n",err)
			break
		}
		fmt.Printf("data %v \n",string(buf[0:n]))
	}
    	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func NewWSServer(hostname string) *WebSocketServer{
	return &WebSocketServer{host : hostname}
	
}

func (this *WebSocketServer) Run() {
	http.Handle("/ws", websocket.Handler(EchoServer))
    	go func() {
		err := http.ListenAndServe(this.host, nil)
	    	if err != nil {
			panic("ListenAndServe: " + err.Error())
	    	}
	}()
}
