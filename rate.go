package conversions

import (
	"strconv"
	"time"
)

type BinaryPrefix int

const (
	PREFIX_BITS BinaryPrefix = 1
	PREFIX_KILOBITS BinaryPrefix = 1 * 1000
	PREFIX_MEGABITS BinaryPrefix = 1 * 1000 * 1000
	PREFIX_GIGABITS BinaryPrefix = 1 * 1000 * 1000 * 1000
	PREFIX_TERABITS BinaryPrefix = 1 * 1000 * 1000 * 1000 * 1000

	BITS_PER_BYTE = 8
)

func ConvertStringToUInt64orNil(uintStr string) *uint64 {
	value, err := strconv.ParseUint(uintStr, 10, 64)
	if err != nil {
		return nil
	}
	return &value
}

func ConvertStringToUint64or0(uintStr string) uint64 {
	value := ConvertStringToUInt64orNil(uintStr)
	if value == nil {
		return 0
	}
	return *value
}

func ConvertToRateMbps(lastBytes uint64, currentBytes uint64, lastTime int64, currentTime int64) float32 {
	return convertByteCountsToRate(lastBytes, currentBytes, lastTime, currentTime, PREFIX_MEGABITS)
}

func convertByteCountsToRate(lastBytes uint64, currentBytes uint64, lastTime int64, currentTime int64, units BinaryPrefix) float32 {
	period := float32(currentTime - lastTime) / float32(int64(time.Second))
	change := (currentBytes - lastBytes) * BITS_PER_BYTE
	rate := float32(change) / period
	return rate / float32(units)
}