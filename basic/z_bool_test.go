package basic_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
	}{
		{true, true},
		{false, false},
		{1, true},
		{0, false},
		{"true", true},
		{"false", false},
		{nil, false},
	}

	for _, test := range tests {
		if got := mconv.ToBool(test.input); got != test.expected {
			t.Errorf("mconv.ToBool(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToBoolE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
		isErr    bool
	}{
		{true, true, false},
		{false, false, false},
		{1, true, false},
		{0, false, false},
		{"true", true, false},
		{"false", false, false},
		{"yes", true, false},
		{"no", false, false},
		{"invalid", false, true},
		{nil, false, false},
	}

	for _, test := range tests {
		got, err := mconv.ToBoolE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToBoolE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToBoolE(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("mconv.ToBoolE(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
