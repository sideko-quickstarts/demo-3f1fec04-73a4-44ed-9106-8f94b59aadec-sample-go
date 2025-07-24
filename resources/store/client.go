package store

import (
	http "net/http"
	sdkcore "sample_go/core"
	order "sample_go/resources/store/order"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	Order      *order.Client
}
type RequestModifier = func(req *http.Request) error

// Instantiate a new resource client
func NewClient(coreClient *sdkcore.CoreClient) *Client {
	client := Client{
		coreClient: coreClient,
		Order:      order.NewClient(coreClient),
	}

	return &client
}
