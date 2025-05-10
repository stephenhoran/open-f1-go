package openf1go

// Fetches real-time interval data between drivers and their gap to the race leader.
// Available during races only, with updates approximately every 4 seconds.

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// Base endpoint for intervals API
const intervalsBase = "/intervals"

// IntervalsResponse represents a slice of Interval objects
type IntervalsResponse []Interval

// Interval represents the structure of an interval record
type Interval struct {
	Date         time.Time       `json:"date"`          // Date of the interval
	DriverNumber int             `json:"driver_number"` // Driver's unique number
	GapToLeader  json.RawMessage `json:"gap_to_leader"` // Gap to the leader in raw JSON format
	Interval     json.RawMessage `json:"interval"`      // Interval to the next driver in raw JSON format
	MeetingKey   int             `json:"meeting_key"`   // Unique key for the meeting
	SessionKey   int             `json:"session_key"`   // Unique key for the session
}

// getIntervalsURL constructs the full URL for the intervals API
func (c *Client) getIntervalsURL() string {
	return c.baseUrl + intervalsBase
}

// GetIntervals fetches intervals data based on the provided Interval filter
func (c *Client) GetIntervals(interval Interval) (IntervalsResponse, error) {
	var intervalsResponse IntervalsResponse

	// Build the URL with query parameters
	url, err := UrlBuilder(c.getIntervalsURL(), buildArgs(interval))
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into IntervalsResponse
	if err := json.Unmarshal(resp, &intervalsResponse); err != nil {
		return nil, err
	}

	return intervalsResponse, nil
}

// GetAllDriversCurrentIntervals fetches the current intervals for all drivers
func (c *Client) GetAllDriversCurrentIntervals() (IntervalsResponse, error) {
	var intervalsResponse IntervalsResponse

	// Build the URL with the latest session arguments
	url, err := UrlBuilder(c.getIntervalsURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into IntervalsResponse
	if err := json.Unmarshal(resp, &intervalsResponse); err != nil {
		return nil, err
	}

	return intervalsResponse, nil
}

// GetDriverCurrentIntervals fetches the current intervals for a specific driver
func (c *Client) GetDriverCurrentIntervals(driver Driver) (IntervalsResponse, error) {
	var intervalsResponse IntervalsResponse

	// Validate that the driver has a valid driver number
	if driver.DriverNumber == 0 {
		return IntervalsResponse{}, errors.New("provided driver missing driver number or empty driver")
	}

	// Build query arguments with the driver's number and latest session arguments
	args := []Arg{{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}}
	args = append(args, c.getLatestSessionArgs()...)

	// Build the URL with the query arguments
	url, err := UrlBuilder(c.getIntervalsURL(), args)
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into IntervalsResponse
	if err := json.Unmarshal(resp, &intervalsResponse); err != nil {
		return nil, err
	}

	return intervalsResponse, nil
}
