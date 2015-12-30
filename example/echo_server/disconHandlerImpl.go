package main

import (
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/client"
)

type DisconHandlerImpl struct {
}

func (this *DisconHandlerImpl) Handle(client *client.Client) {
	logger.Info("TODO: i am the handler \n")
}
