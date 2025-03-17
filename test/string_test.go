package mconv_test

import (
	"testing"

	"github.com/mingzaily/mconv"
)

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{123, "123"},
		{123.45, "123.45"},
		{true, "true"},
		{[]byte("test"), "test"},
		{nil, ""},
	}

	for _, test := range tests {
		if got := mconv.ToString(test.input); got != test.expected {
			t.Errorf("mconv.ToString(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToStringE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		isErr    bool
	}{
		{123, "123", false},
		{123.45, "123.45", false},
		{true, "true", false},
		{[]byte("test"), "test", false},
		{nil, "", false},
	}

	for _, test := range tests {
		got, err := mconv.ToStringE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToStringE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToStringE(%v) unexpected error: %v", test.input, err)
		}
		if got != test.expected {
			t.Errorf("mconv.ToStringE(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
