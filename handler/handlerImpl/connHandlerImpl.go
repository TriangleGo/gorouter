package handlerImpl

import (
	"fmt"
	"gorouter/types"
)

type ConnHandlerImpl struct {
}

func (this *ConnHandlerImpl) Handle(client *types.Client) *types.Client {
	fmt.Printf("TODO: i am the handler \n")

	return nil
}
