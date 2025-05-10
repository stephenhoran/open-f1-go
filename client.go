package openf1go

// Package openf1go provides a client for accessing the OpenF1 API.

import (
	"net/http"
	"time"
)

// Client is a struct that wraps an HTTP client and a base URL for API requests.
type Client struct {
	client  http.Client // HTTP client used to make requests
	baseUrl string      // Base URL for the API
}

// New creates and returns a new instance of the Client struct.
// It initializes the HTTP client with a timeout of 15 seconds and sets the base URL.
func New() *Client {
	c := http.Client{Timeout: time.Duration(15) * time.Second} // Set HTTP client timeout
	return &Client{
		client:  c,
		baseUrl: "https://api.openf1.org/v1", // Base URL for the OpenF1 API
	}
}

// getLatestSessionArgs returns a slice of Arg structs representing
// the query parameters for fetching the latest session data.
func (c *Client) getLatestSessionArgs() []Arg {
	return []Arg{
		{Key: "meeting_key", Value: "latest"}, // Query parameter for the latest meeting
		{Key: "session_key", Value: "latest"}, // Query parameter for the latest session
	}
}
