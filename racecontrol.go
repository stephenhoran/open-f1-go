package openf1go

// Provides information about race control (racing incidents, flags, safety car, ...).

import (
	"encoding/json"
	"strconv"
	"time"
)

// Base path for race control-related API endpoints
const raceControlBase = "/race_control"

// RaceControlResponse represents a list of RaceControl events
type RaceControlResponse []RaceControl

// RaceControl represents a single race control event
// TODO: Handle the nullable fields in the response
type RaceControl struct {
	Category     string          `json:"category"`      // Event category (e.g., CarEvent, DRS, Flag, SafetyCar, Weather)
	Date         time.Time       `json:"date"`          // Timestamp of the event
	DriverNumber int             `json:"driver_number"` // Driver's number (if applicable)
	Flag         string          `json:"flag"`          // Flag type (e.g., Yellow, Red)
	LapNumber    int             `json:"lap_number"`    // Lap number when the event occurred
	MeetingKey   int             `json:"meeting_key"`   // Unique identifier for the meeting
	Message      string          `json:"message"`       // Message or description of the event
	Scope        string          `json:"scope"`         // Scope of the event (e.g., Track, Driver, Sector)
	Sector       json.RawMessage `json:"sector"`        // Sector information (nullable, typically an int)
	SessionKey   int             `json:"session_key"`   // Unique identifier for the session
}

// getRaceControlURL constructs the full URL for race control API requests
func (c *Client) getRaceControlURL() string {
	return c.baseUrl + raceControlBase
}

// GetRaceControl fetches race control data based on the provided RaceControl filters
func (c *Client) GetRaceControl(raceControl RaceControl) (RaceControlResponse, error) {
	var raceControlResponse RaceControlResponse

	// Build the URL with query parameters based on the RaceControl filters
	url, err := UrlBuilder(c.getRaceControlURL(), buildArgs(raceControl))
	if err != nil {
		return nil, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the RaceControlResponse struct
	if err := json.Unmarshal(resp, &raceControlResponse); err != nil {
		return nil, err
	}

	return raceControlResponse, nil
}

// GetAllDriversLatestRaceControl fetches the latest race control data for all drivers
func (c *Client) GetAllDriversLatestRaceControl() (RaceControlResponse, error) {
	var raceControlResponse RaceControlResponse

	// Build the URL for the latest session data
	url, err := UrlBuilder(c.getRaceControlURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the RaceControlResponse struct
	if err := json.Unmarshal(resp, &raceControlResponse); err != nil {
		return nil, err
	}

	return raceControlResponse, nil
}

// GetDriverLatestRaceControl fetches the latest race control data for a specific driver
func (c *Client) GetDriverLatestRaceControl(driver Driver) (RaceControlResponse, error) {
	var raceControlResponse RaceControlResponse

	// Build the URL with the driver number as a query parameter
	url, err := UrlBuilder(c.getRaceControlURL(), append(c.getLatestSessionArgs(), Arg{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}))
	if err != nil {
		return nil, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the RaceControlResponse struct
	if err := json.Unmarshal(resp, &raceControlResponse); err != nil {
		return nil, err
	}

	return raceControlResponse, nil
}
