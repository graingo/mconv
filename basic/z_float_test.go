package basic_test

import (
	"math"
	"testing"

	"github.com/graingo/mconv"
)

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
	}{
		{123, 123.0},
		{int64(123), 123.0},
		{123.45, 123.45},
		{float32(123.45), 123.45},
		{"123.45", 123.45},
		{true, 1.0},
		{false, 0.0},
		{nil, 0.0},
		{complex(1.2, 0), 1.2},
	}

	for _, test := range tests {
		got := mconv.ToFloat64(test.input)
		if f32, ok := test.input.(float32); ok {
			if math.Abs(got-float64(f32)) > 1e-6 {
				t.Errorf("mconv.ToFloat64(%v) = %v; want %v", test.input, got, test.expected)
			}
		} else {
			if math.Abs(got-test.expected) > 1e-6 {
				t.Errorf("mconv.ToFloat64(%v) = %v; want %v", test.input, got, test.expected)
			}
		}
	}
}

func TestToFloat64E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
		isErr    bool
	}{
		{123, 123.0, false},
		{int64(12345), 12345.0, false},
		{int32(123), 123.0, false},
		{int16(123), 123.0, false},
		{int8(123), 123.0, false},
		{uint(123), 123.0, false},
		{uint64(123), 123.0, false},
		{uint32(123), 123.0, false},
		{uint16(123), 123.0, false},
		{uint8(123), 123.0, false},
		{123.45, 123.45, false},
		{float32(123.45), 123.45, false},
		{complex64(complex(1.2, 3.4)), 1.2, false},
		{complex128(complex(5.6, 7.8)), 5.6, false},
		{"123.45", 123.45, false},
		{"abc", 0.0, true},
		{true, 1.0, false},
		{false, 0.0, false},
		{nil, 0.0, false},
		{struct{}{}, 0.0, true},
	}

	for _, test := range tests {
		got, err := mconv.ToFloat64E(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToFloat64E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToFloat64E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr {
			if f32, ok := test.input.(float32); ok {
				if math.Abs(got-float64(f32)) > 1e-6 {
					t.Errorf("mconv.ToFloat64E(%v) = %v; want %v", test.input, got, test.expected)
				}
			} else {
				if math.Abs(got-test.expected) > 1e-6 {
					t.Errorf("mconv.ToFloat64E(%v) = %v; want %v", test.input, got, test.expected)
				}
			}
		}
	}
}

func TestToFloat32E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float32
		isErr    bool
	}{
		{123, 123.0, false},
		{int64(123), 123.0, false},
		{int32(123), 123.0, false},
		{int16(123), 123.0, false},
		{int8(123), 123.0, false},
		{uint(123), 123.0, false},
		{uint64(123), 123.0, false},
		{uint32(123), 123.0, false},
		{uint16(123), 123.0, false},
		{uint8(123), 123.0, false},
		{123.45, float32(123.45), false},
		{float64(123.45), float32(123.45), false},
		{complex64(complex(1.2, 3.4)), 1.2, false},
		{complex128(complex(5.6, 7.8)), 5.6, false},
		{"123.45", 123.45, false},
		{"abc", 0.0, true},
		{true, 1.0, false},
		{false, 0.0, false},
		{nil, 0.0, false},
		{struct{}{}, 0.0, true},
	}

	for _, test := range tests {
		got, err := mconv.ToFloat32E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToFloat32E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToFloat32E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && float32(math.Abs(float64(got-test.expected))) > 1e-6 {
			t.Errorf("ToFloat32E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
