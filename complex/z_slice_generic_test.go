package complex_test

import (
	"reflect"
	"testing"

	"github.com/graingo/mconv/complex"
)

func TestToSliceT(t *testing.T) {
	t.Run("to string slice", func(t *testing.T) {
		source := []interface{}{"a", 1, true}
		expected := []string{"a", "1", "true"}
		result := complex.ToSliceT[string](source)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("to int slice", func(t *testing.T) {
		source := []string{"1", "2", "3"}
		expected := []int{1, 2, 3}
		result := complex.ToSliceT[int](source)
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		var result []string = complex.ToSliceT[string](nil)
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})
}

func TestToSliceTE(t *testing.T) {
	t.Run("successful conversion", func(t *testing.T) {
		source := []interface{}{1, 2.5, "3"}
		expected := []float64{1.0, 2.5, 3.0}
		result, err := complex.ToSliceTE[float64](source)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("conversion error", func(t *testing.T) {
		source := []interface{}{"a", "b", "c"}
		_, err := complex.ToSliceTE[int](source)
		if err == nil {
			t.Error("expected an error but got nil")
		}
	})

	t.Run("unsupported type", func(t *testing.T) {
		source := map[string]int{"a": 1}
		_, err := complex.ToSliceTE[int](source)
		if err == nil {
			t.Error("expected an error but got nil")
		}
	})
}
