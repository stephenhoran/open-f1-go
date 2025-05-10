package openf1go

import (
	"encoding/json"
)

const driversBase = "/drivers"

type DriversResponse []Driver

type Driver struct {
	BroadcastName string `json:"broadcast_name"`
	CountryCode   string `json:"country_code"`
	DriverNumber  int    `json:"driver_number"`
	FirstName     string `json:"first_name"`
	FullName      string `json:"full_name"`
	HeadshotURL   string `json:"headshot_url"`
	LastName      string `json:"last_name"`
	MeetingKey    int    `json:"meeting_key"`
	NameAcronym   string `json:"name_acronym"`
	SessionKey    int    `json:"session_key"`
	TeamColour    string `json:"team_colour"`
	TeamName      string `json:"team_name"`
}

func (c *Client) getDriversURL() string {
	return c.baseUrl + driversBase
}

func (c *Client) GetDrivers(driver Driver) (DriversResponse, error) {
	var driversResponse DriversResponse

	url, err := UrlBuilder(c.getDriversURL(), buildArgs(driver))
	if err != nil {
		return DriversResponse{}, err
	}

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &driversResponse); err != nil {
		return nil, err
	}

	return driversResponse, nil
}

func (c *Client) GetLatestDrivers() (DriversResponse, error) {
	var driversResponse DriversResponse

	args := []Arg{{Key: "meeting_key", Value: "latest"}, {Key: "session_key", Value: "latest"}}

	url, err := UrlBuilder(c.getDriversURL(), args)
	if err != nil {
		return DriversResponse{}, err
	}

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return DriversResponse{}, err
	}

	if err := json.Unmarshal(resp, &driversResponse); err != nil {
		return DriversResponse{}, err
	}

	return driversResponse, nil
}
