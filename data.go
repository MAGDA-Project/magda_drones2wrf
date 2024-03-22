package magda_drones2wrf

import (
	"time"
)

// Observation represents data for all sensor classes of
// a station at a moment in time
type Observation struct {
	Elevation   float64
	StationID   string
	StationName string
	ObsTimeUtc  time.Time
	Lat, Lon    float64

	Measures []Measure
}

// Measure contains a subset of values
// contained in an Observation
type Measure struct {
	Temperature   Value
	Dewpoint      Value
	WindSpeed     Value
	Pressure      Value
	Precipitation Value
	Humidity      Value
	WindDirection Value
	Altitude      Value
}
