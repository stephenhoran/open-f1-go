# open-f1-go

`open-f1-go` is a Go client library for interacting with the Open F1 API. It provides methods to fetch data about drivers, meetings, sessions, laps, intervals, and car telemetry.

**Notice:** This project is currently in **beta**. Some features that will be added in the future include:
- Better error handling after gaining more experience with the API.
- Helpers for working with time.

## Installation

To use this library, add it to your Go project by running:

```sh
go get github.com/stephenhoran/open-f1-go
```

## Usage

### Initialize the Client

To start using the library, create a new client instance:

```go
import openf1go "github.com/stephenhoran/open-f1-go"

func main() {
	client := openf1go.New()
}
```

## API Endpoints

### Drivers
Fetch information about drivers participating in a session.

#### Example: Fetch Latest Drivers
```go
drivers, err := client.GetLatestDrivers()
if err != nil {
	fmt.Println("Error fetching drivers:", err)
	return
}

for _, driver := range drivers {
	fmt.Printf("Driver: %s (%d)\n", driver.FullName, driver.DriverNumber)
}
```

#### Example: Fetch a Specific Driver
```go
driver, err := client.GetDriver(openf1go.Driver{DriverNumber: 44})
if err != nil {
	fmt.Println("Error fetching driver:", err)
	return
}

fmt.Printf("Driver: %s (%d)\n", driver.FullName, driver.DriverNumber)
```

---

### Meetings
Fetch information about Grand Prix or testing weekends.

#### Example: Fetch Latest Meeting
```go
meeting, err := client.GetLatestMeeting()
if err != nil {
	fmt.Println("Error fetching latest meeting:", err)
	return
}

fmt.Printf("Meeting: %s (%s)\n", meeting.MeetingName, meeting.CountryName)
}
```

---

### Sessions
Fetch information about specific sessions (practice, qualifying, race, etc.).

#### Example: Fetch Sessions for a Meeting
```go
sessions, err := client.GetSessions(openf1go.Session{Year: 2023, CountryName: "Belgium"})
if err != nil {
	fmt.Println("Error fetching sessions:", err)
	return
}

for _, session := range sessions {
	fmt.Printf("Session: %s (%s)\n", session.SessionName, session.SessionType)
}
```

#### Example: Fetch Latest Session
```go
session, err := client.GetLatestSessions()
if err != nil {
	fmt.Println("Error fetching latest session:", err)
	return
}

fmt.Printf("Session: %s (%s)\n", session.SessionName, session.SessionType)
}
```

---

### Weather
Fetch weather data for the track.

#### Example: Fetch Latest Weather
```go
weather, err := client.GetLatestWeather()
if err != nil {
	fmt.Println("Error fetching weather:", err)
	return
}

fmt.Printf("Air Temperature: %.2f°C, Track Temperature: %.2f°C\n", weather.AirTemperature, weather.TrackTemperature)
}
```

---

### Team Radio
Fetch team radio communications between drivers and their teams.

#### Example: Fetch Latest Team Radio for All Drivers
```go
teamRadios, err := client.GetAllDriversLatestTeamRadio()
if err != nil {
	fmt.Println("Error fetching team radio:", err)
	return
}

for _, radio := range teamRadios {
	fmt.Printf("Driver %d: %s\n", radio.DriverNumber, radio.RecordingURL)
}
```

---

### Stints
Fetch information about driver stints during a session.

#### Example: Fetch Latest Stints for All Drivers
```go
stints, err := client.GetAllDriversLatestStints()
if err != nil {
	fmt.Println("Error fetching stints:", err)
	return
}

for _, stint := range stints {
	fmt.Printf("Driver %d: Stint %d, Compound: %s\n", stint.DriverNumber, stint.StintNumber, stint.Compound)
}
```

---

### Race Control
Fetch race control events such as flags, safety car deployments, and incidents.

#### Example: Fetch Latest Race Control Events
```go
events, err := client.GetAllDriversLatestRaceControl()
if err != nil {
	fmt.Println("Error fetching race control events:", err)
	return
}

for _, event := range events {
	fmt.Printf("Event: %s, Message: %s\n", event.Category, event.Message)
}
```

---

### Positions
Fetch driver positions during a session.

#### Example: Fetch Latest Positions for All Drivers
```go
positions, err := client.GetAllDriversLatestPositions()
if err != nil {
	fmt.Println("Error fetching positions:", err)
	return
}

for _, position := range positions {
	fmt.Printf("Driver %d: Position %d\n", position.DriverNumber, position.Position)
}
```

---

### Pit Stops
Fetch information about pit stops during a session.

#### Example: Fetch Latest Pit Stops for All Drivers
```go
pits, err := client.GetAllDriversLatestPits()
if err != nil {
	fmt.Println("Error fetching pit stops:", err)
	return
}

for _, pit := range pits {
	fmt.Printf("Driver %d: Pit Duration %.2fs\n", pit.DriverNumber, pit.PitDuration)
}
```

---

### Locations
Fetch the approximate location of cars on the track.

#### Example: Fetch Latest Locations for All Drivers
```go
locations, err := client.GetAllDriversLatestLocations()
if err != nil {
	fmt.Println("Error fetching locations:", err)
	return
}

for _, location := range locations {
	fmt.Printf("Driver %d: X=%d, Y=%d, Z=%d\n", location.DriverNumber, location.X, location.Y, location.Z)
}
```

---

### Laps
Fetch detailed information about individual laps.

#### Example: Fetch Latest Laps for a Driver
```go
driver := openf1go.Driver{DriverNumber: 44}
laps, err := client.GetLatestLapsByDriver(driver)
if err != nil {
	fmt.Println("Error fetching laps:", err)
	return
}

for _, lap := range laps {
	fmt.Printf("Lap %d: Duration %.2fs\n", lap.LapNumber, lap.LapDuration)
}
```

---

### Intervals
Fetch real-time interval data between drivers and their gap to the race leader.

#### Example: Fetch Current Intervals for All Drivers
```go
intervals, err := client.GetAllDriversCurrentIntervals()
if err != nil {
	fmt.Println("Error fetching intervals:", err)
	return
}

for _, interval := range intervals {
	fmt.Printf("Driver %d: Gap to Leader: %s\n", interval.DriverNumber, string(interval.GapToLeader))
}
```

---

### Car Data
Fetch telemetry data for cars, such as speed, RPM, and throttle.

#### Example: Fetch Latest Car Data for a Driver
```go
driver := openf1go.Driver{DriverNumber: 44}
carData, err := client.GetLatestCarDataByDriver(driver)
if err != nil {
	fmt.Println("Error fetching car data:", err)
	return
}

for _, data := range carData {
	fmt.Printf("Speed: %d, RPM: %d\n", data.Speed, data.Rpm)
}
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the MIT License.
