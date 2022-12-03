package docker

import (
	doc "github.com/docker/docker/client"
)

type Client struct {
	docker *doc.Client
}

func NewClient() (*Client, error) {
	cli, err := doc.NewClientWithOpts(
		doc.FromEnv,
		doc.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}

	return &Client{cli}, nil
}
