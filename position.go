package openf1go

// Provides driver positions throughout a session, including initial placement and subsequent changes.

import (
	"encoding/json"
	"strconv"
	"time"
)

// Base endpoint for position-related API calls
const positionBase = "/position"

// PostionsResponse represents a slice of Position objects returned by the API
type PostionsResponse []Position

// Position represents the structure of a single position record
type Position struct {
	Date         time.Time `json:"date"`          // Date of the position record
	DriverNumber int       `json:"driver_number"` // Unique identifier for the driver
	MeetingKey   int       `json:"meeting_key"`   // Key identifying the meeting
	Position     int       `json:"position"`      // Position of the driver
	SessionKey   int       `json:"session_key"`   // Key identifying the session
}

// getPositionsURL constructs the full URL for position-related API calls
func (c *Client) getPositionsURL() string {
	return c.baseUrl + positionBase
}

// GetPositions fetches position data for a specific position filter
func (c *Client) GetPositions(position Position) (PostionsResponse, error) {
	var positionsResponse PostionsResponse

	// Build the URL with query parameters based on the provided position filter
	url, err := UrlBuilder(c.getPositionsURL(), buildArgs(position))
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the positionsResponse struct
	if err := json.Unmarshal(resp, &positionsResponse); err != nil {
		return nil, err
	}

	return positionsResponse, nil
}

// GetAllDriversLatestPositions fetches the latest positions for all drivers
func (c *Client) GetAllDriversLatestPositions() (PostionsResponse, error) {
	var positionsResponse PostionsResponse

	// Build the URL with query parameters for the latest session
	url, err := UrlBuilder(c.getPositionsURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the positionsResponse struct
	if err := json.Unmarshal(resp, &positionsResponse); err != nil {
		return nil, err
	}

	return positionsResponse, nil
}

// GetDriversLatestPositions fetches the latest positions for a specific driver
func (c *Client) GetDriversLatestPositions(driver Driver) (PostionsResponse, error) {
	var positionsResponse PostionsResponse

	// Build the URL with query parameters for the latest session and the specific driver
	url, err := UrlBuilder(c.getPositionsURL(), append(c.getLatestSessionArgs(), Arg{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}))
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the positionsResponse struct
	if err := json.Unmarshal(resp, &positionsResponse); err != nil {
		return nil, err
	}

	return positionsResponse, nil
}
