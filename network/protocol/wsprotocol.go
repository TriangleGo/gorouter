package protocol

import (
	"github.com/bitly/go-simplejson"
	"github.com/TriangleGo/gorouter/logger"
)


type WsProtocol struct {
	Module string
	Command string
	Data []byte
}


func NewWsProtocol() *WsProtocol{
	return &WsProtocol{}
}

func NewWsProtocolFromParams(m,c ,data string) *WsProtocol{
	return &WsProtocol{Module:m,Command:c,Data: []byte(data)}
}

func NewWsProtocolFromData(data []byte) (*WsProtocol,error) {
	json,err := simplejson.NewJson(data)
	if err != nil {
		logger.Info("Parsing data error:%v\n",err)
		return nil,err 
	}
	
	m := json.Get("module").MustString()
	c := json.Get("command").MustString()
	d,_ := json.Get("data").MarshalJSON()
	
	wp := NewWsProtocol()
	wp.Module = m
	wp.Command = c
	wp.Data = d
	
	return wp,nil
}

func (this *WsProtocol) ToBytes() []byte {
	dataJson,_ := simplejson.NewJson(this.Data)
	
	json := simplejson.New()
	json.Set("module",this.Module)
	json.Set("command",this.Command)
	json.Set("data",dataJson)
	b , _ := json.MarshalJSON()
	return b
}