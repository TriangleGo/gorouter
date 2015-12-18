package handler

import (
	"gorouter/logger"
	"gorouter/client"
	"gorouter/util"
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
		case "login":
			//TODO:
			break
		default:
			//TODO:
			break
	}
	return nil
}
