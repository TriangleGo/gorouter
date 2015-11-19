set GOARCH=386
set GOOS=windows
rem set GODEBUG=gctrace=1 
go build gorouter.go
go  run gorouter.go 

pause