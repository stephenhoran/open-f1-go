package openf1go

// Provides information about drivers for each session.

import (
	"encoding/json"
	"errors"
)

// Base endpoint for drivers-related API calls
const driversBase = "/drivers"

// DriversResponse represents a list of Driver objects
type DriversResponse []Driver

// Driver represents the structure of a driver object returned by the API
type Driver struct {
	BroadcastName string `json:"broadcast_name"` // Name used in broadcasts
	CountryCode   string `json:"country_code"`   // Country code of the driver
	DriverNumber  int    `json:"driver_number"`  // Driver's number
	FirstName     string `json:"first_name"`     // Driver's first name
	FullName      string `json:"full_name"`      // Driver's full name
	HeadshotURL   string `json:"headshot_url"`   // URL to the driver's headshot image
	LastName      string `json:"last_name"`      // Driver's last name
	MeetingKey    int    `json:"meeting_key"`    // Key identifying the meeting
	NameAcronym   string `json:"name_acronym"`   // Acronym of the driver's name
	SessionKey    int    `json:"session_key"`    // Key identifying the session
	TeamColour    string `json:"team_colour"`    // Team's color
	TeamName      string `json:"team_name"`      // Name of the team
}

// getDriversURL constructs the full URL for the drivers endpoint
func (c *Client) getDriversURL() string {
	return c.baseUrl + driversBase
}

// GetDrivers fetches a list of drivers based on the provided driver filters
func (c *Client) GetDrivers(driver Driver) (DriversResponse, error) {
	var driversResponse DriversResponse

	// Build the URL with query parameters
	url, err := UrlBuilder(c.getDriversURL(), buildArgs(driver))
	if err != nil {
		return DriversResponse{}, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the DriversResponse struct
	if err := json.Unmarshal(resp, &driversResponse); err != nil {
		return nil, err
	}

	return driversResponse, nil
}

// GetDriver fetches a single driver based on the provided driver filters
func (c *Client) GetDriver(driver Driver) (Driver, error) {
	var driversResponse DriversResponse

	// Validate that at least one search field is provided
	if driver.FirstName == "" && driver.LastName == "" && driver.FullName == "" && driver.NameAcronym == "" && driver.DriverNumber == 0 {
		return Driver{}, errors.New("search fields for a single driver not met, ensure at least one drive identifier is valid")
	}

	// Build query arguments and include the latest session if keys are not provided
	args := buildArgs(driver)
	if driver.MeetingKey == 0 && driver.SessionKey == 0 {
		args = append(args, c.getLatestSessionArgs()...)
	}

	// Build the URL with query parameters
	url, err := UrlBuilder(c.getDriversURL(), args)
	if err != nil {
		return Driver{}, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Driver{}, err
	}

	// Parse the JSON response into the DriversResponse struct
	if err := json.Unmarshal(resp, &driversResponse); err != nil {
		return Driver{}, err
	}

	// Ensure exactly one driver is returned
	if len(driversResponse) != 1 {
		return Driver{}, errors.New("driver not found or too many drivers returned from search")
	}

	return driversResponse[0], nil
}

// GetLatestDrivers fetches the latest drivers for the most recent meeting and session
func (c *Client) GetLatestDrivers() (DriversResponse, error) {
	var driversResponse DriversResponse

	// Define arguments for the latest meeting and session
	args := []Arg{{Key: "meeting_key", Value: "latest"}, {Key: "session_key", Value: "latest"}}

	// Build the URL with query parameters
	url, err := UrlBuilder(c.getDriversURL(), args)
	if err != nil {
		return DriversResponse{}, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return DriversResponse{}, err
	}

	// Parse the JSON response into the DriversResponse struct
	if err := json.Unmarshal(resp, &driversResponse); err != nil {
		return DriversResponse{}, err
	}

	return driversResponse, nil
}
