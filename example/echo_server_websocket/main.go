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
	
	
	logger.Info("Go run WSServer \r\n")
	
	wsServer := server.NewWSServer(":9092")
	//register handler
	wsServer.GetRouter().SetWsHandler(map[string]router.WSHandler{
		"system": &WsUserHandlerImpl{},
		"user": &WsUserHandlerImpl{},
		"chat": &WsUserHandlerImpl{},
	})
	
	//startup
	wsServer.Run()
	
	<-exitChan
}


