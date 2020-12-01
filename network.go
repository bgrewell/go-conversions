package conversions

func StringBitRateToInt(rate string) (value int64, err error) {
	numeric := ""
	units := ""
	valuef := 0.0
	multiplier := 1.0
	for _, c := range rate {
		switch {
		case c >= '0' && c <= '9' || c == '.':
			numeric += string(c)
		default:
			units += string(c)
		}
	}
	valuef, err = strconv.ParseFloat(numeric, 64)
	if err != nil {
		return 0, fmt.Errorf("couldn't parse '%v' into float", rate)
	}
	unitsLower := strings.ToLower(units)
	switch {
	case strings.HasPrefix(unitsLower, "k"):
		multiplier = 1000
	case strings.HasPrefix(unitsLower, "m"):
		multiplier = 1000 * 1000
	case strings.HasPrefix(unitsLower, "g"):
		multiplier = 1000 * 1000 * 1000
	case unitsLower == "":
		multiplier = 1
	default:
		return 0, fmt.Errorf("invalid units specified '%v'", units)
	}
	value = int64(math.Round(valuef * multiplier))
	return value, err
}

func IntBitRateToString(rate int64) string {
	suffix := "bit"
	value := float64(rate)
	if value >= 1000 {
		suffix = "kbit"
		value = value / 1000
	}
	if value >= 1000 {
		suffix = "mbit"
		value = value / 1000
	}
	if value >= 1000 {
		suffix = "gbit"
		value = value / 1000
	}
	if value >= 1000 {
		suffix = "tbit"
		value = value / 1000
	}
	srate := fmt.Sprintf("%.2f%s", value, suffix)
	log.WithFields(log.Fields{
		"rate":   rate,
		"value":  value,
		"suffix": suffix,
		"srate":  srate,
	}).Debug("results of bits to string conversion")
	return srate
}
