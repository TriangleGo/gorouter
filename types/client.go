package types

import (
	"net"
)

type Client struct {
	ClientID int
	Token    string
	Conn     net.Conn
}

func NewClient(c net.Conn) *Client {
	return &Client{Conn: c}
}

func (this *Client) GetConn() net.Conn {
	return this.Conn
}
