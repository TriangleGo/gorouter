package router

import (

)

var _global_router *Router

func InitRouter() *Router {
	if _global_router != nil {
		return _global_router
	}
	_global_router = NewRouter()
	
	//init all module	
	GetRouter().Init()

	return _global_router

}

func GetRouter() *Router {
	if _global_router == nil {
		_global_router = NewRouter()
	}
	return _global_router
}
