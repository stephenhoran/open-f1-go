package openf1go

import (
	"encoding/json"
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

func (c *Client) getMeetingsURL() string {
	return c.baseUrl + meetingBase
}

func (c *Client) GetMeetings(meeting Meeting) (MeetingResponse, error) {
	var meetingResponse MeetingResponse

	url, err := UrlBuilder(c.getMeetingsURL(), buildArgs(meeting))
	if err != nil {
		return MeetingResponse{}, err
	}

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
	args := []Arg{{Key: "meeting_key", Value: "latest"}}

	url, err := UrlBuilder(c.getMeetingsURL(), args)
	if err != nil {
		return Meeting{}, err
	}

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
