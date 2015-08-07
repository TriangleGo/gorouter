package handlerImpl

import (
	"fmt"
	"gorouter/types"
)

type DisconHandlerImpl struct {
}

func (this *DisconHandlerImpl) Handle(client *types.Client) {
	fmt.Printf("TODO: i am the handler \n")
}
