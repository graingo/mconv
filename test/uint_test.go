package mconv_test

import (
	"math"
	"testing"

	"github.com/graingo/mconv"
)

func TestToUint(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint
		isError  bool
	}{
		{nil, 0, false},
		{uint(123), uint(123), false},
		{uint64(123), uint(123), false},
		{uint32(123), uint(123), false},
		{uint16(123), uint(123), false},
		{uint8(123), uint(123), false},
		{int(123), uint(123), false},
		{int64(123), uint(123), false},
		{int32(123), uint(123), false},
		{int16(123), uint(123), false},
		{int8(123), uint(123), false},
		{float64(123.45), uint(123), false},
		{float32(123.45), uint(123), false},
		{complex(123.45, 0), uint(123), false},
		{complex64(complex(123.45, 0)), uint(123), false},
		{true, uint(1), false},
		{false, uint(0), false},
		{"123", uint(123), false},
		{"0x7B", uint(123), false},
		{"0173", uint(123), false},
		{"123", uint(123), false},

		// Error cases
		{int(-123), uint(0), true},
		{int64(-123), uint(0), true},
		{float64(-123.45), uint(0), true},
		{complex(123.45, 1), uint(0), true},
		{"abc", uint(0), true},
		{struct{}{}, uint(0), true},
	}

	for _, test := range tests {
		result, err := mconv.ToUintE(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("mconv.ToUintE(%v) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("mconv.ToUintE(%v) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("mconv.ToUintE(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}

		// Test non-error version
		if !test.isError {
			result := mconv.ToUint(test.input)
			if result != test.expected {
				t.Errorf("mconv.ToUint(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToUint64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint64
		isError  bool
	}{
		{nil, 0, false},
		{uint(123), uint64(123), false},
		{uint64(123), uint64(123), false},
		{uint32(123), uint64(123), false},
		{uint16(123), uint64(123), false},
		{uint8(123), uint64(123), false},
		{int(123), uint64(123), false},
		{int64(123), uint64(123), false},
		{int32(123), uint64(123), false},
		{int16(123), uint64(123), false},
		{int8(123), uint64(123), false},
		{float64(123.45), uint64(123), false},
		{float32(123.45), uint64(123), false},
		{complex(123.45, 0), uint64(123), false},
		{complex64(complex(123.45, 0)), uint64(123), false},
		{true, uint64(1), false},
		{false, uint64(0), false},
		{"123", uint64(123), false},
		{"0x7B", uint64(123), false},
		{"0173", uint64(123), false},
		{"123", uint64(123), false},

		// Error cases
		{int(-123), uint64(0), true},
		{int64(-123), uint64(0), true},
		{float64(-123.45), uint64(0), true},
		{complex(123.45, 1), uint64(0), true},
		{"abc", uint64(0), true},
		{struct{}{}, uint64(0), true},
	}

	for _, test := range tests {
		result, err := mconv.ToUint64E(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("ToUint64E(%v) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToUint64E(%v) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToUint64E(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}

		// Test non-error version
		if !test.isError {
			result := mconv.ToUint64(test.input)
			if result != test.expected {
				t.Errorf("ToUint64(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToUint32(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint32
		isError  bool
	}{
		{nil, 0, false},
		{uint(123), uint32(123), false},
		{uint64(123), uint32(123), false},
		{uint32(123), uint32(123), false},
		{uint16(123), uint32(123), false},
		{uint8(123), uint32(123), false},
		{int(123), uint32(123), false},
		{int64(123), uint32(123), false},
		{int32(123), uint32(123), false},
		{int16(123), uint32(123), false},
		{int8(123), uint32(123), false},
		{float64(123.45), uint32(123), false},
		{float32(123.45), uint32(123), false},
		{complex(123.45, 0), uint32(123), false},
		{complex64(complex(123.45, 0)), uint32(123), false},
		{true, uint32(1), false},
		{false, uint32(0), false},
		{"123", uint32(123), false},
		{"0x7B", uint32(123), false},
		{"0173", uint32(123), false},
		{"123", uint32(123), false},

		// Error cases
		{uint64(math.MaxUint32 + 1), uint32(0), true},
		{int(-123), uint32(0), true},
		{int64(-123), uint32(0), true},
		{float64(-123.45), uint32(0), true},
		{complex(123.45, 1), uint32(0), true},
		{"abc", uint32(0), true},
		{struct{}{}, uint32(0), true},
	}

	for _, test := range tests {
		result, err := mconv.ToUint32E(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("ToUint32E(%v) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToUint32E(%v) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToUint32E(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}

		// Test non-error version
		if !test.isError {
			result := mconv.ToUint32(test.input)
			if result != test.expected {
				t.Errorf("ToUint32(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToUint16(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint16
		isError  bool
	}{
		{nil, 0, false},
		{uint(123), uint16(123), false},
		{uint64(123), uint16(123), false},
		{uint32(123), uint16(123), false},
		{uint16(123), uint16(123), false},
		{uint8(123), uint16(123), false},
		{int(123), uint16(123), false},
		{int64(123), uint16(123), false},
		{int32(123), uint16(123), false},
		{int16(123), uint16(123), false},
		{int8(123), uint16(123), false},
		{float64(123.45), uint16(123), false},
		{float32(123.45), uint16(123), false},
		{complex(123.45, 0), uint16(123), false},
		{complex64(complex(123.45, 0)), uint16(123), false},
		{true, uint16(1), false},
		{false, uint16(0), false},
		{"123", uint16(123), false},
		{"0x7B", uint16(123), false},
		{"0173", uint16(123), false},
		{"123", uint16(123), false},

		// Error cases
		{uint64(math.MaxUint16 + 1), uint16(0), true},
		{int(-123), uint16(0), true},
		{int64(-123), uint16(0), true},
		{float64(-123.45), uint16(0), true},
		{complex(123.45, 1), uint16(0), true},
		{"abc", uint16(0), true},
		{struct{}{}, uint16(0), true},
	}

	for _, test := range tests {
		result, err := mconv.ToUint16E(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("ToUint16E(%v) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToUint16E(%v) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToUint16E(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}

		// Test non-error version
		if !test.isError {
			result := mconv.ToUint16(test.input)
			if result != test.expected {
				t.Errorf("ToUint16(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}

func TestToUint8(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint8
		isError  bool
	}{
		{nil, 0, false},
		{uint(123), uint8(123), false},
		{uint64(123), uint8(123), false},
		{uint32(123), uint8(123), false},
		{uint16(123), uint8(123), false},
		{uint8(123), uint8(123), false},
		{int(123), uint8(123), false},
		{int64(123), uint8(123), false},
		{int32(123), uint8(123), false},
		{int16(123), uint8(123), false},
		{int8(123), uint8(123), false},
		{float64(123.45), uint8(123), false},
		{float32(123.45), uint8(123), false},
		{complex(123.45, 0), uint8(123), false},
		{complex64(complex(123.45, 0)), uint8(123), false},
		{true, uint8(1), false},
		{false, uint8(0), false},
		{"123", uint8(123), false},
		{"0x7B", uint8(123), false},
		{"0173", uint8(123), false},
		{"123", uint8(123), false},

		// Error cases
		{uint64(math.MaxUint8 + 1), uint8(0), true},
		{int(-123), uint8(0), true},
		{int64(-123), uint8(0), true},
		{float64(-123.45), uint8(0), true},
		{complex(123.45, 1), uint8(0), true},
		{"abc", uint8(0), true},
		{struct{}{}, uint8(0), true},
	}

	for _, test := range tests {
		result, err := mconv.ToUint8E(test.input)
		if test.isError {
			if err == nil {
				t.Errorf("ToUint8E(%v) should return error", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("ToUint8E(%v) returned error: %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("ToUint8E(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}

		// Test non-error version
		if !test.isError {
			result := mconv.ToUint8(test.input)
			if result != test.expected {
				t.Errorf("ToUint8(%v) = %v, expected %v", test.input, result, test.expected)
			}
		}
	}
}
