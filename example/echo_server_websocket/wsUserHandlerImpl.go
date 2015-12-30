package main

import (
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/client"
	"github.com/TriangleGo/gorouter/util"
)

type WsUserHandlerImpl struct {
}

func (this *WsUserHandlerImpl) Init()  {
	logger.Info("LoginHandleImple loaded \n")
}

func (this *WsUserHandlerImpl) Handle(client *client.Client,command string,data string) *client.Client {
	defer util.TraceCrashStack()
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
