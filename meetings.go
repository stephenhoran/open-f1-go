package openf1go

// Provides information about meetings.
// A meeting refers to a Grand Prix or testing weekend and usually includes multiple sessions (practice, qualifying, race, ...).

import (
	"encoding/json"
	"time"
)

const meetingBase = "/meetings" // Base path for meeting-related API endpoints

// MeetingResponse represents a slice of Meeting objects returned by the API
type MeetingResponse []Meeting

// Meeting represents the structure of a single meeting object
type Meeting struct {
	CircuitKey          int       `json:"circuit_key"`           // Unique identifier for the circuit
	CircuitShortName    string    `json:"circuit_short_name"`    // Short name of the circuit
	CountryCode         string    `json:"country_code"`          // ISO country code
	CountryKey          int       `json:"country_key"`           // Unique identifier for the country
	CountryName         string    `json:"country_name"`          // Full name of the country
	DateStart           time.Time `json:"date_start"`            // Start date of the meeting
	GMTOffset           string    `json:"gmt_offset"`            // GMT offset for the meeting location
	Location            string    `json:"location"`              // Location of the meeting
	MeetingKey          int       `json:"meeting_key"`           // Unique identifier for the meeting
	MeetingName         string    `json:"meeting_name"`          // Name of the meeting
	MeetingOfficialName string    `json:"meeting_official_name"` // Official name of the meeting
	Year                int       `json:"year"`                  // Year of the meeting
}

// getMeetingsURL constructs the full URL for the meetings API endpoint
func (c *Client) getMeetingsURL() string {
	return c.baseUrl + meetingBase
}

// GetMeetings fetches a list of meetings based on the provided meeting filter
func (c *Client) GetMeetings(meeting Meeting) (MeetingResponse, error) {
	var meetingResponse MeetingResponse

	// Build the URL with query parameters based on the meeting filter
	url, err := UrlBuilder(c.getMeetingsURL(), buildArgs(meeting))
	if err != nil {
		return MeetingResponse{}, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON response into the MeetingResponse struct
	if err := json.Unmarshal(resp, &meetingResponse); err != nil {
		return nil, err
	}

	return meetingResponse, nil
}

// GetLatestMeeting fetches the most recent meeting based on the start date
func (c *Client) GetLatestMeeting() (Meeting, error) {
	var meetingResponse MeetingResponse // Holds the list of meetings from the API
	var meeting Meeting                 // Holds the latest meeting

	// Define the query parameter to fetch the latest meeting
	args := []Arg{{Key: "meeting_key", Value: "latest"}}

	// Build the URL for the API request
	url, err := UrlBuilder(c.getMeetingsURL(), args)
	if err != nil {
		return Meeting{}, err
	}

	// Perform the HTTP GET request
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Meeting{}, err
	}

	// Unmarshal the JSON response into the MeetingResponse struct
	if err := json.Unmarshal(resp, &meetingResponse); err != nil {
		return Meeting{}, err
	}

	// Iterate through the meetings to find the one with the latest start date
	for _, m := range meetingResponse {
		if m.DateStart.After(meeting.DateStart) {
			meeting = m
		}
	}

	return meeting, nil
}
