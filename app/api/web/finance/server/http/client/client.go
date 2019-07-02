package client

import (
	"mall/lib/net/http"
)

type Client struct {
	c *http.Client
}

func New(host string) *Client {
	return &Client{
		c: http.NewClient(host),
	}
}
