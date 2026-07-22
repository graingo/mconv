package basic_test

import (
	"math"
	"strconv"
	"testing"
	"time"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/internal"
)

func TestIntegerConversionsRejectNonFiniteAndRoundedBoundaries(t *testing.T) {
	signed := []func(interface{}) error{
		func(v interface{}) error { _, err := basic.ToIntE(v); return err },
		func(v interface{}) error { _, err := basic.ToInt64E(v); return err },
		func(v interface{}) error { _, err := basic.ToInt32E(v); return err },
		func(v interface{}) error { _, err := basic.ToInt16E(v); return err },
		func(v interface{}) error { _, err := basic.ToInt8E(v); return err },
	}
	unsigned := []func(interface{}) error{
		func(v interface{}) error { _, err := basic.ToUintE(v); return err },
		func(v interface{}) error { _, err := basic.ToUint64E(v); return err },
		func(v interface{}) error { _, err := basic.ToUint32E(v); return err },
		func(v interface{}) error { _, err := basic.ToUint16E(v); return err },
		func(v interface{}) error { _, err := basic.ToUint8E(v); return err },
	}
	for _, convert := range append(signed, unsigned...) {
		if err := convert(math.NaN()); err == nil {
			t.Fatal("NaN conversion should fail")
		}
	}
	if _, err := basic.ToInt64E(float64(math.MaxInt64)); err == nil {
		t.Fatal("rounded int64 upper boundary should fail")
	}
	if _, err := basic.ToUint64E(float64(math.MaxUint64)); err == nil {
		t.Fatal("rounded uint64 upper boundary should fail")
	}
	if strconv.IntSize == 64 {
		if _, err := basic.ToInt64E(^uint(0)); err == nil {
			t.Fatal("uint larger than MaxInt64 should fail")
		}
	}
}

func TestTimeCacheIncludesParsingFormat(t *testing.T) {
	internal.ClearTimeCache()
	value := "01/02/2006"
	if _, err := basic.ToTimeE(value, "01/02/2006"); err != nil {
		t.Fatalf("first format should parse: %v", err)
	}
	if _, err := basic.ToTimeE(value, "2006-01-02"); err == nil {
		t.Fatal("a cached result from a different format must not be reused")
	}
	if _, err := basic.ToTimeE(uint64(math.MaxUint64)); err == nil {
		t.Fatal("uint64 timestamp overflow should fail")
	}
	if _, err := basic.ToTimeE("2026-07-22", time.RFC3339, "2006-01-02"); err != nil {
		t.Fatalf("all supplied time formats should be tried: %v", err)
	}
	if _, err := basic.ToDurationE(uint64(math.MaxUint64)); err == nil {
		t.Fatal("uint64 duration overflow should fail")
	}
}

func TestFloatAndComplexNarrowingChecksRealOverflow(t *testing.T) {
	if _, err := basic.ToFloat32E(math.MaxFloat64); err == nil {
		t.Fatal("float64 outside float32 range should fail")
	}
	value, err := basic.ToComplex64E(complex(-1.1, -2.2))
	if err != nil {
		t.Fatalf("representable negative complex value should succeed: %v", err)
	}
	if real(value) >= 0 || imag(value) >= 0 {
		t.Fatalf("unexpected converted value: %v", value)
	}
	if _, err := basic.ToComplex64E(complex(math.MaxFloat64, 0)); err == nil {
		t.Fatal("complex component outside float32 range should fail")
	}
}

func TestNumericStringsAllowSurroundingWhitespace(t *testing.T) {
	if value, err := basic.ToInt64E(" 42 "); err != nil || value != 42 {
		t.Fatalf("int conversion failed: %d, %v", value, err)
	}
	if value, err := basic.ToFloat64E(" 1.5 "); err != nil || value != 1.5 {
		t.Fatalf("float conversion failed: %v, %v", value, err)
	}
	if value, err := basic.ToBoolE(" yes "); err != nil || !value {
		t.Fatalf("bool conversion failed: %v, %v", value, err)
	}
}
