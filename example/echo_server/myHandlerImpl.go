package main

import (
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/client"
	"github.com/TriangleGo/gorouter/network/protocol"
	"github.com/TriangleGo/gorouter/util"
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
