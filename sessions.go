package openf1go

// Provides information about sessions.
// A session refers to a distinct period of track activity during a Grand Prix or testing weekend (practice, qualifying, sprint, race, ...).

import (
	"encoding/json"
	"time"
)

// Base endpoint for sessions API
const sessionsBase = "/sessions"

// SessionResponse represents a list of sessions returned by the API
type SessionResponse []Session

// Session represents the structure of a session object returned by the API
type Session struct {
	CircuitKey       int       `json:"circuit_key"`        // Unique identifier for the circuit
	CircuitShortName string    `json:"circuit_short_name"` // Short name of the circuit
	CountryCode      string    `json:"country_code"`       // ISO country code
	CountryKey       int       `json:"country_key"`        // Unique identifier for the country
	CountryName      string    `json:"country_name"`       // Full name of the country
	DateEnd          time.Time `json:"date_end"`           // End date of the session
	DateStart        time.Time `json:"date_start"`         // Start date of the session
	GmtOffset        string    `json:"gmt_offset"`         // GMT offset for the session location
	Location         string    `json:"location"`           // Location of the session
	MeetingKey       int       `json:"meeting_key"`        // Unique identifier for the meeting
	SessionKey       int       `json:"session_key"`        // Unique identifier for the session
	SessionName      string    `json:"session_name"`       // Name of the session
	SessionType      string    `json:"session_type"`       // Type of the session (e.g., practice, qualifying, race)
	Year             int       `json:"year"`               // Year of the session
}

// getSessionsURL constructs the full URL for the sessions API
func (c *Client) getSessionsURL() string {
	return c.baseUrl + sessionsBase
}

// GetSessions fetches sessions based on the provided session filter
func (c *Client) GetSessions(session Session) (SessionResponse, error) {
	var sessionResponse SessionResponse

	// Build the URL with query parameters
	url, err := UrlBuilder(c.getSessionsURL(), buildArgs(session))
	if err != nil {
		return SessionResponse{}, nil // Return empty response if URL building fails
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err // Return error if the request fails
	}

	// Parse the JSON response into the SessionResponse struct
	if err := json.Unmarshal(resp, &sessionResponse); err != nil {
		return nil, err // Return error if JSON unmarshalling fails
	}

	return sessionResponse, nil // Return the parsed session response
}

// GetLatestSessions fetches the most recent session
func (c *Client) GetLatestSessions() (Session, error) {
	var sessionResponse SessionResponse

	// Build the URL for the latest session
	url, err := UrlBuilder(c.getSessionsURL(), c.getLatestSessionArgs())
	if err != nil {
		return Session{}, nil // Return empty session if URL building fails
	}

	// Make the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Session{}, err // Return error if the request fails
	}

	// Parse the JSON response into the SessionResponse struct
	if err := json.Unmarshal(resp, &sessionResponse); err != nil {
		return Session{}, err // Return error if JSON unmarshalling fails
	}

	// Return the first session from the response as the latest session
	return sessionResponse[0], nil
}
