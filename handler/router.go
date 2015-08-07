package handler

type Router struct {
	ConnHandler    ConnectHandler
	DisconnHandler DisconnectHandler
	TcpHandler     map[uint8]Handler
	Ipchandler     map[uint8]Handler
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

func (this *Router) SetIpcHandler(p map[uint8]Handler) {
	this.Ipchandler = p
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

func (this *Router) GetIpcHandler() map[uint8]Handler {
	return this.Ipchandler
}
