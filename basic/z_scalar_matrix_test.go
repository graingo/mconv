package basic_test

import (
	"errors"
	"math"
	"testing"
	"time"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/internal"
)

type signedConversion struct {
	name   string
	strict func(interface{}) (int64, error)
	loose  func(interface{}) int64
}

type unsignedConversion struct {
	name   string
	strict func(interface{}) (uint64, error)
	loose  func(interface{}) uint64
}

func TestSignedIntegerConversionMatrix(t *testing.T) {
	converters := []signedConversion{
		{
			name: "int",
			strict: func(value interface{}) (int64, error) {
				converted, err := basic.ToIntE(value)
				return int64(converted), err
			},
			loose: func(value interface{}) int64 { return int64(basic.ToInt(value)) },
		},
		{name: "int64", strict: basic.ToInt64E, loose: basic.ToInt64},
		{
			name: "int32",
			strict: func(value interface{}) (int64, error) {
				converted, err := basic.ToInt32E(value)
				return int64(converted), err
			},
			loose: func(value interface{}) int64 { return int64(basic.ToInt32(value)) },
		},
		{
			name: "int16",
			strict: func(value interface{}) (int64, error) {
				converted, err := basic.ToInt16E(value)
				return int64(converted), err
			},
			loose: func(value interface{}) int64 { return int64(basic.ToInt16(value)) },
		},
		{
			name: "int8",
			strict: func(value interface{}) (int64, error) {
				converted, err := basic.ToInt8E(value)
				return int64(converted), err
			},
			loose: func(value interface{}) int64 { return int64(basic.ToInt8(value)) },
		},
	}
	inputs := []struct {
		name  string
		value interface{}
		want  int64
	}{
		{name: "int", value: int(12), want: 12},
		{name: "int64", value: int64(12), want: 12},
		{name: "int32", value: int32(12), want: 12},
		{name: "int16", value: int16(12), want: 12},
		{name: "int8", value: int8(12), want: 12},
		{name: "uint", value: uint(12), want: 12},
		{name: "uint64", value: uint64(12), want: 12},
		{name: "uint32", value: uint32(12), want: 12},
		{name: "uint16", value: uint16(12), want: 12},
		{name: "uint8", value: uint8(12), want: 12},
		{name: "float64", value: float64(12.75), want: 12},
		{name: "float32", value: float32(12.75), want: 12},
		{name: "complex64", value: complex64(12.75), want: 12},
		{name: "complex128", value: complex128(12.75), want: 12},
		{name: "bool", value: true, want: 1},
		{name: "string", value: " 12 ", want: 12},
		{name: "nil", value: nil, want: 0},
	}

	for _, converter := range converters {
		converter := converter
		t.Run(converter.name, func(t *testing.T) {
			for _, input := range inputs {
				input := input
				t.Run(input.name, func(t *testing.T) {
					got, err := converter.strict(input.value)
					if err != nil {
						t.Fatalf("strict conversion failed: %v", err)
					}
					if got != input.want {
						t.Fatalf("got %d, want %d", got, input.want)
					}
				})
			}

			if got := converter.loose("invalid"); got != 0 {
				t.Fatalf("loose conversion returned %d for invalid input", got)
			}
			for _, value := range []interface{}{complex(1, 1), struct{}{}} {
				if _, err := converter.strict(value); err == nil {
					t.Fatalf("strict conversion accepted %#v", value)
				}
			}
		})
	}
}

func TestSignedIntegerWidthBoundaries(t *testing.T) {
	tests := []struct {
		name    string
		convert func() error
	}{
		{name: "int8", convert: func() error { _, err := basic.ToInt8E(int64(128)); return err }},
		{name: "int16", convert: func() error { _, err := basic.ToInt16E(int32(32768)); return err }},
		{name: "int32", convert: func() error { _, err := basic.ToInt32E(int64(1 << 31)); return err }},
		{name: "int64", convert: func() error { _, err := basic.ToInt64E(uint64(math.MaxUint64)); return err }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.convert(); !errors.Is(err, internal.ErrOverflow) {
				t.Fatalf("got %v, want ErrOverflow", err)
			}
		})
	}
}

