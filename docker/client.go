package docker

import "os/exec"

type Client struct {
}

func NewClient() (*Client, error) {

	return &Client{}, nil
}

func ExecCmd(cmd *exec.Cmd) (string, error) {
	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), err
	}

	return "", nil
}
