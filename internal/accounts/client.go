package accounts

import "github.com/kpurdon/apir/pkg/requester"

type Client struct {
	api requester.Requester
}

func NewClient(api requester.Requester) *Client {
	return &Client{api: api}
}
