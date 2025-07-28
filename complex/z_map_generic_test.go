package complex_test

import (
	"reflect"
	"testing"

	"github.com/graingo/mconv/complex"
)

func TestToMapT(t *testing.T) {
	t.Run("string to int", func(t *testing.T) {
		source := map[string]interface{}{"a": "1", "b": 2.0}
		expected := map[string]int{"a": 1, "b": 2}
		result := complex.ToMapT[string, int](source)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("int to float64", func(t *testing.T) {
		source := map[interface{}]interface{}{1: "1.1", "2": 2}
		expected := map[int]float64{1: 1.1, 2: 2.0}
		result := complex.ToMapT[int, float64](source)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		var result map[string]string = complex.ToMapT[string, string](nil)
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}

func TestToMapTE(t *testing.T) {
	t.Run("successful conversion", func(t *testing.T) {
		source := map[string]interface{}{"a": "1", "b": true}
		expected := map[string]string{"a": "1", "b": "true"}
		result, err := complex.ToMapTE[string, string](source)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("key conversion error", func(t *testing.T) {
		// complex key cannot be converted to int
		source := map[interface{}]interface{}{1 + 2i: "value"}
		_, err := complex.ToMapTE[int, string](source)
		if err == nil {
			t.Error("expected an error but got nil")
		}
	})

	t.Run("value conversion error", func(t *testing.T) {
		source := map[string]interface{}{"a": "not-an-int"}
		_, err := complex.ToMapTE[string, int](source)
		if err == nil {
			t.Error("expected an error but got nil")
		}
	})

	t.Run("unsupported source", func(t *testing.T) {
		source := "a string"
		_, err := complex.ToMapTE[string, string](source)
		if err == nil {
			t.Error("expected an error but got nil")
		}
	})
}
