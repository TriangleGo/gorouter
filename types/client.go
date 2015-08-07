package types

type Client struct {
	ClientID int
	Token    string
}

func NewClient() *Client {
	return &Client{}
}
