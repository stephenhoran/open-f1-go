package openf1go

import (
	"encoding/json"
	"time"
)

const sessionsBase = "/sessions"

type SessionResponse []Session

type SessionParams struct {
}

type Session struct {
	CircuitKey       int       `json:"circuit_key"`
	CircuitShortName string    `json:"circuit_short_name"`
	CountryCode      string    `json:"country_code"`
	CountryKey       int       `json:"country_key"`
	CountryName      string    `json:"country_name"`
	DateEnd          time.Time `json:"date_end"`
	DateStart        time.Time `json:"date_start"`
	GmtOffset        string    `json:"gmt_offset"`
	Location         string    `json:"location"`
	MeetingKey       int       `json:"meeting_key"`
	SessionKey       int       `json:"session_key"`
	SessionName      string    `json:"session_name"`
	SessionType      string    `json:"session_type"`
	Year             int       `json:"year"`
}

func (c *Client) getSessionsURL() string {
	return c.baseUrl + sessionsBase
}

func (c *Client) GetSessions(session Session) (SessionResponse, error) {
	var sessionResponse SessionResponse

	url, err := UrlBuilder(c.getSessionsURL(), buildArgs(session))
	if err != nil {
		return SessionResponse{}, nil
	}

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &sessionResponse); err != nil {
		return nil, err
	}

	return sessionResponse, nil

}

func (c *Client) GetLatestSessions() (Session, error) {
	var sessionResponse SessionResponse

	args := []Arg{{Key: "meeting_key", Value: "latest"}, {Key: "session_key", Value: "latest"}}

	url, err := UrlBuilder(c.getSessionsURL(), args)
	if err != nil {
		return Session{}, nil
	}

	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Session{}, err
	}

	if err := json.Unmarshal(resp, &sessionResponse); err != nil {
		return Session{}, err
	}

	return sessionResponse[0], nil

}
