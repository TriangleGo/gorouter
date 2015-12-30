package server

import (
	"github.com/TriangleGo/gorouter/network/socket"
	"github.com/TriangleGo/gorouter/network"
	"github.com/TriangleGo/gorouter/router"
	"github.com/TriangleGo/gorouter/logger"
	"github.com/TriangleGo/gorouter/util"
	"net"
	"time"
)


//   tcp Server
type TCPServer struct {
	dsn         string //127.0.0.1:port
	protocol    string //tcp
	listener    net.Listener
	router	*router.Router
}

func NewTCPServer(_protocol string, _dsn string) *TCPServer {
	return &TCPServer{protocol: _protocol,
		dsn:         _dsn,
		router: router.NewRouter()}
}

func (this *TCPServer) ServerListen() error {
	var err error
	this.listener, err = net.Listen(this.protocol, this.dsn)
	if err != nil {
		logger.Info("TCPServer Start Failed %v \r\n", err)
		return err
	}
	return nil
}

func (this *TCPServer) GetRouter() *router.Router {
	return router.GetRouter()
}

func (this *TCPServer) Run() {
	if this.ServerListen() != nil {
		return
	}
	go this.ServerAccpet()
}


func (this *TCPServer) ServerAccpet() {
	defer util.TraceCrashStack()
	//Looping Accept
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			logger.Info("TCPServer Accepting Failed %v \r\n", err)
			time.Sleep(time.Second * 10)
			continue
		}
		// recv data
		network.GetConnectionManager().Produce(socket.NewBaseSocket(conn)).AsyncServe()
		//this.connManager.Produce(&conn).Serve()
	}

}