package complex_test

import (
	"testing"

	"github.com/graingo/mconv"
)

func TestToMap(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected map[string]interface{}
	}{
		{map[string]interface{}{"a": 1, "b": 2}, map[string]interface{}{"a": 1, "b": 2}},
		{map[string]string{"a": "1", "b": "2"}, map[string]interface{}{"a": "1", "b": "2"}},
		{map[interface{}]interface{}{"a": 1, "b": 2}, map[string]interface{}{"a": 1, "b": 2}},
		{nil, nil},
	}

	for _, test := range tests {
		got := mconv.ToMap(test.input)
		if len(got) != len(test.expected) {
			t.Errorf("mconv.ToMap(%v) length = %v; want %v", test.input, len(got), len(test.expected))
			continue
		}
		for k, v := range test.expected {
			if got[k] != v {
				t.Errorf("mconv.ToMap(%v)[%s] = %v; want %v", test.input, k, got[k], v)
			}
		}
	}
}

func TestToMapE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected map[string]interface{}
		isErr    bool
	}{
		{map[string]interface{}{"a": 1, "b": 2}, map[string]interface{}{"a": 1, "b": 2}, false},
		{map[string]string{"a": "1", "b": "2"}, map[string]interface{}{"a": "1", "b": "2"}, false},
		{map[interface{}]interface{}{"a": 1, "b": 2}, map[string]interface{}{"a": 1, "b": 2}, false},
		{123, nil, true},
		{nil, nil, false},
	}

	for _, test := range tests {
		got, err := mconv.ToMapE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToMapE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToMapE(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr {
			if len(got) != len(test.expected) {
				t.Errorf("mconv.ToMapE(%v) length = %v; want %v", test.input, len(got), len(test.expected))
				continue
			}
			for k, v := range test.expected {
				if got[k] != v {
					t.Errorf("mconv.ToMapE(%v)[%s] = %v; want %v", test.input, k, got[k], v)
				}
			}
		}
	}
}

func TestToStringMap(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected map[string]string
	}{
		{map[string]string{"a": "1", "b": "2"}, map[string]string{"a": "1", "b": "2"}},
		{map[string]interface{}{"a": 1, "b": 2}, map[string]string{"a": "1", "b": "2"}},
		{map[interface{}]interface{}{"a": 1, "b": 2}, map[string]string{"a": "1", "b": "2"}},
		{nil, nil},
	}

	for _, test := range tests {
		got := mconv.ToStringMap(test.input)
		if len(got) != len(test.expected) {
			t.Errorf("ToStringMap(%v) length = %v; want %v", test.input, len(got), len(test.expected))
			continue
		}
		for k, v := range test.expected {
			if got[k] != v {
				t.Errorf("ToStringMap(%v)[%s] = %v; want %v", test.input, k, got[k], v)
			}
		}
	}
}

func TestToIntMap(t *testing.T) {
	// Test nil value
	result := mconv.ToIntMap(nil)
	if len(result) != 0 {
		t.Errorf("ToIntMap(nil) = %v; want %v", result, map[string]int{})
	}

	// Test integer map
	result = mconv.ToIntMap(map[string]int{"a": 1, "b": 2, "c": 3})
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !intMapEqual(result, expected) {
		t.Errorf("ToIntMap(map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}) = %v; want %v", result, expected)
	}

	// Test string map
	result = mconv.ToIntMap(map[string]string{"a": "1", "b": "2", "c": "3"})
	expected = map[string]int{"a": 1, "b": 2, "c": 3}
	if !intMapEqual(result, expected) {
		t.Errorf("ToIntMap(map[string]string{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}) = %v; want %v", result, expected)
	}

	// Test interface map
	result = mconv.ToIntMap(map[string]interface{}{"a": 1, "b": "2", "c": 3.0})
	expected = map[string]int{"a": 1, "b": 2, "c": 3}
	if !intMapEqual(result, expected) {
		t.Errorf("ToIntMap(map[string]interface{}{\"a\": 1, \"b\": \"2\", \"c\": 3.0}) = %v; want %v", result, expected)
	}

	// Test other types of map
	result = mconv.ToIntMap(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3})
	expected = map[string]int{"a": 1, "b": 2, "c": 3}
	if !intMapEqual(result, expected) {
		t.Errorf("ToIntMap(map[string]float64{\"a\": 1.1, \"b\": 2.2, \"c\": 3.3}) = %v; want %v", result, expected)
	}
}

