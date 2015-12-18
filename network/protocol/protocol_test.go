package protocol

import (
	"fmt"

	"testing"

)



func Test_Protocol(t *testing.T) {
	p := NewProtocalByParams(1,2,[]byte("1234567890"))	
	b := p.ToBytes()
	p.GetCommandId()
	p.GetModuleId()
	fmt.Printf("Tobytes %v\nCommandID %v\nModuleId %v \n",b,p.GetCommandId(),p.GetModuleId())
	
	p2 := NewProtocal()
	p2.PraseFromData(b,len(b))
	p2.ToBytes()
	p2.GetCommandId()
	p2.GetModuleId()
	fmt.Printf("p2 bytes: %v \ncid: %v \nmid:%v\n",p2.ToBytes(),p2.Command,p2.GetModuleId())
	
	wsString := `{"module":"user","command":"log","data":{"uid":1234,"pwd":"123456"}}`
	ws,err := NewWsProtocolFromData([]byte(wsString))
	fmt.Printf("err %v\nbytes %v \n",err,ws.toBytes())
	fmt.Printf("module:%v\ncommand:%v\n",ws.Module,ws.Command)
	fmt.Printf("data %v \n",string(ws.Data))
	
	w := NewWsProtocolFromParams("hahah","helo",`{"test:":1234}`)
	fmt.Printf("Module :%v\nCommand:%v\ndata:%v\n",w.Module,w.Command,string(w.Data))
	fmt.Printf("bytes :%v\n",string(w.toBytes()))
}
