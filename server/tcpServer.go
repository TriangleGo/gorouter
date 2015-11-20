package server

import (
	"gorouter/network/socket"
	"gorouter/network"
	"gorouter/logger"
	"net"
	"time"
)

//**************************************
//   tcp 服务器
//**************************************

type TCPServer struct {
	dsn         string //127.0.0.1:port
	protocol    string //tcp
	listener    net.Listener
	
}

func NewTCPServer(_protocol string, _dsn string) *TCPServer {
	return &TCPServer{protocol: _protocol,
		dsn:         _dsn}
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

func (this *TCPServer) ServerAccpet() {
	//循环接收Accept
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			logger.Info("TCPServer Accepting Failed %v \r\n", err)
			time.Sleep(time.Second * 10)
			continue
		}
		//接收数据
		network.GetConnectionManager().Produce(socket.NewBaseSocket(conn)).AsyncServe()
		//this.connManager.Produce(&conn).Serve()
	}

}

func (this *TCPServer) Run() {
	if this.ServerListen() != nil {
		return
	}
	go this.ServerAccpet()
}
