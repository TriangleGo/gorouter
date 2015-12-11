package handler

import (
	"gorouter/logger"
	"gorouter/handler/client"
	"gorouter/network/protocol"
)

type LoginHandlerImpl struct {
}

func (this *LoginHandlerImpl) Init()  {
	logger.Info("LoginHandleImple loaded \n")
}

func (this *LoginHandlerImpl) Handle(client *client.Client,data *protocol.Protocol) *client.Client {
	logger.Info("TODO: i am the LoginHandlerImpl data =  %v \n",data)
	return nil
}
