package mconv_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func TestToComplex128(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected complex128
	}{
		{complex(1, 2), complex(1, 2)},
		{complex64(complex(3, 4)), complex(3, 4)},
		{1, complex(1, 0)},
		{1.5, complex(1.5, 0)},
		{"1+2i", complex(1, 2)},
		{true, complex(1, 0)},
		{false, complex(0, 0)},
		{nil, complex(0, 0)},
	}

	for _, test := range tests {
		if got := mconv.ToComplex128(test.input); got != test.expected {
			t.Errorf("mconv.ToComplex128(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToComplex128E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected complex128
		isErr    bool
	}{
		{complex(1, 2), complex(1, 2), false},
		{complex64(complex(3, 4)), complex(3, 4), false},
		{1, complex(1, 0), false},
		{1.5, complex(1.5, 0), false},
		{"1+2i", complex(1, 2), false},
		{"invalid", complex(0, 0), true},
		{true, complex(1, 0), false},
		{false, complex(0, 0), false},
		{nil, complex(0, 0), false},
	}

	for _, test := range tests {
		got, err := mconv.ToComplex128E(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToComplex128E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToComplex128E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("mconv.ToComplex128E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToComplex64(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected complex64
	}{
		{complex(1, 2), complex64(complex(1, 2))},
		{complex64(complex(3, 4)), complex64(complex(3, 4))},
		{1, complex64(complex(1, 0))},
		{1.5, complex64(complex(1.5, 0))},
		{"1+2i", complex64(complex(1, 2))},
		{true, complex64(complex(1, 0))},
		{false, complex64(complex(0, 0))},
		{nil, complex64(complex(0, 0))},
	}

	for _, test := range tests {
		if got := mconv.ToComplex64(test.input); got != test.expected {
			t.Errorf("ToComplex64(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToComplex64E(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected complex64
		isErr    bool
	}{
		{complex(1, 2), complex64(complex(1, 2)), false},
		{complex64(complex(3, 4)), complex64(complex(3, 4)), false},
		{1, complex64(complex(1, 0)), false},
		{1.5, complex64(complex(1.5, 0)), false},
		{"1+2i", complex64(complex(1, 2)), false},
		{"invalid", complex64(complex(0, 0)), true},
		{true, complex64(complex(1, 0)), false},
		{false, complex64(complex(0, 0)), false},
		{nil, complex64(complex(0, 0)), false},
	}

	for _, test := range tests {
		got, err := mconv.ToComplex64E(test.input)
		if test.isErr && err == nil {
			t.Errorf("ToComplex64E(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("ToComplex64E(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && got != test.expected {
			t.Errorf("ToComplex64E(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
