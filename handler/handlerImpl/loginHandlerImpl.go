package handlerImpl

import (
	"fmt"
	"gorouter/types"
)

type LoginHandlerImpl struct {
}

func (this *LoginHandlerImpl) Init()  {
	fmt.Printf("LoginHandleImple loaded \n")
}

func (this *LoginHandlerImpl) Handle(client *types.Client,data []byte) *types.Client {
	fmt.Printf("TODO: i am the LoginHandlerImpl data =  %v \n",data)
	return nil
}
