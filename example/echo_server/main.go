package main

import (
	"gorouter/router"
	"gorouter/server"
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
	tcpServer.GetRouter().SetConnHandler(&ConnHandlerImpl{})
	tcpServer.GetRouter().SetDisconHandler(&DisconHandlerImpl{})
	//register handler
	tcpServer.GetRouter().SetTcpHandler(map[uint8]router.Handler{
		0x0: &MyHandlerImpl{},
	})
	//register Ipc handler
	tcpServer.GetRouter().SetIpcHandler(map[string]router.IpcHandler{
		
	})
	
	//startup
	tcpServer.Run()
	
	
	
	<-exitChan
}


