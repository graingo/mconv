package internal

import (
	"errors"
	"math"
	"testing"
)

func TestBasicCacheKeySupportsEveryCachedScalar(t *testing.T) {
	values := []interface{}{
		"value",
		true,
		false,
		int(1),
		int64(1),
		int32(1),
		int16(1),
		int8(1),
		uint(1),
		uint64(1),
		uint32(1),
		uint16(1),
		uint8(1),
		float64(1),
		float32(1),
	}
	for _, value := range values {
		if _, ok := basicCacheKey(value); !ok {
			t.Fatalf("expected a cache key for %T", value)
		}
	}
	if _, ok := basicCacheKey(struct{}{}); ok {
		t.Fatal("struct values should not be cached")
	}

	positiveZero, ok := basicCacheKey(float64(0))
	if !ok {
		t.Fatal("positive zero should have a cache key")
	}
	negativeZero, ok := basicCacheKey(math.Copysign(0, -1))
	if !ok {
		t.Fatal("negative zero should have a cache key")
	}
	if positiveZero == negativeZero {
		t.Fatal("positive and negative zero require different string-cache keys")
	}
}

func TestTimeCacheKeyIncludesEveryFormat(t *testing.T) {
	defaultKey, ok := timeCacheKey("value", nil)
	if !ok || defaultKey.formatCount != 0 {
		t.Fatalf("unexpected default key: %#v", defaultKey)
	}

	singleKey, ok := timeCacheKey("value", []string{"2006-01-02"})
	if !ok || singleKey.singleFormat != "2006-01-02" {
		t.Fatalf("unexpected single-format key: %#v", singleKey)
	}

	first, ok := timeCacheKey("value", []string{"1", "23"})
	if !ok {
		t.Fatal("multi-format key was rejected")
	}
	second, ok := timeCacheKey("value", []string{"12", "3"})
	if !ok {
		t.Fatal("multi-format key was rejected")
	}
	if first == second {
		t.Fatal("length-prefixed formats must not collide")
	}
	if _, ok := timeCacheKey(struct{}{}, nil); ok {
		t.Fatal("unsupported values should not receive a time-cache key")
	}
}

func TestPrependConversionPath(t *testing.T) {
	base := NewConversionError("invalid", "int", ErrConversionFailed)
	err := PrependConversionPath(base, "Age")
	err = PrependConversionPath(err, "[0]")
	err = PrependConversionPath(err, "Users")

	var conversionErr *ConversionError
	if !errors.As(err, &conversionErr) {
		t.Fatalf("expected ConversionError, got %T", err)
	}
	if conversionErr.Path != "Users[0].Age" {
		t.Fatalf("unexpected path: %q", conversionErr.Path)
	}
	if PrependConversionPath(nil, "Field") != nil {
		t.Fatal("nil errors should remain nil")
	}
	if PrependConversionPath(base, "") != base {
		t.Fatal("an empty segment should preserve the original error")
	}

	plain := errors.New("plain")
	wrapped := PrependConversionPath(plain, "Field")
	if !errors.Is(wrapped, plain) {
		t.Fatal("plain errors should remain discoverable")
	}
}
