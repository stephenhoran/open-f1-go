package openf1go

// Some data about each car, at a sample rate of about 3.7 Hz.

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

// carDataBase defines the base endpoint for car data API.
const carDataBase = "/car_data"

// CarDataResponse represents a slice of CarData objects returned by the API.
type CarDataResponse = []CarData

// CarData represents the structure of car telemetry data.
type CarData struct {
	Brake        int       `json:"brake"`         // Brake pressure percentage (0-100).
	Date         time.Time `json:"date"`          // Timestamp of the data.
	DriverNumber int       `json:"driver_number"` // Unique identifier for the driver.
	Drs          int       `json:"drs"`           // DRS status: 0/1 = off, 8 = eligible, 10/12/14 = on.
	MeetingKey   int       `json:"meeting_key"`   // Identifier for the meeting/session.
	NGear        int       `json:"n_gear"`        // Current gear of the car.
	Rpm          int       `json:"rpm"`           // Engine revolutions per minute.
	SessionKey   int       `json:"session_key"`   // Identifier for the session.
	Speed        int       `json:"speed"`         // Speed of the car in km/h.
	Throttle     int       `json:"throttle"`      // Throttle pressure percentage (0-100).
}

// getCarDataURL constructs the full URL for the car data API endpoint.
func (c *Client) getCarDataURL() string {
	return c.baseUrl + carDataBase
}

// GetCarData fetches car telemetry data based on the provided CarData filter.
// Returns a CarDataResponse or an error if the request fails.
func (c *Client) GetCarData(CarData CarData) (CarDataResponse, error) {
	var carDataResponse CarDataResponse

	// Build the URL with query parameters based on the CarData filter.
	url, err := UrlBuilder(c.getCarDataURL(), buildArgs(CarData))
	if err != nil {
		return CarDataResponse{}, err
	}

	// Make the HTTP GET request to fetch car data.
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the CarDataResponse structure.
	if err := json.Unmarshal(resp, &carDataResponse); err != nil {
		return nil, err
	}

	return carDataResponse, nil
}

// GetLatestCarDataByDriver fetches the latest car telemetry data for a specific driver.
// Returns a CarDataResponse or an error if the request fails.
func (c *Client) GetLatestCarDataByDriver(driver Driver) (CarDataResponse, error) {
	var carDataResponse CarDataResponse

	// Validate that the driver has a valid driver number.
	if driver.DriverNumber == 0 {
		return CarDataResponse{}, errors.New("provided driver missing driver number or empty driver")
	}

	// Build query arguments for the driver and the latest session.
	args := []Arg{{Key: "driver_number", Value: strconv.Itoa(driver.DriverNumber)}}
	args = append(args, c.getLatestSessionArgs()...)

	// Build the URL with the query arguments.
	url, err := UrlBuilder(c.getCarDataURL(), args)
	if err != nil {
		return CarDataResponse{}, err
	}

	// Make the HTTP GET request to fetch the latest car data.
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return CarDataResponse{}, err
	}

	// Parse the JSON response into the CarDataResponse structure.
	if err := json.Unmarshal(resp, &carDataResponse); err != nil {
		return CarDataResponse{}, err
	}

	return carDataResponse, nil
}
