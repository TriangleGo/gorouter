package network

import (
	"time"
	"fmt"
	"gorouter/handler"
	"gorouter/network/protocol"
	"gorouter/types"
	"net"
)

// client.go
type Connection struct {
	Conn          net.Conn
	TcpChan       chan protocol.Protocol
	IpcChan       chan types.IPCSolid
	RpcChan       chan protocol.Protocol
	ExitChan      chan string
	FirstDataChan chan []byte
}

func NewConnection(_conn net.Conn) *Connection {
	return &Connection{Conn: _conn,
		IpcChan:       make(chan types.IPCSolid),
		TcpChan:       make(chan protocol.Protocol),
		RpcChan:       make(chan protocol.Protocol),
		ExitChan:      make(chan string),
		FirstDataChan: make(chan []byte, 1024)}
}

func (this *Connection) Serve() {
	go this.serveLoop()
	go this.serveHandle()
}

func (this *Connection) serveLoop() {
	var fristPack = true
	for {
		//looping to recv the client
		buf := make([]byte, 4096)
		this.Conn.SetReadDeadline(time.Now().Add( 45 * time.Second))
		n, err := this.Conn.Read(buf)
		if err != nil {
			fmt.Printf("Client Read Buffer Failed %v %v\r\n", err, n)
			this.ExitChan <- "TCP_CLOSED"
			break
		}

		//when the user connected,the first data will not parse to protocol
		//it will send to the ConnHandler for modify
		if fristPack {
			maxSize := 1024
			if n > maxSize {
				this.FirstDataChan <- buf[0:maxSize]
			} else {
				this.FirstDataChan <- buf[0:n]
			}
		}

		//construct the protocol and send it to the handler
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

	client := types.NewClient(this.Conn)

	//serve when connect
	go handler.GetRouter().ConnHandler.Handle(client, this.FirstDataChan)

	//loop recv protocol
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
			fmt.Printf("IPCHandler %v %v\r\n", data.Data, ok)
			h := handler.GetRouter().GetTcpHandler()[uint8(data.ModuleId)]
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
