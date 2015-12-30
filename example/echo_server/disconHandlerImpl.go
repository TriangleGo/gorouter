package main

import (
	"gorouter/logger"
	"gorouter/client"
)

type DisconHandlerImpl struct {
}

func (this *DisconHandlerImpl) Handle(client *client.Client) {
	logger.Info("TODO: i am the handler \n")
}
