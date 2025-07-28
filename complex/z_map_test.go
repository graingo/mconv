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
		{map[interface{}]interface{}{"a": 1, 2: "b"}, map[string]interface{}{"a": 1, "2": "b"}},
		{map[string]string{"a": "1", "b": "2"}, map[string]interface{}{"a": "1", "b": "2"}},
		{map[string]int{"a": 1, "b": 2}, map[string]interface{}{"a": 1, "b": 2}},
		{map[string]int64{"a": 1, "b": 2}, map[string]interface{}{"a": int64(1), "b": int64(2)}},
		{map[string]float32{"a": 1.1, "b": 2.2}, map[string]interface{}{"a": float32(1.1), "b": float32(2.2)}},
		{map[string]float64{"a": 1.1, "b": 2.2}, map[string]interface{}{"a": 1.1, "b": 2.2}},
		{map[string]complex64{"a": 1, "b": 2}, map[string]interface{}{"a": complex64(1), "b": complex64(2)}},
		{map[string]complex128{"a": 1, "b": 2}, map[string]interface{}{"a": complex128(1), "b": complex128(2)}},
		{map[string]bool{"a": true, "b": false}, map[string]interface{}{"a": true, "b": false}},
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
		isErr    bool
	}{
		{map[string]string{"a": "1", "b": "2"}, map[string]string{"a": "1", "b": "2"}, false},
		{map[string]interface{}{"a": 1, "b": 2}, map[string]string{"a": "1", "b": "2"}, false},
		{map[interface{}]interface{}{"a": 1, "b": 2}, map[string]string{"a": "1", "b": "2"}, false},
		{nil, nil, false},
		{"not a map", nil, true},
	}

	for _, test := range tests {
		got, err := mconv.ToStringMapE(test.input)
		if test.isErr {
			if err == nil {
				t.Errorf("ToStringMapE(%v) expected error", test.input)
			}
			continue
		}
		if err != nil {
			t.Errorf("ToStringMapE(%v) unexpected error: %v", test.input, err)
		}

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

	// TODO: Fix this test
	// Test a struct
	// result = mconv.ToIntMap(struct{ A, B int }{1, 2})
	// expected = map[string]int{"A": 1, "B": 2}
	// // The result of converting a struct to a map may have a different order of keys.
	// // So we need to check the length and the values of the keys.
	// if len(result) != len(expected) {
	// 	t.Errorf("ToIntMap(struct) len = %v, want %v", len(result), len(expected))
	// }
	// for k, v := range expected {
	// 	if result[k] != v {
	// 		t.Errorf("ToIntMap(struct) [%s] = %v, want %v", k, result[k], v)
	// 	}
	// }
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

	_, err = mconv.ToIntMapE("not a map")
	if err == nil {
		t.Errorf("ToIntMapE with non-map type should return an error")
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

	// TODO: Fix this test
	// Test a struct
	// result = mconv.ToFloat64Map(struct{ A, B float64 }{1.1, 2.2})
	// expectedF := map[string]float64{"A": 1.1, "B": 2.2}
	// // The result of converting a struct to a map may have a different order of keys.
	// // So we need to check the length and the values of the keys.
	// if len(result) != len(expectedF) {
	// 	t.Errorf("ToFloat64Map(struct) len = %v, want %v", len(result), len(expectedF))
	// }
	// for k, v := range expectedF {
	// 	if result[k] != v {
	// 		t.Errorf("ToFloat64Map(struct) [%s] = %v, want %v", k, result[k], v)
	// 	}
	// }
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

	_, err = mconv.ToFloat64MapE("not a map")
	if err == nil {
		t.Errorf("ToFloat64MapE with non-map type should return an error")
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
