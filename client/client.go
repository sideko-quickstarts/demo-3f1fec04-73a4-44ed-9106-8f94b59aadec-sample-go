package client

import (
	sdkcore "sample_go/core"
	pet "sample_go/resources/pet"
	store "sample_go/resources/store"
)

type Client struct {
	coreClient *sdkcore.CoreClient
	Pet        *pet.Client
	Store      *store.Client
}

// Instantiate a new API client
func NewClient(builders ...func(*sdkcore.CoreClient)) *Client {
	defaultEnv := sdkcore.DefaultBaseURL(Environment.String())
	coreClient := sdkcore.NewCoreClient(defaultEnv)
	for _, b := range builders {
		b(coreClient)
	}

	client := Client{
		coreClient: coreClient,
		Pet:        pet.NewClient(coreClient),
		Store:      store.NewClient(coreClient),
	}

	return &client
}
