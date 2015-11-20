package router

import (
	"gorouter/handler"
)

var _router *Router

func InitRouter() {
	if _router != nil {
		return
	}
	_router = NewRouter()

	//setup handler
	GetRouter().SetConnHandler(&handler.ConnHandlerImpl{})
	GetRouter().SetDisconHandler(&handler.DisconHandlerImpl{})

	Dispatchs := map[uint8]Handler{
		0: &handler.LoginHandlerImpl{},
	}
	
	for _,v := range Dispatchs {
		//init all module
		v.Init()
	}
	
	IpcDispatchs := map[uint8]IpcHandler{
		
	}
	
	GetRouter().SetTcpHandler(Dispatchs)
	GetRouter().SetIpcHandler(IpcDispatchs)

}

func GetRouter() *Router {
	if _router == nil {
		InitRouter()
	}
	return _router
}
