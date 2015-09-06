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
)

//global
var (
	exitChan chan bool
)

func main() {
	fmt.Printf("Go router running \r\n")
	fmt.Printf("Go run TCPServer \r\n")
	server.NewTCPServer("tcp", ":9090").Run()
	fmt.Printf("Go run HTTPServer \r\n")
	server.NewHTTPServer(":9091")
	fmt.Printf("Go run WSServer \r\n")
	server.NewWSServer(":9092")
	
	a ,_:= net.Dial("tcp","127.0.0.1:9090")
	fmt.Printf(" remote %v\n",a.RemoteAddr())
	socket := network.NewBaseSocket(a)
	socket.LocalAddr()
	socket.RemoteAddr()
	socket.SetDeadline(time.Now().Add( time.Second * 10))
	socket.Write([]byte("abc"))
	socket.Close()
	fmt.Printf(" remote %v\n",a.RemoteAddr())

	fmt.Printf("CpuProfile %v \n", string(runtime.CPUProfile()))
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
