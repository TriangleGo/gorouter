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

	//setup handler
	GetRouter().SetConnHandler(&handlerImpl.ConnHandlerImpl{})
	GetRouter().SetDisconHandler(&handlerImpl.DisconHandlerImpl{})

	Dispatchs := map[uint8]Handler{
		0: &handlerImpl.LoginHandlerImpl{},
	}
	
	for _,v := range Dispatchs {
		//init all module
		v.Init()
	}
	
	GetRouter().SetTcpHandler(Dispatchs)
	GetRouter().SetIpcHandler(Dispatchs)

}

func GetRouter() *Router {
	if _router == nil {
		InitRouter()
	}
	return _router
}
