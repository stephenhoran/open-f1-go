package openf1go

// The weather over the track, updated every minute.

import (
	"encoding/json"
	"time"
)

// Base endpoint for weather-related API requests
const weatherBase = "/weather"

// WeatherResponse represents a slice of Weather objects returned by the API
type WeatherResponse []Weather

// Weather struct defines the structure of weather data returned by the API
type Weather struct {
	AirTemperature   float64   `json:"air_temperature"`   // Air temperature in degrees Celsius
	Date             time.Time `json:"date"`              // Date and time of the weather data
	Humidity         float64   `json:"humidity"`          // Humidity percentage
	MeetingKey       int       `json:"meeting_key"`       // Identifier for the meeting
	Pressure         float64   `json:"pressure"`          // Atmospheric pressure in hPa
	Rainfall         int       `json:"rainfall"`          // Rainfall in millimeters
	SessionKey       int       `json:"session_key"`       // Identifier for the session
	TrackTemperature float64   `json:"track_temperature"` // Track temperature in degrees Celsius
	WindDirection    int       `json:"wind_direction"`    // Wind direction in degrees
	WindSpeed        float64   `json:"wind_speed"`        // Wind speed in meters per second
}

// getWeatherURL constructs the full URL for weather-related API requests
func (c *Client) getWeatherURL() string {
	return c.baseUrl + weatherBase
}

// GetWeather fetches weather data based on the provided Weather struct
func (c *Client) GetWeather(weather Weather) (WeatherResponse, error) {
	var weatherResponse WeatherResponse

	// Build the URL with query parameters based on the provided Weather struct
	url, err := UrlBuilder(c.getWeatherURL(), buildArgs(weather))
	if err != nil {
		return nil, err
	}

	// Make an HTTP GET request to fetch the weather data
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response into the WeatherResponse struct
	if err := json.Unmarshal(resp, &weatherResponse); err != nil {
		return nil, err
	}

	return weatherResponse, nil
}

// GetLatestWeather fetches the most recent weather data for the latest session
func (c *Client) GetLatestWeather() (Weather, error) {
	var weatherResponse WeatherResponse

	// Build the URL with query parameters for the latest session
	url, err := UrlBuilder(c.getWeatherURL(), c.getLatestSessionArgs())
	if err != nil {
		return Weather{}, err
	}

	// Make an HTTP GET request to fetch the latest weather data
	resp, err := GetHTTPRequest(url)
	if err != nil {
		return Weather{}, err
	}

	// Parse the JSON response into the WeatherResponse struct
	if err := json.Unmarshal(resp, &weatherResponse); err != nil {
		return Weather{}, err
	}

	// Return an empty Weather struct if no data is available
	if len(weatherResponse) == 0 {
		return Weather{}, nil
	}

	// Return the first (most recent) weather record
	return weatherResponse[0], nil
}
