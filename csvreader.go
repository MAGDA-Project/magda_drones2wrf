package magda_drones2wrf

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
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

		if err != nil && errors.Unwrap(err) != csv.ErrFieldCount {
			return observation, err
		}
		data = append(data, record)
	}

	header := data[0]
	obsTime, err := time.Parse("2006-01-02 15:04:05", header[0][0:19])
	if err != nil {
		return observation, err
	}
	observation.ObsTimeUtc = obsTime

	observation.StationName = strings.TrimSpace(header[1])
	observation.StationID = "XXX"
	observation.Elevation = 0
	data = data[2:]
	observation.Lat, err = strconv.ParseFloat(data[0][0], 64)
	if err != nil {
		return observation, err
	}
	observation.Lon, err = strconv.ParseFloat(data[0][1], 64)
	if err != nil {
		return observation, err
	}

	const ALTITUDE = 2
	const TEMPERATURE = 3
	const DEW_POINT = 4
	const HUMIDITY = 5
	const PRESSURE = 6
	const WIND_SPEED = 7
	const WIND_DIRECTION = 8

	observation.Measures = make([]Measure, len(data))
	for i, row := range data {
		var m Measure
		var err error
		m.Temperature, err = parseFloat(row[TEMPERATURE])
		if err != nil {
			return observation, err
		}
		m.Dewpoint, err = parseFloat(row[DEW_POINT])
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
