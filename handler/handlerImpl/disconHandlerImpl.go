package handlerImpl

import (
	"gorouter/types"
	"gorouter/logger"
)

type DisconHandlerImpl struct {
}

func (this *DisconHandlerImpl) Handle(client *types.Client) {
	logger.Info("TODO: i am the handler \n")
}
