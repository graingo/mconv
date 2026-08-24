package basic_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func FuzzStringConversions(f *testing.F) {
	for _, seed := range []string{
		"",
		"0",
		"-1",
		"9223372036854775807",
		"18446744073709551615",
		"NaN",
		"+Inf",
		"true",
		"2026-08-24T12:00:00Z",
	} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, value string) {
		_, _ = mconv.ToBoolE(value)
		_, _ = mconv.ToInt64E(value)
		_, _ = mconv.ToUint64E(value)
		_, _ = mconv.ToFloat64E(value)
		_, _ = mconv.ToComplex128E(value)
		_, _ = mconv.ToTimeE(value)
		_, _ = mconv.ToDurationE(value)
	})
}

func FuzzIntegerBoundaries(f *testing.F) {
	for _, seed := range []int64{
		-1 << 63,
		-1 << 31,
		-1,
		0,
		1,
		1<<31 - 1,
		1<<63 - 1,
	} {
		f.Add(seed)
	}

	f.Fuzz(func(t *testing.T, value int64) {
		_, _ = mconv.ToInt8E(value)
		_, _ = mconv.ToInt16E(value)
		_, _ = mconv.ToInt32E(value)
		_, _ = mconv.ToIntE(value)
		_, _ = mconv.ToUint8E(value)
		_, _ = mconv.ToUint16E(value)
		_, _ = mconv.ToUint32E(value)
		_, _ = mconv.ToUint64E(value)
		_, _ = mconv.ToFloat32E(value)
	})
}
