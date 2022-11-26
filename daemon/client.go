package daemon

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
)

type Client http.Client

func NewClient() Client {
	return Client(http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", unixDomainSockPath)
			},
		},
	})
}

func (c *Client) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	return (*http.Client)(c).Post(url, contentType, body)
}

func (c *Client) Notify(name string) error {
	res, err := c.Post("http://unix/"+name, "", nil)
	if err != nil {
		return err
	}

	if res.StatusCode >= 400 {
		return fmt.Errorf("request failed with status: %s", res.Status)
	}

	return nil
}
