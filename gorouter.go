package main

import (
	"time"
	"net"
	"fmt"
	//. "gorouter/util"
	"os"
	"runtime"
	"runtime/pprof"
	"gorouter/server"
	"gorouter/network"
	"gorouter/logger"
)

//global
var (
	exitChan chan bool
)

func main() {
	logger.GetLogger().SetLogLevel(5,5)
	logger.GetLogger().SetOutputDir(`c:\`)
	logger.GetLogger().Init("testlog")
	
	logger.Info("Go router running \r\n")
	logger.Info("Go run TCPServer \r\n")
	server.NewTCPServer("tcp", ":9090").Run()
	logger.Info("Go run HTTPServer \r\n")
	server.NewHTTPServer(":9091").Run()
	logger.Info("Go run WSServer \r\n")
	server.NewWSServer(":9092").Run()
	
	a ,_:= net.Dial("tcp","127.0.0.1:9090")
	logger.Info(" remote %v\n",a.RemoteAddr())
	socket := network.NewBaseSocket(a)
	socket.LocalAddr()
	socket.RemoteAddr()
	socket.SetDeadline(time.Now().Add( time.Second * 10))
	socket.Write([]byte{0,0,0,4,0,0,0,1})
	socket.Close()
	logger.Info(" remote %v\n",a.RemoteAddr())

	logger.Info("CpuProfile %v \n", string(runtime.CPUProfile()))
	for {
		var cmd string
		fmt.Scanf("%s", &cmd)
		if cmd == "monitor" {
			pprof.StartCPUProfile(os.Stdout)
		}
	}
	<-exitChan
}

//自定义
