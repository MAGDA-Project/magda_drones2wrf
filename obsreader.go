package magda_ws2wrf

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
	HumidityAvg Value
	WinddirAvg  Value
	Metric      ObservationMetric
}

// ObservationMetric contains a subset of values
// contained in an Observation
type ObservationMetric struct {
	TempAvg      Value
	DewptAvg     Value
	WindspeedAvg Value
	Pressure     Value
	PrecipTotal  Value
	PressureMin  Value
	PressureMax  Value
}

// ObsReader is implemented by types that
// are ables to read `Observation`.
type ObsReader interface {
	// ReadAll returns a slice of types.Observation read
	// from path argument, filtered by `domain` and
	// `date` arguments.
	// If an error occurred, it is returned as second value,
	// with the first one nil.
	ReadAll(path string, domain Domain, date time.Time) ([]Observation, error)
}
