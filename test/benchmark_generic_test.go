package mconv_test

import (
	"reflect"
	"testing"

	"github.com/mingzaily/mconv"
	"github.com/mingzaily/mconv/complex"
	"github.com/mingzaily/mconv/internal"
)

// Complex struct for testing reflection performance
type ComplexStruct struct {
	Name         string
	Age          int
	Score        float64
	IsActive     bool
	Tags         []string
	Scores       map[string]float64
	NestedStruct struct {
		Field1 string
		Field2 int
	}
}

// Generic slice conversion benchmarks
func BenchmarkToSliceT(b *testing.B) {
	slice := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		complex.ToSliceT[string](slice)
	}
}

// Generic Map conversion benchmarks
func BenchmarkToMapT(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		complex.ToMapT[string, string](m)
	}
}

// Direct reflection cache benchmarks
func BenchmarkReflectionCache_TypeInfo(b *testing.B) {
	// Create a complex struct for testing
	testStruct := ComplexStruct{
		Name:     "Test User",
		Age:      30,
		Score:    95.5,
		IsActive: true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Scores:   map[string]float64{"math": 95.5, "science": 92.0, "history": 88.5},
	}
	testStruct.NestedStruct.Field1 = "Nested Field"
	testStruct.NestedStruct.Field2 = 42

	// First run without cache
	b.Run("WithoutCache", func(b *testing.B) {
		// Clear cache before testing
		mconv.ClearTypeInfoCache()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Get type info directly using reflect
			t := reflect.TypeOf(testStruct)
			kind := t.Kind()
			numField := 0
			if kind == reflect.Struct {
				numField = t.NumField()
				for j := 0; j < numField; j++ {
					field := t.Field(j)
					_ = field.Name
					_ = field.Type
				}
			}
		}
	})

	// Run with cache
	b.Run("WithCache", func(b *testing.B) {
		// Warm up cache
		typeInfo := internal.GetTypeInfo(reflect.TypeOf(testStruct))

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			// Get type info using cache
			_ = internal.GetTypeInfo(reflect.TypeOf(testStruct))
			for _, fieldName := range typeInfo.FieldNames {
				field := typeInfo.Fields[fieldName]
				_ = field.Name
				_ = field.Type
			}
		}
	})
}

// Complex struct conversion benchmarks
func BenchmarkComplexStructConversion(b *testing.B) {
	// Create a complex struct for testing
	testStruct := ComplexStruct{
		Name:     "Test User",
		Age:      30,
		Score:    95.5,
		IsActive: true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Scores:   map[string]float64{"math": 95.5, "science": 92.0, "history": 88.5},
	}
	testStruct.NestedStruct.Field1 = "Nested Field"
	testStruct.NestedStruct.Field2 = 42

	// First run without cache
	b.Run("WithoutCache", func(b *testing.B) {
		// Clear cache before testing
		mconv.ClearTypeInfoCache()
		mconv.ClearConversionCache()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mconv.ToMap(testStruct)
		}
	})

	// Run with cache
	b.Run("WithCache", func(b *testing.B) {
		// Warm up cache
		mconv.ToMap(testStruct)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			mconv.ToMap(testStruct)
		}
	})
}

// Large slice conversion benchmarks
func BenchmarkLargeSliceConversion(b *testing.B) {
	// Create a large slice for testing
	largeIntSlice := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		largeIntSlice[i] = i
	}

	// First run without cache
	b.Run("WithoutCache", func(b *testing.B) {
		// Clear cache before testing
		mconv.ClearTypeInfoCache()
		mconv.ClearConversionCache()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			complex.ToSliceT[string](largeIntSlice)
		}
	})

	// Run with cache
	b.Run("WithCache", func(b *testing.B) {
		// Warm up cache
		complex.ToSliceT[string](largeIntSlice)

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			complex.ToSliceT[string](largeIntSlice)
		}
	})
}
