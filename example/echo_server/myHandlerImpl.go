package main

import (
	"gorouter/logger"
	"gorouter/client"
	"gorouter/network/protocol"
	"gorouter/util"
)

type MyHandlerImpl struct {
}

func (this *MyHandlerImpl) Init()  {
	logger.Info("LoginHandleImple loaded \n")
}

func (this *MyHandlerImpl) Handle(client *client.Client,data *protocol.Protocol) *client.Client {
	defer util.TraceCrashStack()
	logger.Info("TODO: MyHandlerImpl data =  %v \n",data)
	switch(data.Command) {
		case 0:
			//TODO:
			client.Send(data.ModuleId,data.Command,data.Data)
			break
		default:
			//TODO:
			break
	}
	return nil
}
