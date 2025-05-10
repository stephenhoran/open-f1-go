package openf1go

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const carDataBase = "/car_data"

type CarDataResponse = []CarData

type CarData struct {
	Brake        int       `json:"brake"`
	Date         time.Time `json:"date"`
	DriverNumber int       `json:"driver_number"`
	Drs          int       `json:"drs"` // 0 == off, 1 == off, 8 == eligible, 10 == on, 12 == on, 14 == on
	MeetingKey   int       `json:"meeting_key"`
	NGear        int       `json:"n_gear"`
	Rpm          int       `json:"rpm"`
	SessionKey   int       `json:"session_key"`
	Speed        int       `json:"speed"`
	Throttle     int       `json:"throttle"`
}

func (c *Client) getCarDataURL() string {
	return c.baseUrl + carDataBase
}

func (c *Client) GetCarData(CarData CarData) (CarDataResponse, error) {
	var carDataResponse CarDataResponse

	url, err := UrlBuilder(c.getCarDataURL(), buildArgs(CarData))
	if err != nil {
		return CarDataResponse{}, err
	}

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &carDataResponse); err != nil {
		return nil, err
	}

	return carDataResponse, nil
}

func (c *Client) GetLatestCarDataByDriver(Driver Driver) (CarDataResponse, error) {
	var carDataResponse CarDataResponse

	if Driver.DriverNumber == 0 {
		return CarDataResponse{}, errors.New("provided driver missing driver number or empty driver")
	}

	args := []Arg{{Key: "driver_number", Value: strconv.Itoa(Driver.DriverNumber)}}
	args = append(args, c.getLatestSessionArgs()...)

	url, err := UrlBuilder(c.getCarDataURL(), args)
	if err != nil {
		return CarDataResponse{}, err
	}

	fmt.Println(url)
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return CarDataResponse{}, err
	}

	if err := json.Unmarshal(resp, &carDataResponse); err != nil {
		return CarDataResponse{}, err
	}

	return carDataResponse, nil
}
