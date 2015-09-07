package network


import (
	"time"
	"net"
	"fmt"
	"golang.org/x/net/websocket"
	"reflect"
)

/*	 rewrite method 
	net.Conn.LocalAddr()
	net.Conn.RemoteAddr()
	net.Conn.SetDeadline()
	net.Conn.SetReadDeadline()
	net.Conn.SetWriteDeadline()
	net.Conn.Read()
	net.Conn.Write()
*/

type BaseSocket struct {
	/* websocket.Conn or net.Conn */
	socket interface{}
	socktype string
}


func NewBaseSocket(conn interface{}) *BaseSocket{
	fmt.Printf("reflect.TypeOf(conn) %v \n",reflect.TypeOf(conn).String())
	var stype string
	switch  conn.(type)  {
		case  *net.Conn:
			stype = "socket"
			break
		case *websocket.Conn:
			stype = "websocket"
			break
		default :
			fmt.Printf("unknow\n")
	}

	return &BaseSocket{socket:conn,socktype:stype}
}


func (this *BaseSocket) Close() error{
	in := make([]reflect.Value,0)
	ret := reflect.ValueOf(this.socket).MethodByName("Close").Call(in)
	fmt.Printf("ret = %v \n",ret[0].Interface())
	return nil
}

func (this *BaseSocket) LocalAddr() net.Addr{
	in := make([]reflect.Value,0)
	ret := reflect.ValueOf(this.socket).MethodByName("LocalAddr").Call(in)
	//fmt.Printf("ret = %v \n",ret[0].Interface())
	return (ret[0].Interface()).(net.Addr)

}


func (this *BaseSocket) RemoteAddr() net.Addr{
	in := make([]reflect.Value,0)
	ret := reflect.ValueOf(this.socket).MethodByName("RemoteAddr").Call(in)
	//fmt.Printf("ret = %v \n",ret[0].Interface())
	return (ret[0].Interface()).(net.Addr)
}


func (this *BaseSocket) SetDeadline(t time.Time) error {
	in := []reflect.Value{reflect.ValueOf(t)}
	ret := reflect.ValueOf(this.socket).MethodByName("SetDeadline").Call(in)
	//fmt.Printf("ret = %v \n",ret[0].Interface()) 
	
	if ret[0].IsNil() {
		return nil 
	}
	return (ret[0].Interface()).(error)
}

func (this *BaseSocket) SetReadDeadline(t time.Time) error {
	in := []reflect.Value{reflect.ValueOf(t)}
	ret := reflect.ValueOf(this.socket).MethodByName("SetReadDeadline").Call(in)
	//fmt.Printf("ret = %v \n",ret[0].Interface())
	if ret[0].IsNil() {
		return nil 
	}
	return (ret[0].Interface()).(error)
}

func (this *BaseSocket) SetWriteDeadline(t time.Time) error {
	in := []reflect.Value{reflect.ValueOf(t)}
	ret := reflect.ValueOf(this.socket).MethodByName("SetWriteDeadline").Call(in)
	//fmt.Printf("ret = %v \n",ret[0].Interface())
	if ret[0].IsNil() {
		return nil 
	}
	return (ret[0].Interface()).(error)
}


func (this *BaseSocket) Read(b []byte) (n int, err error){
	in := []reflect.Value{reflect.ValueOf(b)}
	ret := reflect.ValueOf(this.socket).MethodByName("Read").Call(in)
	//fmt.Printf("ret = %v %v \n",ret[0].Interface(),ret[1].Interface())
	if ret[1].IsNil() {
		return (ret[0].Interface()).(int),nil 
	}
	return (ret[0].Interface()).(int),(ret[1].Interface()).(error)
}


func (this *BaseSocket) Write(b []byte) (n int, err error){
	in := []reflect.Value{reflect.ValueOf(b)}
	ret := reflect.ValueOf(this.socket).MethodByName("Write").Call(in)
	//fmt.Printf("ret = %v %v \n",ret[0].Interface(),ret[1].Interface())
	if ret[1].IsNil() {
		return (ret[0].Interface()).(int),nil 
	}
	return (ret[0].Interface()).(int),(ret[1].Interface()).(error)
}