package main

import (
	"time"
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/client"
	"github.com/TriangleGo/gorouter/util"
)

type ConnHandlerImpl struct {
}

func (this *ConnHandlerImpl) Handle(client *client.Client, ch chan []byte) *client.Client {
	go util.TraceCrashStack()
	logger.Info("TODO: i am the conn handler \n")
	select {
	case res := <-ch:
		//it handle the data like flash sandbox and other
		logger.Info("TODO: ConnHandlerImpl first data %v \n", string(res))
	case <-time.After(time.Second * 60):
		logger.Info("when somebody connect but no packs to send , this operation will disconnect it\n")
		//client.GetConn().Close()
	}
	return nil
}
