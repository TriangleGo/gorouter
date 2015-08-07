package main

import (
	"fmt"
	"gorouter/network"
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
	connManager *network.ConnectionManager
}

func NewTCPServer(_protocol string, _dsn string) *TCPServer {
	return &TCPServer{protocol: _protocol,
		dsn:         _dsn,
		connManager: network.NewConnectionManager()}
}

func (this *TCPServer) ServerListen() error {
	var err error
	this.listener, err = net.Listen(this.protocol, this.dsn)
	if err != nil {
		fmt.Printf("TCPServer Start Failed %v \r\n", err)
		return err
	}
	return nil
}

func (this *TCPServer) ServerAccpet() {
	//循环接收Accept
	for {
		conn, err := this.listener.Accept()
		if err != nil {
			fmt.Printf("TCPServer Accepting Failed %v \r\n", err)
			time.Sleep(time.Second * 10)
			continue
		}
		//接收数据
		this.connManager.Produce(&conn).Serve()
	}

}

func (this *TCPServer) Run() {
	if this.ServerListen() != nil {
		return
	}
	go this.ServerAccpet()
}
