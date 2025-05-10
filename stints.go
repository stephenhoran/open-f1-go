package openf1go

// Provides information about individual stints.
// A stint refers to a period of continuous driving by a driver during a session.

import (
	"encoding/json"
	"strconv"
)

// Base endpoint for stints-related API calls
const stintsBase = "/stints"

// StintsReponse represents a slice of Stint objects returned by the API
type StintsReponse []Stint

// Stint represents a single stint in a race session
type Stint struct {
	Compound       string `json:"compound"`          // Tyre compound used during the stint
	DriverNumber   int    `json:"driver_number"`     // Unique identifier for the driver
	LapEnd         int    `json:"lap_end"`           // Last lap of the stint
	LapStart       int    `json:"lap_start"`         // First lap of the stint
	MeetingKey     int    `json:"meeting_key"`       // Identifier for the meeting (event)
	SessionKey     int    `json:"session_key"`       // Identifier for the session
	StintNumber    int    `json:"stint_number"`      // Stint number within the session
	TyreAgeAtStart int    `json:"tyre_age_at_start"` // Age of the tyres at the start of the stint
}

// getStintsURL constructs the full URL for the stints API endpoint
func (c *Client) getStintsURL() string {
	return c.baseUrl + stintsBase
}

// GetStints retrieves stints data for a specific stint configuration
func (c *Client) GetStints(stint Stint) (StintsReponse, error) {
	var stintsResponse StintsReponse

	// Build the URL with query parameters based on the provided stint
	url, err := UrlBuilder(c.getStintsURL(), buildArgs(stint))
	if err != nil {
		return nil, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the StintsReponse structure
	if err := json.Unmarshal(resp, &stintsResponse); err != nil {
		return nil, err
	}

	return stintsResponse, nil
}

// GetAllDriversLatestStints retrieves the latest stints for all drivers in the current session
func (c *Client) GetAllDriversLatestStints() (StintsReponse, error) {
	var stintsResponse StintsReponse

	// Build the URL with query parameters for the latest session
	url, err := UrlBuilder(c.getStintsURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the StintsReponse structure
	if err := json.Unmarshal(resp, &stintsResponse); err != nil {
		return nil, err
	}

	return stintsResponse, nil
}

// GetDriverLatestStints retrieves the latest stints for a specific driver in the current session
func (c *Client) GetDriverLatestStints(driver Driver) (StintsReponse, error) {
	var stintsResponse StintsReponse

	// Build the URL with query parameters for the latest session and the specific driver
	url, err := UrlBuilder(c.getStintsURL(), append(c.getLatestSessionArgs(), Arg{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}))
	if err != nil {
		return nil, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the StintsReponse structure
	if err := json.Unmarshal(resp, &stintsResponse); err != nil {
		return nil, err
	}

	return stintsResponse, nil
}
