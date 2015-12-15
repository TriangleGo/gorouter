package handler

import (
	"gorouter/logger"
	"gorouter/client"
	"gorouter/network/protocol"
	"gorouter/util"
)

type LoginHandlerImpl struct {
}

func (this *LoginHandlerImpl) Init()  {
	logger.Info("LoginHandleImple loaded \n")
}

func (this *LoginHandlerImpl) Handle(client *client.Client,data *protocol.Protocol) *client.Client {
	defer util.TraceCrashStack()
	logger.Info("TODO: LoginHandlerImpl data =  %v \n",data)
	switch(data.Command) {
		case 0:
			//TODO:
			break
		default:
			//TODO:
			break
	}
	return nil
}
