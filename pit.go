package openf1go

// Provides information about cars going through the pit lane.

import (
	"encoding/json"
	"strconv"
	"time"
)

// Base endpoint for pit-related API calls
const pitBase = "/pit"

// PitResponse represents a collection of pit data
type PitResponse []Pit

// Pit represents the details of a single pit stop
type Pit struct {
	Date         time.Time `json:"date"`          // Date of the pit stop
	DriverNumber int       `json:"driver_number"` // Driver's unique number
	LapNumber    int       `json:"lap_number"`    // Lap number when the pit stop occurred
	MeetingKey   int       `json:"meeting_key"`   // Identifier for the meeting/session
	PitDuration  float64   `json:"pit_duration"`  // Duration of the pit stop in seconds
	SessionKey   int       `json:"session_key"`   // Identifier for the session
}

// getPitURL constructs the full URL for pit-related API calls
func (c *Client) getPitURL() string {
	// Combine the base URL of the client with the pit endpoint
	return c.baseUrl + pitBase
}

// GetPits fetches pit data based on the provided Pit struct
func (c *Client) GetPits(pit Pit) (PitResponse, error) {
	var pitResponse PitResponse

	// Build the URL with query parameters based on the Pit struct
	url, err := UrlBuilder(c.getPitURL(), buildArgs(pit))
	if err != nil {
		return nil, err // Return an error if URL construction fails
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err // Return an error if the HTTP request fails
	}

	// Parse the JSON response into the PitResponse struct
	if err := json.Unmarshal(resp, &pitResponse); err != nil {
		return nil, err // Return an error if JSON unmarshalling fails
	}

	return pitResponse, nil // Return the parsed pit data
}

// GetAllDriversLatestPits fetches the latest pit data for all drivers
func (c *Client) GetAllDriversLatestPits() (PitResponse, error) {
	var pitResponse PitResponse

	// Build the URL with query parameters for the latest session
	url, err := UrlBuilder(c.getPitURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err // Return an error if URL construction fails
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err // Return an error if the HTTP request fails
	}

	// Parse the JSON response into the PitResponse struct
	if err := json.Unmarshal(resp, &pitResponse); err != nil {
		return nil, err // Return an error if JSON unmarshalling fails
	}

	return pitResponse, nil // Return the parsed pit data
}

// GetDriverLatestPits fetches the latest pit data for a specific driver
func (c *Client) GetDriverLatestPits(driver Driver) (PitResponse, error) {
	var pitResponse PitResponse

	// Build the URL with query parameters for the latest session and specific driver
	url, err := UrlBuilder(c.getPitURL(), append(c.getLatestSessionArgs(), Arg{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}))
	if err != nil {
		return nil, err // Return an error if URL construction fails
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err // Return an error if the HTTP request fails
	}

	// Parse the JSON response into the PitResponse struct
	if err := json.Unmarshal(resp, &pitResponse); err != nil {
		return nil, err // Return an error if JSON unmarshalling fails
	}

	return pitResponse, nil // Return the parsed pit data
}
