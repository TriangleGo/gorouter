package types

import (
	"net"
)

type Client struct {
	ClientID int
	Token    string
	conn     net.Conn

}

func NewClient(c net.Conn) *Client {
	return &Client{conn: c}
}

func (this *Client) GetConn() net.Conn {
	return this.conn
}
