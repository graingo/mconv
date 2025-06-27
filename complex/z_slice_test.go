package complex_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func TestToSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []interface{}
	}{
		{[]int{1, 2, 3}, []interface{}{1, 2, 3}},
		{[]string{"a", "b"}, []interface{}{"a", "b"}},
		{123, []interface{}{123}},
		{nil, nil},
	}

	for _, test := range tests {
		got := mconv.ToSlice(test.input)
		if len(got) != len(test.expected) {
			t.Errorf("mconv.ToSlice(%v) length = %v; want %v", test.input, len(got), len(test.expected))
			continue
		}
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("mconv.ToSlice(%v)[%d] = %v; want %v", test.input, i, got[i], test.expected[i])
			}
		}
	}
}

func TestToSliceE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []interface{}
		isErr    bool
	}{
		{[]int{1, 2, 3}, []interface{}{1, 2, 3}, false},
		{[]string{"a", "b"}, []interface{}{"a", "b"}, false},
		{123, []interface{}{123}, false},
		{nil, nil, false},
	}

	for _, test := range tests {
		got, err := mconv.ToSliceE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToSliceE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToSliceE(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr {
			if len(got) != len(test.expected) {
				t.Errorf("mconv.ToSliceE(%v) length = %v; want %v", test.input, len(got), len(test.expected))
				continue
			}
			for i := range got {
				if got[i] != test.expected[i] {
					t.Errorf("mconv.ToSliceE(%v)[%d] = %v; want %v", test.input, i, got[i], test.expected[i])
				}
			}
		}
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []string
	}{
		{[]string{"a", "b"}, []string{"a", "b"}},
		{[]int{1, 2}, []string{"1", "2"}},
		{"a", []string{"a"}},
		{nil, nil},
	}

	for _, test := range tests {
		got := mconv.ToStringSlice(test.input)
		if len(got) != len(test.expected) {
			t.Errorf("ToStringSlice(%v) length = %v; want %v", test.input, len(got), len(test.expected))
			continue
		}
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("ToStringSlice(%v)[%d] = %v; want %v", test.input, i, got[i], test.expected[i])
			}
		}
	}
}

func TestToIntSlice(t *testing.T) {
	// Test nil value
	result := mconv.ToIntSlice(nil)
	if len(result) != 0 {
		t.Errorf("ToIntSlice(nil) = %v; want %v", result, []int{})
	}

	// Test integer slice
	result = mconv.ToIntSlice([]int{1, 2, 3})
	expected := []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSlice([]int{1, 2, 3}) = %v; want %v", result, expected)
	}

	// Test string slice
	result = mconv.ToIntSlice([]string{"1", "2", "3"})
	expected = []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSlice([]string{\"1\", \"2\", \"3\"}) = %v; want %v", result, expected)
	}

	// Test interface slice
	result = mconv.ToIntSlice([]interface{}{1, "2", 3.0})
	expected = []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSlice([]interface{}{1, \"2\", 3.0}) = %v; want %v", result, expected)
	}

	// Test other types of slice
	result = mconv.ToIntSlice([]float64{1.1, 2.2, 3.3})
	expected = []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSlice([]float64{1.1, 2.2, 3.3}) = %v; want %v", result, expected)
	}

	// Test single value
	result = mconv.ToIntSlice(123)
	expected = []int{123}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSlice(123) = %v; want %v", result, expected)
	}
}

