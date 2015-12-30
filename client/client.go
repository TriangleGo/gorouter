package client

import (
	"github.com/TriangleGo/gorouter/network/socket"
	"github.com/TriangleGo/gorouter/network/protocol"
)

type Client struct {
	ClientID int
	Token    string
	conn     *socket.BaseSocket

}

func NewClient(c *socket.BaseSocket) *Client {
	return &Client{conn: c}
}

func (this *Client) GetConn() *socket.BaseSocket {
	return this.conn
}

func (this *Client) Send(CMD1 uint8,CMD2 uint8,Data []byte) {
	dat := protocol.NewProtocalByParams(CMD1,CMD2,Data).ToBytes()
	this.conn.Write(dat)
}

func (this *Client) WsSend( module,command,data string ) {
	wp := protocol.NewWsProtocolFromParams(module,command,data)
	this.GetConn().Write(wp.ToBytes())
}



