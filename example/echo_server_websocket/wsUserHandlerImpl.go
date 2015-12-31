package main

import (
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/client"
)

type WsUserHandlerImpl struct {
}

func (this *WsUserHandlerImpl) Init()  {
	logger.Info("LoginHandleImple loaded \n")
}

func (this *WsUserHandlerImpl) Handle(client *client.Client,command string,data string) *client.Client {
	logger.Info("TODO: WsUserHandlerImpl data =  %v \n",data)
	switch(command) {
		case "echo":
			//TODO:
			logger.Info("echo server \n")
			client.WsSend("user","echo",data)
			break
		default:
			//TODO:
			break
	}
	return nil
}
