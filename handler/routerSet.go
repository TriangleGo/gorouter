package handler

import (
	"gorouter/handler/handlerImpl"
)

var _router *Router

func InitRouter() {
	if _router != nil {
		return
	}
	_router = NewRouter()

	//装载handler
	GetRouter().SetConnHandler(&handlerImpl.ConnHandlerImpl{})
	GetRouter().SetDisconHandler(&handlerImpl.DisconHandlerImpl{})

	Dispatchs := map[uint8]Handler{
		0: &handlerImpl.LoginHandlerImpl{},
	}
	GetRouter().SetIpcHandler(Dispatchs)

}

func GetRouter() *Router {
	if _router == nil {
		InitRouter()
	}
	return _router
}
