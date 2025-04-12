package openf1go

import "fmt"

const meetingBase = "/meeting"

type MeetingResponse []Meeting

type MeetingParams struct {
	Year        string
	CountryName string
}

type Meeting struct {
	CircuitKey          int    `json:"circuit_key"`
	CircuitShortName    string `json:"circuit_short_name"`
	CountryCode         string `json:"country_code"`
	CountryKey          int    `json:"country_key"`
	CountryName         string `json:"country_name"`
	DateStart           string `json:"date_start"`
	GMTOffset           string `json:"gmt_offset"`
	Location            string `json:"location"`
	MeetingKey          string `json:"meeting_key"`
	MeetingName         string `json:"meeting_name"`
	MeetingOfficialName string `json:"meeting_official_name"`
	Year                int    `json:"year"`
}

func (c *Client) GetMeetings(params MeetingParams) (MeetingResponse, error) {
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

	fmt.Println(c.baseUrl + meetingBase + urlArgs)

	return nil, nil
}
