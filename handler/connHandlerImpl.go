package handler

import (
	"time"
	"gorouter/logger"
	"gorouter/types"
)

type ConnHandlerImpl struct {
}

func (this *ConnHandlerImpl) Handle(client *types.Client, ch chan []byte) *types.Client {
	logger.Info("TODO: i am the conn handler \n")
	/*client.GetConn().Write([]byte{0,0,0,2,0,0})*/
	select {
	case res := <-ch:
		//it handle the data like flash sandbox and other

		logger.Info("TODO: ConnHandlerImpl first data %v \n", string(res))
/*
		client.GetConn().Write([]byte{0,0,0,6,0,2})
		time.Sleep(3 * time.Second)
		client.GetConn().Write([]byte{0,3})
		time.Sleep(3 * time.Second)
		client.GetConn().Write([]byte{0,4})
*/
	case <-time.After(time.Second * 60):
		logger.Info("when somebody connect but pack to send , this operation will disconnect it\n")
		//client.GetConn().Close()
	}
	return nil
}
