package main

import (
	_"time"
	_"net"
	"fmt"
	"gorouter/handler"
	"gorouter/router"
	"os"
	"runtime"
	"runtime/pprof"
	"gorouter/server"
	_"gorouter/network"
	"gorouter/logger"
)

//global
var (
	exitChan chan bool
)

func main() {
	logger.Info("Go router running \r\n")
	//init logger
	logger.GetLogger().SetLogLevel(5,5)
	logger.GetLogger().SetOutputDir(`c:\`)
	logger.GetLogger().Init("testlog")
	
	//init tcp server
	logger.Info("Go run TCPServer \r\n")
	tcpServer := server.NewTCPServer("tcp", ":9093")

	//setup handler
	tcpServer.GetRouter().SetConnHandler(&handler.ConnHandlerImpl{})
	tcpServer.GetRouter().SetDisconHandler(&handler.DisconHandlerImpl{})
	//register handler
	tcpServer.GetRouter().SetTcpHandler(map[uint8]router.Handler{
		0x0: &handler.LoginHandlerImpl{},
		0x1: &handler.LoginHandlerImpl{},
		0x2: &handler.LoginHandlerImpl{},
	})
	//register Ipc handler
	tcpServer.GetRouter().SetIpcHandler(map[string]router.IpcHandler{
		
	})
	
	//startup
	tcpServer.Run()
	
	
	logger.Info("Go run HTTPServer \r\n")
	server.NewHTTPServer(":9091").Run()
	logger.Info("Go run WSServer \r\n")
	
	wsServer := server.NewWSServer(":9092")
	//register handler
	wsServer.GetRouter().SetWsHandler(map[string]router.WSHandler{
		"system": &handler.WsUserHandlerImpl{},
		"user": &handler.WsUserHandlerImpl{},
		"chat": &handler.WsUserHandlerImpl{},
	})
	
	//startup
	wsServer.Run()
	
	/*
	a ,_:= net.Dial("tcp","127.0.0.1:9093")
	logger.Info(" remote %v\n",a.RemoteAddr())
	socket := network.NewBaseSocket(a)
	socket.LocalAddr()
	socket.RemoteAddr()
	socket.SetDeadline(time.Now().Add( time.Second * 10))
	socket.Write([]byte{0,0,0,4,0,0,0,1})
	socket.Close()
	logger.Info(" remote %v\n",a.RemoteAddr())
*/
	logger.Info("CpuProfile %v \n", string(runtime.CPUProfile()))
	for {
		var cmd string
		fmt.Scanf("%s", &cmd)
		if cmd == "monitor" {
			pprof.StartCPUProfile(os.Stdout)
		}
	}
	<-exitChan
}


