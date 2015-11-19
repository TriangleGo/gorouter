package network

import (
)

type Client struct {
	ClientID int
	Token    string
	conn     BaseSocket

}

func NewClient(c BaseSocket) *Client {
	return &Client{conn: c}
}

func (this *Client) GetConn() BaseSocket {
	return this.conn
}

func (this *Client) Send(CMD1 uint8,CMD2 uint8,Data []byte) {
	
}
