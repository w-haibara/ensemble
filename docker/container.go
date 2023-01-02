package docker

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/w-haibara/docker-wrapper/docker"
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

	cmd := docker.DockerContainerRunCmd(
		docker.DockerContainerRunOption{Name: &req.ImageName},
		[]string{},
	)
	out, err := ExecCmd(cmd)
	log.Println("ExecCmd: ContainerRun\n", out)
	if err != nil {
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

	cmd := docker.DockerContainerStartCmd(
		docker.DockerContainerStartOption{},
		[]string{req.ContainerID},
	)
	out, err := ExecCmd(cmd)
	log.Println("ExecCmd: ContainerStart\n", out)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerStop(ctx context.Context, r Request) error {
	type Req struct {
		ContainerID string
		Time        string
	}
	req := Req{}
	if err := r.Unmarshal(&req); err != nil {
		return err
	}

	if req.ContainerID == "" {
		return errors.New("ContainerID is needed")
	}

	time, err := func() (*int, error) {
		t, err := strconv.Atoi(req.Time)
		if err != nil {
			return nil, err
		}

		return &t, nil
	}()
	if err != nil {
		return err
	}

	cmd := docker.DockerContainerStopCmd(
		docker.DockerContainerStopOption{Time: time},
		[]string{req.ContainerID},
	)
	out, err := ExecCmd(cmd)
	log.Println("ExecCmd: ContainerStop\n", out)
	if err != nil {
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

	cmd := docker.DockerContainerKillCmd(
		docker.DockerContainerKillOption{Signal: &req.Signal},
		[]string{req.ContainerID},
	)
	out, err := ExecCmd(cmd)
	log.Println("ExecCmd: ContainerKill\n", out)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) ContainerRm(ctx context.Context, r Request) error {
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

	cmd := docker.DockerContainerRmCmd(
		docker.DockerContainerRmOption{},
		[]string{req.ContainerID},
	)
	out, err := ExecCmd(cmd)
	log.Println("ExecCmd: ContainerRm\n", out)
	if err != nil {
		return err
	}

	return nil
}
