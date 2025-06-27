package mconv_test

import (
	"testing"

	"github.com/graingo/mconv"
	"github.com/graingo/mconv/complex"
)

// --- Test Basic for Benchmarking ---

func BenchmarkToString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mconv.ToString(123)
	}
}

func BenchmarkToStringE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToStringE(123)
	}
}

func BenchmarkToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mconv.ToInt("123")
	}
}

func BenchmarkToIntE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToIntE("123")
	}
}

func BenchmarkToFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mconv.ToFloat64("123.45")
	}
}

func BenchmarkToFloat64E(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToFloat64E("123.45")
	}
}

func BenchmarkToBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mconv.ToBool("true")
	}
}

func BenchmarkToBoolE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToBoolE("true")
	}
}

func BenchmarkToTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mconv.ToTime("2006-01-02T15:04:05Z")
	}
}

func BenchmarkToTimeE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToTimeE("2006-01-02T15:04:05Z")
	}
}

// Complex type conversion benchmarks
func BenchmarkToSlice(b *testing.B) {
	arr := [3]int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		mconv.ToSlice(arr)
	}
}

func BenchmarkToSliceE(b *testing.B) {
	arr := [3]int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToSliceE(arr)
	}
}

func BenchmarkToStringSlice(b *testing.B) {
	slice := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		mconv.ToStringSlice(slice)
	}
}

func BenchmarkToStringSliceE(b *testing.B) {
	slice := []int{1, 2, 3}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToStringSliceE(slice)
	}
}

func BenchmarkToIntSlice(b *testing.B) {
	slice := []string{"1", "2", "3"}
	for i := 0; i < b.N; i++ {
		mconv.ToIntSlice(slice)
	}
}

func BenchmarkToIntSliceE(b *testing.B) {
	slice := []string{"1", "2", "3"}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToIntSliceE(slice)
	}
}

func BenchmarkToFloat64Slice(b *testing.B) {
	slice := []string{"1.1", "2.2", "3.3"}
	for i := 0; i < b.N; i++ {
		mconv.ToFloat64Slice(slice)
	}
}

func BenchmarkToFloat64SliceE(b *testing.B) {
	slice := []string{"1.1", "2.2", "3.3"}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToFloat64SliceE(slice)
	}
}

func BenchmarkToMap(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		mconv.ToMap(m)
	}
}

func BenchmarkToMapE(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToMapE(m)
	}
}

func BenchmarkToStringMap(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		mconv.ToStringMap(m)
	}
}

func BenchmarkToStringMapE(b *testing.B) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToStringMapE(m)
	}
}

func BenchmarkToIntMap(b *testing.B) {
	m := map[string]string{"a": "1", "b": "2", "c": "3"}
	for i := 0; i < b.N; i++ {
		mconv.ToIntMap(m)
	}
}

func BenchmarkToIntMapE(b *testing.B) {
	m := map[string]string{"a": "1", "b": "2", "c": "3"}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToIntMapE(m)
	}
}

func BenchmarkToFloat64Map(b *testing.B) {
	m := map[string]string{"a": "1.1", "b": "2.2", "c": "3.3"}
	for i := 0; i < b.N; i++ {
		mconv.ToFloat64Map(m)
	}
}

func BenchmarkToFloat64MapE(b *testing.B) {
	m := map[string]string{"a": "1.1", "b": "2.2", "c": "3.3"}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToFloat64MapE(m)
	}
}

// JSON conversion benchmarks
func BenchmarkToJSON(b *testing.B) {
	m := map[string]interface{}{"name": "John", "age": 30}
	for i := 0; i < b.N; i++ {
		mconv.ToJSON(m)
	}
}

func BenchmarkToJSONE(b *testing.B) {
	m := map[string]interface{}{"name": "John", "age": 30}
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToJSONE(m)
	}
}

func BenchmarkFromJSON(b *testing.B) {
	json := `{"name":"John","age":30}`
	var result map[string]interface{}
	for i := 0; i < b.N; i++ {
		mconv.FromJSON(json, &result)
	}
}

func BenchmarkFromJSONE(b *testing.B) {
	json := `{"name":"John","age":30}`
	var result map[string]interface{}
	for i := 0; i < b.N; i++ {
		_ = mconv.FromJSONE(json, &result)
	}
}

func BenchmarkToMapFromJSON(b *testing.B) {
	json := `{"name":"John","age":30}`
	for i := 0; i < b.N; i++ {
		mconv.ToMapFromJSON(json)
	}
}

func BenchmarkToMapFromJSONE(b *testing.B) {
	json := `{"name":"John","age":30}`
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToMapFromJSONE(json)
	}
}

func BenchmarkToSliceFromJSON(b *testing.B) {
	json := `["John","Jane"]`
	for i := 0; i < b.N; i++ {
		mconv.ToSliceFromJSON(json)
	}
}

func BenchmarkToSliceFromJSONE(b *testing.B) {
	json := `["John","Jane"]`
	for i := 0; i < b.N; i++ {
		_, _ = mconv.ToSliceFromJSONE(json)
	}
}

// --- Test Structs for Benchmarking ---

func BenchmarkStructConversion(b *testing.B) {
	source := map[string]interface{}{"user_id": 2, "user_name": "Bob"}
	var target UserWithTags

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = complex.StructE(source, &target)
	}
}

func BenchmarkStructConversionParallel(b *testing.B) {
	source := map[string]interface{}{"user_id": 2, "user_name": "Bob"}

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		var target UserWithTags
		for pb.Next() {
			_ = complex.StructE(source, &target)
		}
	})
}
