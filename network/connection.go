package network

import (
	"fmt"
	"gorouter/handler"
	"gorouter/network/protocol"
	"gorouter/types"
	"net"
)

// client.go
type Connection struct {
	Conn     net.Conn
	TcpChan  chan protocol.Protocol
	IpcChan  chan protocol.Protocol
	ExitChan chan string
}

func NewConnection(_conn net.Conn) *Connection {
	return &Connection{Conn: _conn,
		IpcChan:  make(chan protocol.Protocol),
		TcpChan:  make(chan protocol.Protocol),
		ExitChan: make(chan string)}
}

func (this *Connection) Serve() {
	go this.serveLoop()
	go this.serveHandle()
}

func (this *Connection) serveLoop() {

	for {
		buf := make([]byte, 4096)
		n, err := this.Conn.Read(buf)
		if err != nil {
			fmt.Printf("Client Read Buffer Failed %v %v\r\n", err, n)
			this.ExitChan <- "TCP_CLOSED"
			break
		}

		proto := protocol.NewProtocal()
		_, err = proto.PraseFromData(buf[0:n], n)
		if err != nil {
			fmt.Printf("Data Parse failed %v\r\n", err)
			continue
		}
		this.TcpChan <- *proto
	}
}

func (this *Connection) serveHandle() {
	fmt.Printf("TCPHandle looping tcp \n")
	
	defer this.Conn.Close()
	
	client := types.NewClient()
	handler.GetRouter().ConnHandler.Handle(client)

	for {
		select {
		case data, ok := <-this.TcpChan:
			fmt.Printf("TCPHandler %v %v\r\n", data, ok)
			h := handler.GetRouter().GetTcpHandler()[data.ModuleId]
			if h != nil {
				c := h.Handle(client)
				if c != nil {
					client = c
				}
			}
			break
		case data, ok := <-this.IpcChan:
			fmt.Printf("IPCHandler %v %v\r\n", string(data.Data[0:10]), ok)
			h := handler.GetRouter().GetTcpHandler()[data.ModuleId]
			if h != nil {
				c := h.Handle(client)
				if c != nil {
					client = c
				}
			}
			break
		case data, ok := <-this.ExitChan:
			fmt.Printf("ExitHandler %v %v\r\n", data, ok)
			handler.GetRouter().GetDisconHandler().Handle(client)
			return
		}
	}

	fmt.Printf("ServeHandle ending \n")
}
