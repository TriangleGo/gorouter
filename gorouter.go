package main

import (
	"fmt"
	//. "gorouter/util"
	"os"
	"runtime"
	"runtime/pprof"
)

//global
var (
	exitChan chan bool
)

func main() {
	fmt.Printf("Go router running \r\n")
	fmt.Printf("Go run TCPServer \r\n")
	NewTCPServer("tcp", ":9090").Run()
	fmt.Printf("Go run HTTPServer \r\n")
	go HTTPServer()
	fmt.Printf("Go run WSServer \r\n")
	go WSServer()

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
