package magda_drones2wrf

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"time"
)

func parseFloat(s string) (Value, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return NaN(), err
	}
	return Value(f), nil
}

func ReadAll(dataPath string) (Observation, error) {
	observation := Observation{}

	obsF, err := os.Open(dataPath)
	if err != nil {
		return observation, err
	}
	defer obsF.Close()
	csvReader := csv.NewReader(obsF)
	csvReader.Comma = ';'
	var data [][]string
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return observation, err
		}
		data = append(data, record)
	}

	const PRESSURE = 3
	const ALTITUDE = 4
	const TEMPERATURE = 5
	const HUMIDITY = 6
	const WIND_SPEED = 7
	const WIND_DIRECTION = 8

	const HEAD_DATE = 0
	const HEAD_LAT = 1
	const HEAD_LON = 2

	header := data[1]
	obsTime, err := time.Parse("2006-01-02 15:04", header[HEAD_DATE])
	if err != nil {
		return observation, err
	}
	observation.ObsTimeUtc = obsTime

	observation.StationName = "XXX"
	observation.StationID = "XXX"
	observation.Elevation = ConfigValues.Elevation
	data = data[1:]
	observation.Lat, err = strconv.ParseFloat(header[HEAD_LAT], 64)
	if err != nil {
		return observation, err
	}
	observation.Lon, err = strconv.ParseFloat(header[HEAD_LON], 64)
	if err != nil {
		return observation, err
	}

	observation.Measures = make([]Measure, len(data))
	for i, row := range data {
		var m Measure
		var err error
		m.Dewpoint = NaN()
		m.Temperature, err = parseFloat(row[TEMPERATURE])
		if err != nil {
			return observation, err
		}

		m.WindSpeed, err = parseFloat(row[WIND_SPEED])
		if err != nil {
			return observation, err
		}
		m.WindDirection, err = parseFloat(row[WIND_DIRECTION])
		if err != nil {
			return observation, err
		}
		m.Pressure, err = parseFloat(row[PRESSURE])
		if err != nil {
			return observation, err
		}
		m.Precipitation = NaN()
		m.Humidity, err = parseFloat(row[HUMIDITY])
		if err != nil {
			return observation, err
		}
		m.Altitude, err = parseFloat(row[ALTITUDE])
		if err != nil {
			return observation, err
		}
		observation.Measures[i] = m

	}

	return observation, nil
}
