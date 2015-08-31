package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

// This example demonstrates a trivial echo server.
func WSServer() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
