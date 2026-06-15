package conversions

import "testing"

func TestStringBitRateToInt(t *testing.T) {
	cases := map[string]int64{
		"1000":     1000,
		"100":      100,
		"100bps":   100,
		"1Kbps":    1000,
		"1K":       1000,
		"100Mbps":  100_000_000, // SI decimal
		"1.5G":     1_500_000_000,
		"1Gbps":    1_000_000_000,
		"1Kibps":   1024,        // IEC binary
		"100Mibps": 104_857_600, // 100 * 2^20
		"1Gibps":   1 << 30,
	}
	for in, want := range cases {
		got, err := StringBitRateToInt(in)
		if err != nil || got != want {
			t.Errorf("StringBitRateToInt(%q) = %d, %v; want %d", in, got, err, want)
		}
	}
	// The crux: Mbps and Mibps must differ.
	mbps, _ := StringBitRateToInt("100Mbps")
	mibps, _ := StringBitRateToInt("100Mibps")
	if mbps == mibps {
		t.Errorf("100Mbps (%d) must differ from 100Mibps (%d)", mbps, mibps)
	}
	for _, bad := range []string{"abc", "10Z", "1e9", "99999999999G"} {
		if _, err := StringBitRateToInt(bad); err == nil {
			t.Errorf("StringBitRateToInt(%q) should error", bad)
		}
	}
}

func TestStringByteSizeToInt(t *testing.T) {
	cases := map[string]int64{
		"1000":   1000,
		"1KB":    1000, // SI decimal
		"1K":     1000,
		"100MB":  100_000_000,
		"1.5MB":  1_500_000,
		"1GB":    1_000_000_000,
		"1KiB":   1024, // IEC binary
		"100MiB": 104_857_600,
		"1GiB":   1 << 30,
		"512B":   512,
	}
	for in, want := range cases {
		got, err := StringByteSizeToInt(in)
		if err != nil || got != want {
			t.Errorf("StringByteSizeToInt(%q) = %d, %v; want %d", in, got, err, want)
		}
	}
	// The crux: MB and MiB must differ.
	mb, _ := StringByteSizeToInt("100MB")
	mib, _ := StringByteSizeToInt("100MiB")
	if mb == mib {
		t.Errorf("100MB (%d) must differ from 100MiB (%d)", mb, mib)
	}
	for _, bad := range []string{"", "abc", "10Z", "99999999999999999999G"} {
		if _, err := StringByteSizeToInt(bad); err == nil {
			t.Errorf("StringByteSizeToInt(%q) should error", bad)
		}
	}
}

func TestByteSizeFormatting(t *testing.T) {
	if got := IntByteSizeToString(1_500_000); got != "1.50MB" {
		t.Errorf("IntByteSizeToString = %q, want 1.50MB", got)
	}
	if got := IntByteSizeToStringIEC(1_572_864); got != "1.50MiB" {
		t.Errorf("IntByteSizeToStringIEC = %q, want 1.50MiB", got)
	}
	if got := IntByteSizeToString(512); got != "512.00B" {
		t.Errorf("IntByteSizeToString(512) = %q, want 512.00B", got)
	}
}

// Round-trip: parse then format (SI) stays in the same family.
func TestRoundTrip(t *testing.T) {
	n, _ := StringByteSizeToInt("250MB")
	if got := IntByteSizeToString(n); got != "250.00MB" {
		t.Errorf("round-trip 250MB -> %d -> %q", n, got)
	}
}
