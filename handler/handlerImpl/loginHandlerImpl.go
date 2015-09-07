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

func (this *LoginHandlerImpl) Handle(client *types.Client) *types.Client {
	fmt.Printf("TODO: i am the handler \n")
	return nil
}
