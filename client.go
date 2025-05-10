package openf1go

import (
	"net/http"
	"time"
)

type Client struct {
	client  http.Client
	baseUrl string
}

func New() *Client {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	return &Client{
		client:  c,
		baseUrl: "https://api.openf1.org/v1",
	}
}

func (c *Client) getLatestSessionArgs() []Arg {
	return []Arg{{Key: "meeting_key", Value: "latest"}, {Key: "session_key", Value: "latest"}}
}
