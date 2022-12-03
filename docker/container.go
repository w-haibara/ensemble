package docker

import (
	"context"
)

func (c Client) ContainerRun(ctx context.Context, r Request) error {
	return nil
}

func (c Client) ContainerStart(ctx context.Context, r Request) error {
	return nil
}

func (c Client) ContainerStop(ctx context.Context, r Request) error {
	if err := c.docker.ContainerStop(ctx, "", nil); err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerRemove(ctx context.Context, r Request) error {
	return nil
}
