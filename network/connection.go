package network

import (
	"time"
	"fmt"
	"github.com/TriangleGo/gorouter/router"
	"github.com/TriangleGo/gorouter/network/protocol"
	"github.com/TriangleGo/gorouter/network/socket"
	"github.com/TriangleGo/gorouter/network/simplebuffer"
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/client"
	"github.com/TriangleGo/gorouter/util"
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
	Running	  bool
	FirstDataChan chan []byte
}

func NewConnection(s *socket.BaseSocket) *Connection {
	return &Connection{Conn: s,
		PacketChan:    make(chan []byte,1024),
		IpcChan:       make(chan protocol.IPCProtocol,512),
		TcpChan:       make(chan protocol.Protocol,1024),
		WsChan:       make(chan protocol.WsProtocol,1024),
		RpcChan:       make(chan protocol.Protocol,512),
		ExitChan:      make(chan string),
		Running:		true,
		FirstDataChan: make(chan []byte, 16)}
}

func (this *Connection) SendTcpChan(p *protocol.Protocol) {
	if this.Running {
		this.TcpChan <- *p
	} else {
		//done
	}
}

func (this *Connection) SendRpcChan(p *protocol.Protocol) {
	if this.Running {
		this.RpcChan <- *p
	} else {
		//done
	}
}

func (this *Connection) SendIpcChan(p *protocol.IPCProtocol) {
	if this.Running {
		this.IpcChan <- *p
	} else {
		//done
	}
}

func (this *Connection) Release() {
	this.Running = false
	close(this.ExitChan)
	close(this.PacketChan)
	close(this.IpcChan)
	close(this.TcpChan)
	close(this.WsChan)
	close(this.RpcChan)
	close(this.FirstDataChan)
}


//will block the accept thread
func (this *Connection) SyncServe() {
	this.serveLoop()
	go this.servePacket()
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
	//prevent crash other goroutines
	defer util.TraceCrashStackAndHandle(func() {
		this.Conn.Close()
	})
	//release the goroutine and connection
	defer func(){
		GetConnectionManager().Release(this)
		logger.Info("Serve Loop Goroutine End !!!\n")
	}()

	var fristPack = true

	for ;this.Running == true; {
		//looping to recv the client
		buf := make([]byte, 4096)
		this.Conn.SetReadDeadline(time.Now().Add( 60 * time.Second))
		n, err := this.Conn.Read(buf)
		if err != nil {
			logger.Info("Client Read Buffer Failed %v %v\r\n", err, n)
			return
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
		//end first pack
		//make the protocal
		this.PacketChan <- buf[0:n]
		//logger.Info("make protocal complete\n")

	} //end for{}
	
	
}


// for handling the packet interrupt
func (this *Connection) servePacket() {
	defer util.TraceCrashStackAndHandle(func() {
		this.Conn.Close()
	})
	const packsize = 20480
	bigBuffer := simplebuffer.NewSimpleBuffer("BigEndian") 
	for {
		select {
		case data, _ := <-this.ExitChan:
			logger.Info("Serve Packet Goroutine End !!! %v \r\n", data)
			return
		case data, _ := <-this.PacketChan:
			// construct the protocal packet
			if len(bigBuffer.Data()) + len(data) > packsize {
				//block this pack
			} else {
				bigBuffer.WriteData(data)
			}
			//construct the protocol and send it to the handler
			
			proto := protocol.NewProtocal()
			ps, err := proto.PraseFromSimpleBuffer(bigBuffer)
			if err != nil {
				logger.Info("BigBuffer remain %v \n",bigBuffer.Data())
				logger.Info("Data Parse failed \n")
				logger.Info("Buffer : %v\n\n\n", bigBuffer.Data())
				continue
			}
			//### parse success ! reset all ### 
			//bigBuffer = simplebuffer.NewSimpleBufferBySize("bigEndian",packsize) // 2 Mb
			for _,v := range ps {
				this.SendTcpChan(v)
			}
		// Packget goroutine no need Ipc
		/*
		case data, ok := <-this.IpcChan:
			logger.Info("ExitHandler %v %v\r\n", data, ok)
			break
		*/
		}
	}
}


// into the handler
func (this *Connection) serveHandle() {
	logger.Info("TCPHandle looping tcp \n")
	defer this.Conn.Close()
	client := client.NewClient(this.Conn)
	
	defer util.TraceCrashStackAndHandle(func() {
		this.Conn.Close()
	})
	
	defer router.GetRouter().GetDisconHandler().Handle(client)

	//serve when connect
	go router.GetRouter().ConnHandler.Handle(client, this.FirstDataChan)

	//loop recv protocol
	for {
		select {
		case  <-this.ExitChan:
			logger.Info("Serve Handle Goroutine Exit !!! \r\n")
			router.GetRouter().GetDisconHandler().Handle(client)
			return
		case data, _ := <-this.TcpChan:
			//logger.Info("TCPHandler %v %v\r\n", data)
			h := router.GetRouter().GetTcpHandler()[data.ModuleId]
			if h != nil {
				c := h.Handle(client,&data)
				if c != nil {
					client = c
				}
			}
		case data, _ := <-this.IpcChan:
			//logger.Info("IPCHandler %v %v\r\n", data.Data)
			h := router.GetRouter().GetIpcHandler()[data.ModuleId]
			if h != nil {
				c := h.Handle(client,data.CommandId,data.Data)
				if c != nil {
					client = c
				}
			}
		}// end select
	} //end for

}



//specify websocket 
func (this *Connection) WSServe() {
	defer util.TraceCrashStackAndHandle(func() {
		this.Conn.Close()
	})
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
	defer util.TraceCrashStackAndHandle(func() {
		this.Conn.Close()
	})
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
				c := h.Handle(client,string(wp.Data))
				if c != nil {
					client = c
				}
			}
		case data, ok := <-this.IpcChan:
			logger.Info("IPCHandler %v %v\r\n", data.Data, ok)
			h := router.GetRouter().GetIpcHandler()[data.ModuleId]
			if h != nil {
				c := h.Handle(client,data.CommandId,data.Data)
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
