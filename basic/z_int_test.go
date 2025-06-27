package basic_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func TestToInt(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
	}{
		{123, 123},
		{123.45, 123},
		{"123", 123},
		{true, 1},
		{false, 0},
		{nil, 0},
	}

	for _, test := range tests {
		if got := mconv.ToInt(test.input); got != test.expected {
			t.Errorf("mconv.ToInt(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToIntE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
		isErr    bool
	}{
		{123, 123, false},
		{123.45, 123, false},
		{"123", 123, false},
		{"abc", 0, true},
		{true, 1, false},
		{false, 0, false},
		{nil, 0, false},
	}

	for _, test := range tests {
		got, err := mconv.ToIntE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToIntE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToIntE(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("mconv.ToIntE(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToInt64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int64
	}{
		{int64(123), int64(123)},
		{123, int64(123)},
		{123.45, int64(123)},
		{"123", int64(123)},
		{true, int64(1)},
		{false, int64(0)},
		{nil, int64(0)},
	}

	for _, test := range tests {
		if got := mconv.ToInt64(test.input); got != test.expected {
			t.Errorf("ToInt64(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToInt64E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int64
		isErr    bool
	}{
		{int64(123), int64(123), false},
		{123, int64(123), false},
		{123.45, int64(123), false},
		{"123", int64(123), false},
		{"abc", int64(0), true},
		{true, int64(1), false},
		{false, int64(0), false},
		{nil, int64(0), false},
	}

	for _, test := range tests {
		got, err := mconv.ToInt64E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToInt64E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToInt64E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("ToInt64E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToInt32E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int32
		isErr    bool
	}{
		{int32(123), int32(123), false},
		{123, int32(123), false},
		{123.45, int32(123), false},
		{"123", int32(123), false},
		{"abc", int32(0), true},
		{true, int32(1), false},
		{false, int32(0), false},
		{nil, int32(0), false},
	}

	for _, test := range tests {
		got, err := mconv.ToInt32E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToInt32E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToInt32E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("ToInt32E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToInt16E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int16
		isErr    bool
	}{
		{int16(123), int16(123), false},
		{123, int16(123), false},
		{123.45, int16(123), false},
		{"123", int16(123), false},
		{"abc", int16(0), true},
		{true, int16(1), false},
		{false, int16(0), false},
		{nil, int16(0), false},
	}

	for _, test := range tests {
		got, err := mconv.ToInt16E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToInt16E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToInt16E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("ToInt16E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToInt8E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int8
		isErr    bool
	}{
		{int8(123), int8(123), false},
		{123, int8(123), false},
		{123.45, int8(123), false},
		{"123", int8(123), false},
		{"abc", int8(0), true},
		{true, int8(1), false},
		{false, int8(0), false},
		{nil, int8(0), false},
	}

	for _, test := range tests {
		got, err := mconv.ToInt8E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToInt8E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToInt8E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("ToInt8E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
