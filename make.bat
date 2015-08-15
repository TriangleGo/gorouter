set GOARCH=amd64
set GOOS=windows
go run gorouter.go httpServer.go tcpServer.go

pause