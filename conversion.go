// Package conversion  implements conversion of a Observation into
// a string of three lines.
// Text is formatted according to WRF ob.ascii text format,
// which is described by this FORTRAN specification:
//
//	INFO  = PLATFORM, DATE, NAME, LEVELS, LATITUDE, LONGITUDE, ELEVATION, ID.
//	SRFC  = SLP, PW (DATA,qc,ERROR).
//	EACH  = PRES, SPEED, DIR, HEIGHT, TEMP, DEW PT, HUMID (DATA,qc,ERROR)*LEVELS.
//	INFO_FMT = (A12,1X,A19,1X,A40,1X,I6,3(F12.3,11X),6X,A40)
//	SRFC_FMT = (F12.3,I4,F7.2,F12.3,I4,F7.3)
//	EACH_FMT = (3(F12.3,I4,F7.2),11X,3(F12.3,I4,F7.2),11X,3(F12.3,I4,F7.2))
package magda_drones2wrf

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// qc is
const qc = 0

func str(s string, ln int) string {
	strFmt := fmt.Sprintf("%%-%ds", ln)
	res := fmt.Sprintf(strFmt, s)
	if len(res) <= 40 {
		return res
	}
	return res[0:40]
}

func integer(i int, len int) string {
	intS := fmt.Sprintf("%d", i)
	strFmt := fmt.Sprintf("%%%ds", len)
	return fmt.Sprintf(strFmt, intS)
}

func num(f Value, len float64) string {
	if f.IsNaN() {
		f = -888888.0
	}

	strFmt := fmt.Sprintf("%% %sf", strconv.FormatFloat(float64(len), 'f', -1, 64))
	return fmt.Sprintf(strFmt, f)
}

func space(n int) string {
	return strings.Repeat(" ", n)
}

func date(dt time.Time) string {
	return dt.Format("2006-01-02_15:04:05")
}

func dataQCError(data string, err float64) string {
	qc := qc
	if strings.Contains(data, "-888") {
		qc = -88
	}

	return data +
		integer(qc, 4) +
		num(Value(err), 7.2)
}

func dataQCError3(data string, err float64) string {
	qc := qc
	if strings.Contains(data, "-888") {
		qc = -88
	}

	return data +
		integer(qc, 4) +
		num(Value(err), 7.3)
}

func onlyletters(s string) string {
	res := ""
	for _, rune := range s {
		if unicode.IsLetter(rune) && rune < unicode.MaxASCII {
			res += string(rune)
		} else {
			res += string('X')
		}
	}
	return res
}

// ToWRFASCII converts a Observation into a string
func ToWRFASCII(obs Observation) string {
	firstLine :=
		str("FM-35 TEMP", 12) +
			" " +
			date(obs.ObsTimeUtc) +
			" " +
			str(onlyletters(obs.StationName), 40) +
			" " +
			integer(1, 6) +
			num(Value(obs.Lat), 12.3) +
			space(11) +
			num(Value(obs.Lon), 12.3) +
			space(11) +
			num(Value(obs.Elevation), 12.3) +
			space(11) +
			space(6) +
			str(onlyletters(obs.StationID), 40)

	surfaceLevelPressure := NaN()
	precipTotal := NaN()

	secondLine :=
		dataQCError(num(surfaceLevelPressure, 12.3), 99.99) +
			dataQCError3(num(precipTotal, 12.3), 99.99)

	var thirstLines []string
	for _, m := range obs.Measures {
		thirstLine :=
			dataQCError(num(m.Pressure, 12.3), 1.0) +
				dataQCError(num(m.WindSpeed, 12.3), 1.0) +
				dataQCError(num(m.WindDirection, 12.3), 3.0) +
				space(11) +
				dataQCError(num(m.Altitude, 12.3), 100.00) +
				dataQCError(num(m.Temperature, 12.3), 1) +
				dataQCError(num(m.Dewpoint, 12.3), 1.0) +
				space(11) +
				dataQCError(num(m.Humidity, 12.3), 2)
		thirstLines = append(thirstLines, thirstLine)
	}
	return firstLine + "\n" + secondLine + "\n" + strings.Join(thirstLines, "\n")
}
