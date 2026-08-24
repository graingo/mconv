package mconv_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/graingo/mconv"
)

func BenchmarkStringCache(b *testing.B) {
	for _, size := range []int{0, 1000} {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			mconv.SetStringCacheSize(size)
			b.Cleanup(func() {
				mconv.SetStringCacheSize(0)
			})
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = mconv.ToStringE(42)
			}
		})
	}
}

func BenchmarkTimeCacheRepeatedValue(b *testing.B) {
	for _, size := range []int{0, 100} {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			mconv.SetTimeCacheSize(size)
			b.Cleanup(func() {
				mconv.SetTimeCacheSize(0)
			})
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = mconv.ToTimeE("2026-08-24T12:00:00Z")
			}
		})
	}
}

func BenchmarkTimeCacheHighCardinality(b *testing.B) {
	values := make([]string, 1000)
	start := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := range values {
		values[i] = start.Add(time.Duration(i) * time.Second).Format(time.RFC3339)
	}

	for _, size := range []int{0, 100} {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			mconv.SetTimeCacheSize(size)
			b.Cleanup(func() {
				mconv.SetTimeCacheSize(0)
			})
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, _ = mconv.ToTimeE(values[i%len(values)])
			}
		})
	}
}
