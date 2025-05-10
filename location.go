package openf1go

// The approximate location of the cars on the circuit, at a sample rate of about 3.7 Hz.
// Useful for gauging their progress along the track, but lacks details about lateral placement — i.e. whether the car is on the left or right side of the track. The origin point (0, 0, 0) appears to be arbitrary and not tied to any specific location on the track.

import (
	"encoding/json"
	"strconv"
	"time"
)

// Base path for location-related API endpoints
const locationBase = "/location"

// LocationResponse represents a slice of Location objects
type LocationResponse []Location

// Location represents the details of a driver's location
type Location struct {
	Date         time.Time `json:"date"`          // Timestamp of the location data
	DriverNumber int       `json:"driver_number"` // Unique identifier for the driver
	MeetingKey   int       `json:"meeting_key"`   // Identifier for the meeting
	SessionKey   int       `json:"session_key"`   // Identifier for the session
	X            int       `json:"x"`             // X-coordinate of the location
	Y            int       `json:"y"`             // Y-coordinate of the location
	Z            int       `json:"z"`             // Z-coordinate of the location
}

// getLocationURL constructs the full URL for location-related API calls
func (c *Client) getLocationURL() string {
	return c.baseUrl + locationBase
}

// GetLocations fetches location data based on the provided Location object
func (c *Client) GetLocations(location Location) (LocationResponse, error) {
	var locationResponse LocationResponse

	// Build the URL with query parameters
	url, err := UrlBuilder(c.getLocationURL(), buildArgs(location))
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the LocationResponse struct
	if err := json.Unmarshal(resp, &locationResponse); err != nil {
		return nil, err
	}

	return locationResponse, nil
}

// GetAllDriversLatestLocations fetches the latest location data for all drivers
func (c *Client) GetAllDriversLatestLocations() (LocationResponse, error) {
	var locationResponse LocationResponse

	// Build the URL with session-specific arguments
	url, err := UrlBuilder(c.getLocationURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the LocationResponse struct
	if err := json.Unmarshal(resp, &locationResponse); err != nil {
		return nil, err
	}

	return locationResponse, nil
}

// GetDriverLatestLocation fetches the latest location data for a specific driver
func (c *Client) GetDriverLatestLocation(driver Driver) (LocationResponse, error) {
	var locationResponse LocationResponse

	// Validate that the driver number is provided
	if driver.DriverNumber == 0 {
		return nil, ErrDriverNumberMissing
	}

	// Build the query arguments with the driver number and session-specific arguments
	args := []Arg{{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}}
	args = append(args, c.getLatestSessionArgs()...)

	// Build the URL with the query arguments
	url, err := UrlBuilder(c.getLocationURL(), args)
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the LocationResponse struct
	if err := json.Unmarshal(resp, &locationResponse); err != nil {
		return nil, err
	}

	return locationResponse, nil
}
