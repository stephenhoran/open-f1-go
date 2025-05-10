package openf1go

// Provides detailed information about individual laps.

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// Base endpoint for laps-related API calls
const lapsBase = "/laps"

// LapsResponse is a type alias for a slice of Lap structs
type LapsResponse = []Lap

// Lap represents the structure of a single lap's data
type Lap struct {
	DateStart       time.Time `json:"date_start"`        // Start time of the lap
	DriverNumber    int       `json:"driver_number"`     // Driver's unique number
	DurationSector1 float64   `json:"duration_sector_1"` // Duration of sector 1 in seconds
	DurationSector2 float64   `json:"duration_sector_2"` // Duration of sector 2 in seconds
	DurationSector3 float64   `json:"duration_sector_3"` // Duration of sector 3 in seconds
	I1Speed         int       `json:"i1_speed"`          // Speed at intermediate 1
	I2Speed         int       `json:"i2_speed"`          // Speed at intermediate 2
	IsPitOutLap     bool      `json:"is_pit_out_lap"`    // Indicates if this is a pit out lap
	LapDuration     float64   `json:"lap_duration"`      // Total lap duration in seconds
	LapNumber       int       `json:"lap_number"`        // Lap number in the session
	MeetingKey      int       `json:"meeting_key"`       // Unique identifier for the meeting
	SegmentsSector1 []int     `json:"segments_sector_1"` // Segment times for sector 1
	SegmentsSector2 []int     `json:"segments_sector_2"` // Segment times for sector 2
	SegmentsSector3 []int     `json:"segments_sector_3"` // Segment times for sector 3
	SessionKey      int       `json:"session_key"`       // Unique identifier for the session
	StSpeed         int       `json:"st_speed"`          // Speed at the start/finish line
}

// getLapURL constructs the full URL for the laps endpoint
func (c *Client) getLapURL() string {
	return c.baseUrl + lapsBase
}

// GetLaps retrieves laps data based on the provided Lap struct as a filter
func (c *Client) GetLaps(lap Lap) (LapsResponse, error) {
	var lapsResponse LapsResponse

	// Build the URL with query parameters based on the provided lap
	url, err := UrlBuilder(c.getLapURL(), buildArgs(lap))
	if err != nil {
		return LapsResponse{}, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the LapsResponse struct
	if err := json.Unmarshal(resp, &lapsResponse); err != nil {
		return nil, err
	}

	return lapsResponse, nil
}

// GetLatestLapsByDriver retrieves the latest laps for a specific driver
func (c *Client) GetLatestLapsByDriver(driver Driver) (LapsResponse, error) {
	var lapsResponse LapsResponse

	// Validate that the driver has a valid driver number
	if driver.DriverNumber == 0 {
		return LapsResponse{}, errors.New("provided driver missing driver number or empty driver")
	}

	// Build query parameters for the driver and session
	args := []Arg{{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}}
	args = append(args, c.getLatestSessionArgs()...)

	// Build the URL with the query parameters
	url, err := UrlBuilder(c.getLapURL(), args)
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the LapsResponse struct
	if err := json.Unmarshal(resp, &lapsResponse); err != nil {
		return nil, err
	}

	return lapsResponse, nil
}

// GetLatestLaps retrieves the latest laps for the current session
func (c *Client) GetLatestLaps() (LapsResponse, error) {
	var lapsResponse LapsResponse

	// Build the URL with query parameters for the latest session
	url, err := UrlBuilder(c.getLapURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the LapsResponse struct
	if err := json.Unmarshal(resp, &lapsResponse); err != nil {
		return nil, err
	}

	return lapsResponse, nil
}
