package handler

import (
	"gorouter/logger"
	"gorouter/types"
)

type DisconHandlerImpl struct {
}

func (this *DisconHandlerImpl) Handle(client *types.Client) {
	logger.Info("TODO: i am the handler \n")
}