func TestToIntMapE(t *testing.T) {
	// Test nil value
	result, err := mconv.ToIntMapE(nil)
	if err != nil {
		t.Errorf("ToIntMapE(nil) returned error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("ToIntMapE(nil) = %v; want %v", result, map[string]int{})
	}

	// Test integer map
	result, err = mconv.ToIntMapE(map[string]int{"a": 1, "b": 2, "c": 3})
	if err != nil {
		t.Errorf("ToIntMapE(map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}) returned error: %v", err)
	}
	expected := map[string]int{"a": 1, "b": 2, "c": 3}
	if !intMapEqual(result, expected) {
		t.Errorf("ToIntMapE(map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}) = %v; want %v", result, expected)
	}

	// Test string map
	result, err = mconv.ToIntMapE(map[string]string{"a": "1", "b": "2", "c": "3"})
	if err != nil {
		t.Errorf("ToIntMapE(map[string]string{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}) returned error: %v", err)
	}
	expected = map[string]int{"a": 1, "b": 2, "c": 3}
	if !intMapEqual(result, expected) {
		t.Errorf("ToIntMapE(map[string]string{\"a\": \"1\", \"b\": \"2\", \"c\": \"3\"}) = %v; want %v", result, expected)
	}

	// Test error cases
	_, err = mconv.ToIntMapE(map[string]string{"a": "1", "b": "abc", "c": "3"})
	if err == nil {
		t.Errorf("ToIntMapE(map[string]string{\"a\": \"1\", \"b\": \"abc\", \"c\": \"3\"}) did not return error")
	}
}

func TestToFloat64Map(t *testing.T) {
	// Test nil value
	result := mconv.ToFloat64Map(nil)
	if len(result) != 0 {
		t.Errorf("ToFloat64Map(nil) = %v; want %v", result, map[string]float64{})
	}

	// Test float map
	result = mconv.ToFloat64Map(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3})
	expected := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	if !float64MapEqual(result, expected) {
		t.Errorf("ToFloat64Map(map[string]float64{\"a\": 1.1, \"b\": 2.2, \"c\": 3.3}) = %v; want %v", result, expected)
	}

	// Test string map
	result = mconv.ToFloat64Map(map[string]string{"a": "1.1", "b": "2.2", "c": "3.3"})
	expected = map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	if !float64MapEqual(result, expected) {
		t.Errorf("ToFloat64Map(map[string]string{\"a\": \"1.1\", \"b\": \"2.2\", \"c\": \"3.3\"}) = %v; want %v", result, expected)
	}

	// Test interface map
	result = mconv.ToFloat64Map(map[string]interface{}{"a": 1.1, "b": "2.2", "c": 3})
	expected = map[string]float64{"a": 1.1, "b": 2.2, "c": 3.0}
	if !float64MapEqual(result, expected) {
		t.Errorf("ToFloat64Map(map[string]interface{}{\"a\": 1.1, \"b\": \"2.2\", \"c\": 3}) = %v; want %v", result, expected)
	}

	// Test integer map
	result = mconv.ToFloat64Map(map[string]int{"a": 1, "b": 2, "c": 3})
	expected = map[string]float64{"a": 1.0, "b": 2.0, "c": 3.0}
	if !float64MapEqual(result, expected) {
		t.Errorf("ToFloat64Map(map[string]int{\"a\": 1, \"b\": 2, \"c\": 3}) = %v; want %v", result, expected)
	}
}

func TestToFloat64MapE(t *testing.T) {
	// Test nil value
	result, err := mconv.ToFloat64MapE(nil)
	if err != nil {
		t.Errorf("ToFloat64MapE(nil) returned error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("ToFloat64MapE(nil) = %v; want %v", result, map[string]float64{})
	}

	// Test float map
	result, err = mconv.ToFloat64MapE(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3})
	if err != nil {
		t.Errorf("ToFloat64MapE(map[string]float64{\"a\": 1.1, \"b\": 2.2, \"c\": 3.3}) returned error: %v", err)
	}
	expected := map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	if !float64MapEqual(result, expected) {
		t.Errorf("ToFloat64MapE(map[string]float64{\"a\": 1.1, \"b\": 2.2, \"c\": 3.3}) = %v; want %v", result, expected)
	}

	// Test string map
	result, err = mconv.ToFloat64MapE(map[string]string{"a": "1.1", "b": "2.2", "c": "3.3"})
	if err != nil {
		t.Errorf("ToFloat64MapE(map[string]string{\"a\": \"1.1\", \"b\": \"2.2\", \"c\": \"3.3\"}) returned error: %v", err)
	}
	expected = map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}
	if !float64MapEqual(result, expected) {
		t.Errorf("ToFloat64MapE(map[string]string{\"a\": \"1.1\", \"b\": \"2.2\", \"c\": \"3.3\"}) = %v; want %v", result, expected)
	}

	// Test error cases
	_, err = mconv.ToFloat64MapE(map[string]string{"a": "1.1", "b": "abc", "c": "3.3"})
	if err == nil {
		t.Errorf("ToFloat64MapE(map[string]string{\"a\": \"1.1\", \"b\": \"abc\", \"c\": \"3.3\"}) did not return error")
	}
}

// Helper function: compare if integer maps are equal
func intMapEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}

// Helper function: compare if float maps are equal
func float64MapEqual(a, b map[string]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if v != b[k] {
			return false
		}
	}
	return true
}