func TestToIntSliceE(t *testing.T) {
	// Test nil value
	result, err := mconv.ToIntSliceE(nil)
	if err != nil {
		t.Errorf("ToIntSliceE(nil) returned error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("ToIntSliceE(nil) = %v; want %v", result, []int{})
	}

	// Test integer slice
	result, err = mconv.ToIntSliceE([]int{1, 2, 3})
	if err != nil {
		t.Errorf("ToIntSliceE([]int{1, 2, 3}) returned error: %v", err)
	}
	expected := []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSliceE([]int{1, 2, 3}) = %v; want %v", result, expected)
	}

	// Test string slice
	result, err = mconv.ToIntSliceE([]string{"1", "2", "3"})
	if err != nil {
		t.Errorf("ToIntSliceE([]string{\"1\", \"2\", \"3\"}) returned error: %v", err)
	}
	expected = []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSliceE([]string{\"1\", \"2\", \"3\"}) = %v; want %v", result, expected)
	}

	// Test interface slice
	result, err = mconv.ToIntSliceE([]interface{}{1, "2", 3.0})
	if err != nil {
		t.Errorf("ToIntSliceE([]interface{}{1, \"2\", 3.0}) returned error: %v", err)
	}
	expected = []int{1, 2, 3}
	if !intSliceEqual(result, expected) {
		t.Errorf("ToIntSliceE([]interface{}{1, \"2\", 3.0}) = %v; want %v", result, expected)
	}

	// Test error cases
	_, err = mconv.ToIntSliceE([]string{"1", "abc", "3"})
	if err == nil {
		t.Errorf("ToIntSliceE([]string{\"1\", \"abc\", \"3\"}) did not return error")
	}
}

func TestToFloat64Slice(t *testing.T) {
	// Test nil value
	result := mconv.ToFloat64Slice(nil)
	if len(result) != 0 {
		t.Errorf("ToFloat64Slice(nil) = %v; want %v", result, []float64{})
	}

	// Test float slice
	result = mconv.ToFloat64Slice([]float64{1.1, 2.2, 3.3})
	expected := []float64{1.1, 2.2, 3.3}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64Slice([]float64{1.1, 2.2, 3.3}) = %v; want %v", result, expected)
	}

	// Test string slice
	result = mconv.ToFloat64Slice([]string{"1.1", "2.2", "3.3"})
	expected = []float64{1.1, 2.2, 3.3}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64Slice([]string{\"1.1\", \"2.2\", \"3.3\"}) = %v; want %v", result, expected)
	}

	// Test interface slice
	result = mconv.ToFloat64Slice([]interface{}{1.1, "2.2", 3})
	expected = []float64{1.1, 2.2, 3.0}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64Slice([]interface{}{1.1, \"2.2\", 3}) = %v; want %v", result, expected)
	}

	// Test integer slice
	result = mconv.ToFloat64Slice([]int{1, 2, 3})
	expected = []float64{1.0, 2.0, 3.0}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64Slice([]int{1, 2, 3}) = %v; want %v", result, expected)
	}

	// Test single value
	result = mconv.ToFloat64Slice(123.45)
	expected = []float64{123.45}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64Slice(123.45) = %v; want %v", result, expected)
	}
}

func TestToFloat64SliceE(t *testing.T) {
	// Test nil value
	result, err := mconv.ToFloat64SliceE(nil)
	if err != nil {
		t.Errorf("ToFloat64SliceE(nil) returned error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("ToFloat64SliceE(nil) = %v; want %v", result, []float64{})
	}

	// Test float slice
	result, err = mconv.ToFloat64SliceE([]float64{1.1, 2.2, 3.3})
	if err != nil {
		t.Errorf("ToFloat64SliceE([]float64{1.1, 2.2, 3.3}) returned error: %v", err)
	}
	expected := []float64{1.1, 2.2, 3.3}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64SliceE([]float64{1.1, 2.2, 3.3}) = %v; want %v", result, expected)
	}

	// Test string slice
	result, err = mconv.ToFloat64SliceE([]string{"1.1", "2.2", "3.3"})
	if err != nil {
		t.Errorf("ToFloat64SliceE([]string{\"1.1\", \"2.2\", \"3.3\"}) returned error: %v", err)
	}
	expected = []float64{1.1, 2.2, 3.3}
	if !float64SliceEqual(result, expected) {
		t.Errorf("ToFloat64SliceE([]string{\"1.1\", \"2.2\", \"3.3\"}) = %v; want %v", result, expected)
	}

	// Test error cases
	_, err = mconv.ToFloat64SliceE([]string{"1.1", "abc", "3.3"})
	if err == nil {
		t.Errorf("ToFloat64SliceE([]string{\"1.1\", \"abc\", \"3.3\"}) did not return error")
	}
}

// Helper function: compare if integer slices are equal
func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Helper function: compare if float slices are equal
func float64SliceEqual(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
