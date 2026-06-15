package conversions

import "fmt"

// StringByteSizeToInt parses a byte-size string to bytes. It distinguishes SI
// decimal prefixes (KB/MB/GB/TB and bare K/M/G/T → 1e3/1e6/1e9/1e12) from IEC
// binary prefixes (KiB/MiB/GiB/TiB → 2^10/2^20/2^30/2^40), so "100MB" (10⁶) and
// "100MiB" (2²⁰) are not the same. A trailing "B"/"byte" is optional, and the
// result is overflow-guarded.
//
//	StringByteSizeToInt("100MB")  -> 100000000
//	StringByteSizeToInt("100MiB") -> 104857600
func StringByteSizeToInt(size string) (int64, error) {
	return parseScaled(size)
}

// IntByteSizeToString renders a byte count with SI decimal prefixes (KB, MB, …,
// powers of 1000), e.g. 1500000 -> "1.50MB".
func IntByteSizeToString(size int64) string {
	v, suffix := humanize(float64(size), 1000, []string{"", "K", "M", "G", "T"})
	return fmt.Sprintf("%.2f%sB", v, suffix)
}

// IntByteSizeToStringIEC renders a byte count with IEC binary prefixes (KiB,
// MiB, …, powers of 1024), e.g. 1572864 -> "1.50MiB".
func IntByteSizeToStringIEC(size int64) string {
	v, suffix := humanize(float64(size), 1024, []string{"", "Ki", "Mi", "Gi", "Ti"})
	return fmt.Sprintf("%.2f%sB", v, suffix)
}

// humanize scales value down by base, stepping through suffixes, until it is
// below base (or the suffixes run out). Shared by the size/rate formatters.
func humanize(value, base float64, suffixes []string) (float64, string) {
	i := 0
	for value >= base && i < len(suffixes)-1 {
		value /= base
		i++
	}
	return value, suffixes[i]
}