func TestUnsignedIntegerConversionMatrix(t *testing.T) {
	converters := []unsignedConversion{
		{
			name: "uint",
			strict: func(value interface{}) (uint64, error) {
				converted, err := basic.ToUintE(value)
				return uint64(converted), err
			},
			loose: func(value interface{}) uint64 { return uint64(basic.ToUint(value)) },
		},
		{name: "uint64", strict: basic.ToUint64E, loose: basic.ToUint64},
		{
			name: "uint32",
			strict: func(value interface{}) (uint64, error) {
				converted, err := basic.ToUint32E(value)
				return uint64(converted), err
			},
			loose: func(value interface{}) uint64 { return uint64(basic.ToUint32(value)) },
		},
		{
			name: "uint16",
			strict: func(value interface{}) (uint64, error) {
				converted, err := basic.ToUint16E(value)
				return uint64(converted), err
			},
			loose: func(value interface{}) uint64 { return uint64(basic.ToUint16(value)) },
		},
		{
			name: "uint8",
			strict: func(value interface{}) (uint64, error) {
				converted, err := basic.ToUint8E(value)
				return uint64(converted), err
			},
			loose: func(value interface{}) uint64 { return uint64(basic.ToUint8(value)) },
		},
	}
	inputs := []struct {
		name  string
		value interface{}
		want  uint64
	}{
		{name: "uint", value: uint(12), want: 12},
		{name: "uint64", value: uint64(12), want: 12},
		{name: "uint32", value: uint32(12), want: 12},
		{name: "uint16", value: uint16(12), want: 12},
		{name: "uint8", value: uint8(12), want: 12},
		{name: "int", value: int(12), want: 12},
		{name: "int64", value: int64(12), want: 12},
		{name: "int32", value: int32(12), want: 12},
		{name: "int16", value: int16(12), want: 12},
		{name: "int8", value: int8(12), want: 12},
		{name: "float64", value: float64(12.75), want: 12},
		{name: "float32", value: float32(12.75), want: 12},
		{name: "complex64", value: complex64(12.75), want: 12},
		{name: "complex128", value: complex128(12.75), want: 12},
		{name: "bool", value: true, want: 1},
		{name: "string", value: " 12 ", want: 12},
		{name: "nil", value: nil, want: 0},
	}

	for _, converter := range converters {
		converter := converter
		t.Run(converter.name, func(t *testing.T) {
			for _, input := range inputs {
				input := input
				t.Run(input.name, func(t *testing.T) {
					got, err := converter.strict(input.value)
					if err != nil {
						t.Fatalf("strict conversion failed: %v", err)
					}
					if got != input.want {
						t.Fatalf("got %d, want %d", got, input.want)
					}
				})
			}

			if got := converter.loose("invalid"); got != 0 {
				t.Fatalf("loose conversion returned %d for invalid input", got)
			}
			for _, value := range []interface{}{int64(-1), complex(1, 1), struct{}{}} {
				if _, err := converter.strict(value); err == nil {
					t.Fatalf("strict conversion accepted %#v", value)
				}
			}
		})
	}
}

func TestUnsignedIntegerWidthBoundaries(t *testing.T) {
	tests := []struct {
		name    string
		convert func() error
	}{
		{name: "uint8", convert: func() error { _, err := basic.ToUint8E(uint16(256)); return err }},
		{name: "uint16", convert: func() error { _, err := basic.ToUint16E(uint32(1 << 16)); return err }},
		{name: "uint32", convert: func() error { _, err := basic.ToUint32E(uint64(1 << 32)); return err }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.convert(); !errors.Is(err, internal.ErrOverflow) {
				t.Fatalf("got %v, want ErrOverflow", err)
			}
		})
	}
}

func TestRemainingScalarBranches(t *testing.T) {
	boolInputs := []struct {
		value interface{}
		want  bool
	}{
		{value: int8(0), want: false},
		{value: int16(1), want: true},
		{value: uint8(0), want: false},
		{value: uint16(1), want: true},
		{value: uint32(0), want: false},
		{value: uint64(1), want: true},
		{value: float32(0), want: false},
		{value: complex64(1), want: true},
	}
	for _, input := range boolInputs {
		got, err := basic.ToBoolE(input.value)
		if err != nil || got != input.want {
			t.Fatalf("ToBoolE(%#v) = %v, %v; want %v", input.value, got, err, input.want)
		}
	}

	if got := basic.ToFloat32("1.5"); got != 1.5 {
		t.Fatalf("ToFloat32 returned %v", got)
	}
	if got := basic.ToComplex128(uint64(7)); got != 7 {
		t.Fatalf("ToComplex128 returned %v", got)
	}
	if got := basic.ToComplex64(int16(7)); got != 7 {
		t.Fatalf("ToComplex64 returned %v", got)
	}

	durationInputs := []interface{}{int8(1), int16(1), int32(1), uint8(1), uint16(1), uint32(1), uint64(1), float32(1)}
	for _, input := range durationInputs {
		got, err := basic.ToDurationE(input)
		if err != nil || got != time.Nanosecond {
			t.Fatalf("ToDurationE(%#v) = %v, %v", input, got, err)
		}
	}
}
