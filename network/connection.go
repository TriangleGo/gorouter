package network

import (
	"time"
	"fmt"
	"gorouter/router"
	"gorouter/network/protocol"
	"gorouter/network/socket"
	"gorouter/network/simplebuffer"
	"gorouter/logger"
	"gorouter/client"
	"gorouter/util"
	_"net"
)

// client.go
type Connection struct {
	Conn          *socket.BaseSocket
	PacketChan	  chan []byte
	TcpChan       chan protocol.Protocol
	WsChan       chan protocol.WsProtocol
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
		WsChan:       make(chan protocol.WsProtocol),
		RpcChan:       make(chan protocol.Protocol),
		ExitChan:      make(chan string),
		FirstDataChan: make(chan []byte, 1024)}
}


//will block the accept thread
func (this *Connection) SyncServe() {
	this.serveLoop()
	go this.serveHandle()
}

//will not block the accept thread
func (this *Connection) AsyncServe() {
	go this.serveLoop()
	go this.servePacket()
	go this.serveHandle()
}


// recving the data from socket
func (this *Connection) serveLoop() {
	defer util.TraceCrashStack()
	var fristPack = true
	for {
		//looping to recv the client
		buf := make([]byte, 8192)
		this.Conn.SetReadDeadline(time.Now().Add( 60 * time.Second))
		n, err := this.Conn.Read(buf)
		if err != nil {
			logger.Info("Client Read Buffer Failed %v %v\r\n", err, n)
			this.ExitChan <- err.Error()
			this.ExitChan <- err.Error()
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
	const packsize = 20480
	bigBuffer := simplebuffer.NewSimpleBufferBySize("bigEndian",packsize) // 2 Mb
	for {
		select {
		case data, ok := <-this.PacketChan:
			logger.Info("ExitHandler %v %v\r\n", data, ok)
			// construct the protocal packet
			if len(bigBuffer.Data()) + len(data) > packsize {
				//block this pack
			} else {
				bigBuffer.WriteData(data)
			}
			//construct the protocol and send it to the handler
			
			proto := protocol.NewProtocal()
			_, err := proto.PraseFromData(bigBuffer.Data(), bigBuffer.Size())
			if err != nil {
				logger.Info("Data Parse failed \n")
				logger.Info("Buffer : %v\n\n\n", bigBuffer.Data())
				continue
			}
			//### parse success ! reset all ### 
			bigBuffer = simplebuffer.NewSimpleBufferBySize("bigEndian",packsize) // 2 Mb
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
			h := router.GetRouter().GetIpcHandler()[data.ModuleId]
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



//specify websocket 
func (this *Connection) WSServe() {
	// init handle
	go this.serveWsHandle()
	var firstPack = true
	// parse data
	const packsize = 8192
	for {
		var padding []byte
		buf := make([]byte,packsize)
		//no heartbeat after 60s will disconnect
		this.Conn.SetReadDeadline(time.Now().Add( 60 * time.Second))
		n,err := this.Conn.Read(buf)
		// error handle 
		if err != nil {
			this.ExitChan <- err.Error()
			return 
		}
		// pack handle
		if n >= packsize {
			tmpBuff := make([]byte,20480)
			tn,_ := this.Conn.Read(tmpBuff)
			if tn > 20480 {
				continue
			}
			padding = make([]byte,tn+n)
			copy(padding[0:],buf)
			copy(padding[n:],tmpBuff)
			n = n+tn
		} else {
			padding = buf
		}
		logger.Info("websocket read size %v data %v\n",n,padding[0:n])
		if err != nil {
			return
		}
		//parse data in jso		
		wp,err := protocol.NewWsProtocolFromData(padding[0:n])
		if err != nil {
			continue
		}
		//first pack channel
		if firstPack {
			this.FirstDataChan <- wp.ToBytes()
			firstPack = false
		}
		//pass value to handle goroutine
		this.WsChan <- *wp
	}
}


func (this *Connection) serveWsHandle() {
	defer util.TraceCrashStack()
	logger.Info("Websocket handle looping tcp \n")

	defer this.Conn.Close()

	client := client.NewClient(this.Conn)

	//serve when connect
	go router.GetRouter().ConnHandler.Handle(client, this.FirstDataChan)

	//loop recv protocol
	for {
		select {
		case wp, _ := <-this.WsChan:
			//get handler 
			logger.Info("handle module:%v command:%v data:%v \n",wp.Module,wp.Command,string(wp.Data))
			h := router.GetRouter().GetWsHandler()[wp.Module]
			if h != nil {
				c := h.Handle(client,wp.Command,string(wp.Data))
				if c != nil {
					client = c
				}
			}
		case data, ok := <-this.IpcChan:
			logger.Info("IPCHandler %v %v\r\n", data.Data, ok)
			h := router.GetRouter().GetIpcHandler()[data.ModuleId]
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
