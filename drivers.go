package openf1go

import (
	"encoding/json"
	"errors"
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

func (c *Client) GetDriver(driver Driver) (Driver, error) {
	var driversResponse DriversResponse

	if driver.FirstName == "" && driver.LastName == "" && driver.FullName == "" && driver.NameAcronym == "" && driver.DriverNumber == 0 {
		return Driver{}, errors.New("search fields for a single driver not met, ensure at least one drive identifier is valid")
	}

	args := buildArgs(driver)
	if driver.MeetingKey == 0 && driver.SessionKey == 0 {
		args = append(args, c.getLatestSessionArgs()...)
	}

	url, err := UrlBuilder(c.getDriversURL(), args)
	if err != nil {
		return Driver{}, err
	}

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Driver{}, err
	}

	if err := json.Unmarshal(resp, &driversResponse); err != nil {
		return Driver{}, err
	}

	if len(driversResponse) == 0 {
		return Driver{}, errors.New("driver not found or too many drivers returned from search")
	}

	return driversResponse[0], nil
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
