package handler

import (
	"gorouter/logger"
	"gorouter/types"
)

type LoginHandlerImpl struct {
}

func (this *LoginHandlerImpl) Init()  {
	logger.Info("LoginHandleImple loaded \n")
}

func (this *LoginHandlerImpl) Handle(client *types.Client,data []byte) *types.Client {
	logger.Info("TODO: i am the LoginHandlerImpl data =  %v \n",data)
	return nil
}
