package openf1go

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const meetingBase = "/meetings"

type MeetingResponse []Meeting

type MeetingParams struct {
	Year        string
	CountryName string
}

type Meeting struct {
	CircuitKey          int       `json:"circuit_key"`
	CircuitShortName    string    `json:"circuit_short_name"`
	CountryCode         string    `json:"country_code"`
	CountryKey          int       `json:"country_key"`
	CountryName         string    `json:"country_name"`
	DateStart           time.Time `json:"date_start"`
	GMTOffset           string    `json:"gmt_offset"`
	Location            string    `json:"location"`
	MeetingKey          int       `json:"meeting_key"`
	MeetingName         string    `json:"meeting_name"`
	MeetingOfficialName string    `json:"meeting_official_name"`
	Year                int       `json:"year"`
}

func (c *Client) getMettingsURL() string {
	return c.baseUrl + meetingBase
}

func (c *Client) GetMeetings(params MeetingParams) (MeetingResponse, error) {
	var meetingResponse MeetingResponse
	args := []Arg{}

	if params.Year != "" {
		args = append(args, Arg{
			Key:   "year",
			Value: params.Year,
		})
	}

	if params.CountryName != "" {
		args = append(args, Arg{
			Key:   "country_name",
			Value: params.CountryName,
		})
	}

	urlArgs := ArgBuilder(args)
	url := c.getMettingsURL() + urlArgs

	fmt.Println(url)
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &meetingResponse); err != nil {
		return nil, err
	}

	return meetingResponse, nil
}

func (c *Client) GetLatestMeeting() (Meeting, error) {
	var meetingResponse MeetingResponse
	var meeting Meeting
	args := []Arg{}

	t := time.Now()

	args = append(args, Arg{Key: "year", Value: strconv.Itoa(t.Year())})
	urlArgs := ArgBuilder(args)

	url := c.getMettingsURL() + urlArgs

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Meeting{}, err
	}

	if err := json.Unmarshal(resp, &meetingResponse); err != nil {
		return Meeting{}, err
	}

	for _, m := range meetingResponse {
		if m.DateStart.After(meeting.DateStart) {
			meeting = m
		}
	}

	return meeting, nil
}
