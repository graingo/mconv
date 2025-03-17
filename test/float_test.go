package mconv_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func TestToFloat64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
	}{
		{123, 123.0},
		{123.45, 123.45},
		{"123.45", 123.45},
		{true, 1.0},
		{false, 0.0},
		{nil, 0.0},
	}

	for _, test := range tests {
		if got := mconv.ToFloat64(test.input); got != test.expected {
			t.Errorf("mconv.ToFloat64(%v) = %v; want %v", test.input, got, test.expected)
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
		{123.45, 123.45, false},
		{"123.45", 123.45, false},
		{"abc", 0.0, true},
		{true, 1.0, false},
		{false, 0.0, false},
		{nil, 0.0, false},
	}

	for _, test := range tests {
		got, err := mconv.ToFloat64E(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToFloat64E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToFloat64E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("mconv.ToFloat64E(%v) = %v; want %v", test.input, got, test.expected)
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
		{123.45, 123.45, false},
		{"123.45", 123.45, false},
		{"abc", 0.0, true},
		{true, 1.0, false},
		{false, 0.0, false},
		{nil, 0.0, false},
	}

	for _, test := range tests {
		got, err := mconv.ToFloat32E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToFloat32E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToFloat32E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("ToFloat32E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
