package client

type Client struct {
	address string
}

func New(address string) *Client {
	return &Client{
		address: address,
	}
}
