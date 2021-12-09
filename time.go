package conversions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func ConvertNanosecondsToStringTime(nanoseconds int64) string {
	units := "ns"
	t := float64(nanoseconds)
	if t > 1000 {
		units = "us"
		t = t / 1000
	}
	if t > 1000 {
		units = "ms"
		t = t / 1000
	}
	if t > 1000 {
		units = "s"
		t = t / 1000
	}

	return fmt.Sprintf("%f %s", t, units)
}

func ConvertStringTimeToNanoseconds(value string) (t int64, err error) {

	numeric := ""
	units := ""
	valuef := 0.0
	multiplier := 1.0
	for _, c := range value {
		switch {
		case c >= '0' && c <= '9' || c == '.':
			numeric += string(c)
		default:
			units += string(c)
		}
	}
	valuef, err = strconv.ParseFloat(numeric, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse '%v' into float", value)
	}
	unitsLower := strings.ToLower(strings.TrimSpace(units))
	switch {
	case strings.HasPrefix(unitsLower, "ns"):
		multiplier = float64(1 * time.Nanosecond)
	case strings.HasPrefix(unitsLower, "µs"):
		multiplier = float64(1 * time.Microsecond)
	case strings.HasPrefix(unitsLower, "ms"):
		multiplier = float64(1 * time.Millisecond)
	case strings.HasPrefix(unitsLower, "s"):
		multiplier = float64(1 * time.Second)
	case strings.HasPrefix(unitsLower, "m"):
		multiplier = float64(1 * time.Minute)
	case strings.HasPrefix(unitsLower, "h"):
		multiplier = float64(1 * time.Hour)
	case strings.HasPrefix(unitsLower, "d"):
		multiplier = float64(1 * time.Hour * 24)
	}
	t = int64(math.Round(valuef * multiplier))
	return t, err
}
