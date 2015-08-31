set GOARCH=386
set GOOS=windows
go run gorouter.go httpServer.go tcpServer.go wsServer.go

pause
