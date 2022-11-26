package daemon

import (
	"context"
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
	if _, err := c.Post("http://unix/"+name, "", nil); err != nil {
		return err
	}

	return nil
}
