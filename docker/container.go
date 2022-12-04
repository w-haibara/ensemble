package docker

import (
	"context"
	"errors"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func (c Client) ContainerRun(ctx context.Context, r Request) error {
	type Req struct {
		ImageName string
	}
	req := Req{}
	if err := r.Unmarshal(&req); err != nil {
		return err
	}

	if req.ImageName == "" {
		return errors.New("ImageName is needed")
	}

	resp, err := c.docker.ContainerCreate(ctx, &container.Config{
		Image: req.ImageName,
	}, nil, nil, nil, "")
	if err != nil {
		return err
	}

	if err := c.docker.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerStart(ctx context.Context, r Request) error {
	type Req struct {
		ContainerID string
	}
	req := Req{}
	if err := r.Unmarshal(&req); err != nil {
		return err
	}

	if req.ContainerID == "" {
		return errors.New("ContainerID is needed")
	}

	if err := c.docker.ContainerStart(ctx, req.ContainerID, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerStop(ctx context.Context, r Request) error {
	type Req struct {
		ContainerID string
		Timeout     string
	}
	req := Req{}
	if err := r.Unmarshal(&req); err != nil {
		return err
	}

	if req.ContainerID == "" {
		return errors.New("ContainerID is needed")
	}

	var timeout time.Duration
	if req.Timeout != "" {
		var err error
		timeout, err = time.ParseDuration(req.Timeout)
		if err != nil {
			return err
		}
	}

	if err := c.docker.ContainerStop(ctx, req.ContainerID, &timeout); err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerKill(ctx context.Context, r Request) error {
	type Req struct {
		ContainerID string
		Signal      string
	}
	req := Req{}
	if err := r.Unmarshal(&req); err != nil {
		return err
	}

	if req.ContainerID == "" {
		return errors.New("ContainerID is needed")
	}

	if err := c.docker.ContainerKill(ctx, req.ContainerID, req.Signal); err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerRemove(ctx context.Context, r Request) error {
	type Req struct {
		ContainerID string
	}
	req := Req{}
	if err := r.Unmarshal(&req); err != nil {
		return err
	}

	if req.ContainerID == "" {
		return errors.New("ContainerID is needed")
	}

	if err := c.docker.ContainerRemove(ctx, req.ContainerID, types.ContainerRemoveOptions{}); err != nil {
		return err
	}

	return nil
}
