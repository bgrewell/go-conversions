package conversions

import (
	"fmt"
)

type UnitsFormatter interface {
	Format(units string) string
}

type DefaultFormatter struct {
}

func (df DefaultFormatter) Format(units string) string {
	return fmt.Sprintf("%sbit", units)
}

// StringBitRateToInt parses a bit-rate string to bits/sec. It distinguishes SI
// decimal prefixes (K/M/G/T → 1e3/1e6/1e9/1e12, e.g. "100Mbps") from IEC binary
// prefixes (Ki/Mi/Gi/Ti → 2^10/2^20/2^30/2^40, e.g. "100Mibps"), so 100Mbps and
// 100Mibps are not the same. The "bps"/"bit" suffix is optional, and the result
// is overflow-guarded.
func StringBitRateToInt(rate string) (value int64, err error) {
	return parseScaled(rate)
}

func IntBitRateToString(rate int64) string {
	return IntBitRateToStringFmt(rate, DefaultFormatter{})
}

// IntBitRateToStringIEC renders bits/sec with IEC binary prefixes (Kibit, Mibit,
// …, powers of 1024).
func IntBitRateToStringIEC(rate int64) string {
	v, suffix := humanize(float64(rate), 1024, []string{"", "Ki", "Mi", "Gi", "Ti"})
	return fmt.Sprintf("%.2f%sbit", v, suffix)
}

func IntBitRateToStringFmt(rate int64, formatter UnitsFormatter) string {
	units := "b"
	value := float64(rate)
	if value >= 1000 {
		units = "k"
		value = value / 1000
	}
	if value >= 1000 {
		units = "m"
		value = value / 1000
	}
	if value >= 1000 {
		units = "g"
		value = value / 1000
	}
	if value >= 1000 {
		units = "t"
		value = value / 1000
	}
	suffix := formatter.Format(units)
	srate := fmt.Sprintf("%.2f%s", value, suffix)
	return srate
}
