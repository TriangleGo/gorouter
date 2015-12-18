package router

type Router struct {
	ConnHandler    ConnectHandler
	DisconnHandler DisconnectHandler
	TcpHandler     map[uint8]Handler
	IpcHandler     map[string]IpcHandler
	/*WebSocket handler*/
	WsHandler     map[string]WSHandler
}

func NewRouter() *Router {
	return &Router{}
}

func (this *Router) SetConnHandler(p ConnectHandler) {
	this.ConnHandler = p
}

func (this *Router) SetDisconHandler(p DisconnectHandler) {
	this.DisconnHandler = p
}

func (this *Router) SetTcpHandler(p map[uint8]Handler) {
	this.TcpHandler = p
}

func (this *Router) SetWsHandler(p map[string]WSHandler) {
	this.WsHandler = p
}

func (this *Router) SetIpcHandler(p map[string]IpcHandler) {
	this.IpcHandler = p
}

func (this *Router) GetConnHandler() ConnectHandler {
	return this.ConnHandler
}

func (this *Router) GetDisconHandler() DisconnectHandler {
	return this.DisconnHandler
}

func (this *Router) GetTcpHandler() map[uint8]Handler {
	return this.TcpHandler
}

func (this *Router) GetWsHandler() map[string]WSHandler {
	return this.WsHandler
}

func (this *Router) GetIpcHandler() map[string]IpcHandler {
	return this.IpcHandler
}

func (this *Router) Init()  {
	for _,v := range this.TcpHandler {
		v.Init()
	}
	for _,v := range this.IpcHandler {
		v.Init()
	}
}
