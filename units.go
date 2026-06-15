package conversions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// splitNumericUnits separates the numeric prefix of s from its unit suffix.
func splitNumericUnits(s string) (numeric, units string) {
	for _, c := range s {
		if (c >= '0' && c <= '9') || c == '.' {
			numeric += string(c)
		} else {
			units += string(c)
		}
	}
	return numeric, units
}

// prefixMultiplier returns the scale implied by a (lower-cased) unit suffix.
//
// IEC binary prefixes are powers of 1024 and are recognized *before* their SI
// look-alikes so they don't collapse onto them:
//
//	Ki/Mi/Gi/Ti → 2^10 / 2^20 / 2^30 / 2^40   (e.g. MiB, Mibps)
//	K/M/G/T     → 1e3 / 1e6 / 1e9 / 1e12       (e.g. MB,  Mbps)
//
// A suffix with no scale prefix (e.g. "", "b", "bps", "bit", "byte") is 1. ok is
// false for an unrecognized suffix.
func prefixMultiplier(unitsLower string) (mult float64, ok bool) {
	switch {
	case strings.HasPrefix(unitsLower, "ki"):
		return 1 << 10, true
	case strings.HasPrefix(unitsLower, "mi"):
		return 1 << 20, true
	case strings.HasPrefix(unitsLower, "gi"):
		return 1 << 30, true
	case strings.HasPrefix(unitsLower, "ti"):
		return 1 << 40, true
	case strings.HasPrefix(unitsLower, "k"):
		return 1e3, true
	case strings.HasPrefix(unitsLower, "m"):
		return 1e6, true
	case strings.HasPrefix(unitsLower, "g"):
		return 1e9, true
	case strings.HasPrefix(unitsLower, "t"):
		return 1e12, true
	case unitsLower == "" || unitsLower == "b" || unitsLower == "bps" ||
		unitsLower == "bit" || unitsLower == "bits" || unitsLower == "byte" || unitsLower == "bytes":
		return 1, true
	default:
		return 0, false
	}
}

// parseScaled parses "<number><unit>" into number×scale, distinguishing SI from
// IEC prefixes (see prefixMultiplier). It is the shared core of the bit-rate and
// byte-size parsers; the result is overflow-guarded.
func parseScaled(s string) (int64, error) {
	numeric, units := splitNumericUnits(s)
	valuef, err := strconv.ParseFloat(numeric, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse '%v' into a number", s)
	}
	mult, ok := prefixMultiplier(strings.ToLower(strings.TrimSpace(units)))
	if !ok {
		return 0, fmt.Errorf("invalid units specified '%v'", units)
	}
	v := valuef * mult
	if v < 0 || math.IsNaN(v) || math.IsInf(v, 0) || v > math.MaxInt64 {
		return 0, fmt.Errorf("value out of range '%v'", s)
	}
	return int64(math.Round(v)), nil
}
