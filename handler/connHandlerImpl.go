package handler

import (
	"time"
	"gorouter/logger"
	"gorouter/client"
	"gorouter/util"
)

type ConnHandlerImpl struct {
}

func (this *ConnHandlerImpl) Handle(client *client.Client, ch chan []byte) *client.Client {
	go util.TraceCrashStack()
	logger.Info("TODO: i am the conn handler \n")
	/*client.GetConn().Write([]byte{0,0,0,2,0,0})*/
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
