
## gorouter ##

A simple message server or a simple mesaage framework


## Features ##
* High Concurrence 
* Fit for vertical communities
* Flexible, Modular
* Support websocket
* Very easy to use

## Requirements

* Go 1.2 or higher
## Installation ##
```
go get "github.com/garyburd/redigo/redis"
go get "github.com/go-sql-driver/mysql"
go get "github.com/astaxie/beego/orm"
go get "github.com/bitly/go-simplejson"
go get "github.com/TriangleGo/gorouter"
```

## Protocol ##

#### Tcp/Socket  
>>   [Header:Int32] [ModuleId:Int8] [CommandId:Int8] [MsgBody:Binary]


>>
>> Example:     
>>  Header = len(ModuleId) + len(CommandId) + len(MsgBody)   
>>  [0,0,0,4][0]   [0]    [0,0]   


#### Tcp/WebSocket    
```
{ "moduleid":"user",   "commandid":"getstate",   "data":{"userid":12345}}
```


## Example 
> See gorouter/example     


	package main
	
	import (
		"github.com/TriangleGo/gorouter/router"
		"github.com/TriangleGo/gorouter/server"
		"github.com/TriangleGo/gorouter/logger"
	)
	
	//global
	var (
		exitChan chan bool
	)
	
	func main() {
		logger.Info("Go router running \r\n")
		//init logger
		logger.GetLogger().SetLogLevel(5,5)
		logger.GetLogger().SetOutputDir(`c:\`)
		logger.GetLogger().Init("testlog")
		
		//init tcp server
		logger.Info("Go run TCPServer \r\n")
		tcpServer := server.NewTCPServer("tcp", ":9093")
	
		//setup handler
		tcpServer.GetRouter().SetConnHandler(&ConnHandlerImpl{})
		tcpServer.GetRouter().SetDisconHandler(&DisconHandlerImpl{})
		//register handler
		tcpServer.GetRouter().SetTcpHandler(map[uint8]router.Handler{
			0x0: &MyHandlerImpl{},
		})
		//register Ipc handler
		tcpServer.GetRouter().SetIpcHandler(map[string]router.IpcHandler{
			
		})
		
		//startup
		tcpServer.Run()
		
		<-exitChan
	}
