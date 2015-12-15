package network

import (
	"time"
	"fmt"
	_"gorouter/handler"
	"gorouter/router"
	"gorouter/network/protocol"
	"gorouter/network/socket"
	"gorouter/network/simplebuffer"
	"gorouter/logger"
	"gorouter/client"
	"gorouter/util"
	"gorouter/util/hash"
	_"net"
)

// client.go
type Connection struct {
	Conn          *socket.BaseSocket
	PacketChan	  chan []byte
	TcpChan       chan protocol.Protocol
	IpcChan       chan protocol.IPCProtocol
	RpcChan       chan protocol.Protocol
	ExitChan      chan string
	FirstDataChan chan []byte
}

func NewConnection(s *socket.BaseSocket) *Connection {
	return &Connection{Conn: s,
		PacketChan:    make(chan []byte),
		IpcChan:       make(chan protocol.IPCProtocol),
		TcpChan:       make(chan protocol.Protocol),
		RpcChan:       make(chan protocol.Protocol),
		ExitChan:      make(chan string),
		FirstDataChan: make(chan []byte, 1024)}
}

func (this *Connection) GetHash() string {
	sId := fmt.Sprintf("%x%x%x%x",&this.Conn,&this.IpcChan,&this.TcpChan,&this.RpcChan)
	C32 := hash.HashC32([]byte(sId))
	return fmt.Sprintf("%x",C32)
}


func (this *Connection) SyncServe() {
	this.serveLoop()
	go this.serveHandle()
}

func (this *Connection) AsyncServe() {
	go this.serveLoop()
	go this.servePacket()
	go this.serveHandle()
}

func (this *Connection) serveLoop() {
	defer util.TraceCrashStack()
	var fristPack = true
	for {
		//looping to recv the client
		buf := make([]byte, 8096)
		this.Conn.SetReadDeadline(time.Now().Add( 60 * time.Second))
		n, err := this.Conn.Read(buf)
		if err != nil {
			logger.Info("Client Read Buffer Failed %v %v\r\n", err, n)
			this.ExitChan <- "TCP_CLOSED"
			this.ExitChan <- "TCP_CLOSED"
			break
		}
		
		logger.Info("recv data %v \n",buf[0:n])

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
		//end first pack
		
		//make the protocal
		this.PacketChan <- buf[0:n]

	} //end for{}
}


// for handling the packet interrupt
func (this *Connection) servePacket() {
	defer util.TraceCrashStack()
	bigBuffer := simplebuffer.NewSimpleBufferBySize("bigEndian",20480) // 2 Mb
	for {
		select {
		case data, ok := <-this.PacketChan:
			logger.Info("ExitHandler %v %v\r\n", data, ok)
			// construct the protocal packet
			bigBuffer.WriteData(data)
			//construct the protocol and send it to the handler
			
			proto := protocol.NewProtocal()
			_, err := proto.PraseFromData(bigBuffer.Data(), bigBuffer.Size())
			if err != nil {
				logger.Info("Data Parse failed \n")
				logger.Info("Buffer : %v\n\n\n", bigBuffer.Data())
				continue
			}
			//### parse success ! reset all ### 
			bigBuffer = simplebuffer.NewSimpleBufferBySize("bigEndian",20480) // 2 Mb
			this.TcpChan <- *proto
			break
		// Packget goroutine no need Ipc
		/*
		case data, ok := <-this.IpcChan:
			logger.Info("ExitHandler %v %v\r\n", data, ok)
			break
		*/
		case data, ok := <-this.ExitChan:
			logger.Info("ExitHandler %v %v\r\n", data, ok)
			return
		}
	}
}


// into the handler
func (this *Connection) serveHandle() {
	defer util.TraceCrashStack()
	logger.Info("TCPHandle looping tcp \n")

	defer this.Conn.Close()

	client := client.NewClient(this.Conn)

	//serve when connect
	go router.GetRouter().ConnHandler.Handle(client, this.FirstDataChan)

	//loop recv protocol
	for {
		select {
		case data, ok := <-this.TcpChan:
			logger.Info("TCPHandler %v %v\r\n", data, ok)
			h := router.GetRouter().GetTcpHandler()[data.ModuleId]
			if h != nil {
				c := h.Handle(client,&data)
				if c != nil {
					client = c
				}
			}
			break
		case data, ok := <-this.IpcChan:
			logger.Info("IPCHandler %v %v\r\n", data.Data, ok)
			h := router.GetRouter().GetIpcHandler()[uint8(data.ModuleId)]
			if h != nil {
				c := h.Handle(client,data.Data)
				if c != nil {
					client = c
				}
			}
			break
		case data, ok := <-this.ExitChan:
			logger.Info("ExitHandler %v %v\r\n", data, ok)
			router.GetRouter().GetDisconHandler().Handle(client)
			return
		}
	}

	fmt.Printf("ServeHandle ending \n")
}
