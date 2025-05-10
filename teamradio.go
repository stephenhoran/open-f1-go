package openf1go

// Provides a collection of radio exchanges between Formula 1 drivers and their respective teams during sessions.
// Please note that only a limited selection of communications are included, not the complete record of radio interactions.

import (
	"encoding/json"
	"strconv"
)

const teamRadioBase = "/team_radio"

// TeamRadioResponse represents a collection of team radio exchanges.
type TeamRadioResponse []TeamRadio

// TeamRadio represents a single team radio exchange with relevant metadata.
type TeamRadio struct {
	Date         string `json:"date"`          // Date of the radio exchange.
	DriverNumber int    `json:"driver_number"` // Driver's unique number.
	MeetingKey   int    `json:"meeting_key"`   // Identifier for the meeting/session.
	RecordingURL string `json:"recording_url"` // URL to the recording of the radio exchange.
	SessionKey   int    `json:"session_key"`   // Identifier for the session.
}

// getTeamRadioURL constructs the base URL for team radio API endpoints.
func (c *Client) getTeamRadioURL() string {
	return c.baseUrl + teamRadioBase
}

// GetTeamRadio fetches team radio data for a specific TeamRadio object.
func (c *Client) GetTeamRadio(teamRadio TeamRadio) (TeamRadioResponse, error) {
	var teamRadioResponse TeamRadioResponse

	// Build the URL with the necessary arguments.
	url, err := UrlBuilder(c.getTeamRadioURL(), buildArgs(teamRadio))
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request.
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the TeamRadioResponse struct.
	if err := json.Unmarshal(resp, &teamRadioResponse); err != nil {
		return nil, err
	}

	return teamRadioResponse, nil
}

// GetAllDriversLatestTeamRadio fetches the latest team radio data for all drivers.
func (c *Client) GetAllDriversLatestTeamRadio() (TeamRadioResponse, error) {
	var teamRadioResponse TeamRadioResponse

	// Build the URL with the latest session arguments.
	url, err := UrlBuilder(c.getTeamRadioURL(), c.getLatestSessionArgs())
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request.
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the TeamRadioResponse struct.
	if err := json.Unmarshal(resp, &teamRadioResponse); err != nil {
		return nil, err
	}

	return teamRadioResponse, nil
}

// GetDriverLatestTeamRadio fetches the latest team radio data for a specific driver.
func (c *Client) GetDriverLatestTeamRadio(driver Driver) (TeamRadioResponse, error) {
	var teamRadioResponse TeamRadioResponse

	// Build the URL with the latest session arguments and the driver's number.
	url, err := UrlBuilder(c.getTeamRadioURL(), append(c.getLatestSessionArgs(), Arg{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}))
	if err != nil {
		return nil, err
	}

	// Make the HTTP GET request.
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the TeamRadioResponse struct.
	if err := json.Unmarshal(resp, &teamRadioResponse); err != nil {
		return nil, err
	}

	return teamRadioResponse, nil
}
