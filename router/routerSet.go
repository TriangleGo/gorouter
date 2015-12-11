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

	//register handler
	Dispatchs := map[uint8]Handler{
		0: &handler.LoginHandlerImpl{},
	}
	// register ipc handler
	IpcDispatchs := map[uint8]IpcHandler{
		
	}
	//init all module	
	for _,v := range Dispatchs {
		v.Init()
	}
	for _,v := range IpcDispatchs {
		v.Init()
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
