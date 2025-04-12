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
