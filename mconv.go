package mconv

import (
	"time"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/complex"
	"github.com/graingo/mconv/internal"
)

// HookFunc is an alias of complex.HookFunc.
type HookFunc = complex.HookFunc

// ToString converts any type to string.
func ToString(value interface{}) string { return basic.ToString(value) }

// ToStringE converts any type to string with error.
func ToStringE(value interface{}) (string, error) { return basic.ToStringE(value) }

// ToInt converts any type to int.
func ToInt(value interface{}) int { return basic.ToInt(value) }

// ToIntE converts any type to int with error.
func ToIntE(value interface{}) (int, error) { return basic.ToIntE(value) }

// ToInt64 converts any type to int64.
func ToInt64(value interface{}) int64 { return basic.ToInt64(value) }

// ToInt64E converts any type to int64 with error.
func ToInt64E(value interface{}) (int64, error) { return basic.ToInt64E(value) }

// ToInt32 converts any type to int32.
func ToInt32(value interface{}) int32 { return basic.ToInt32(value) }

// ToInt32E converts any type to int32 with error.
func ToInt32E(value interface{}) (int32, error) { return basic.ToInt32E(value) }

// ToInt16 converts any type to int16.
func ToInt16(value interface{}) int16 { return basic.ToInt16(value) }

// ToInt16E converts any type to int16 with error.
func ToInt16E(value interface{}) (int16, error) { return basic.ToInt16E(value) }

// ToInt8 converts any type to int8.
func ToInt8(value interface{}) int8 { return basic.ToInt8(value) }

// ToInt8E converts any type to int8 with error.
func ToInt8E(value interface{}) (int8, error) { return basic.ToInt8E(value) }

// ToUint converts any type to uint.
func ToUint(value interface{}) uint { return basic.ToUint(value) }

// ToUintE converts any type to uint with error.
func ToUintE(value interface{}) (uint, error) { return basic.ToUintE(value) }

// ToUint64 converts any type to uint64.
func ToUint64(value interface{}) uint64 { return basic.ToUint64(value) }

// ToUint64E converts any type to uint64 with error.
func ToUint64E(value interface{}) (uint64, error) { return basic.ToUint64E(value) }

// ToUint32 converts any type to uint32.
func ToUint32(value interface{}) uint32 { return basic.ToUint32(value) }

// ToUint32E converts any type to uint32 with error.
func ToUint32E(value interface{}) (uint32, error) { return basic.ToUint32E(value) }

// ToUint16 converts any type to uint16.
func ToUint16(value interface{}) uint16 { return basic.ToUint16(value) }

// ToUint16E converts any type to uint16 with error.
func ToUint16E(value interface{}) (uint16, error) { return basic.ToUint16E(value) }

// ToUint8 converts any type to uint8.
func ToUint8(value interface{}) uint8 { return basic.ToUint8(value) }

// ToUint8E converts any type to uint8 with error.
func ToUint8E(value interface{}) (uint8, error) { return basic.ToUint8E(value) }

// ToFloat64 converts any type to float64.
func ToFloat64(value interface{}) float64 { return basic.ToFloat64(value) }

// ToFloat64E converts any type to float64 with error.
func ToFloat64E(value interface{}) (float64, error) { return basic.ToFloat64E(value) }

// ToFloat32 converts any type to float32.
func ToFloat32(value interface{}) float32 { return basic.ToFloat32(value) }

// ToFloat32E converts any type to float32 with error.
func ToFloat32E(value interface{}) (float32, error) { return basic.ToFloat32E(value) }

// ToBool converts any type to bool.
func ToBool(value interface{}) bool { return basic.ToBool(value) }

// ToBoolE converts any type to bool with error.
func ToBoolE(value interface{}) (bool, error) { return basic.ToBoolE(value) }

// ToComplex128 converts any type to complex128.
func ToComplex128(value interface{}) complex128 { return basic.ToComplex128(value) }

// ToComplex128E converts any type to complex128 with error.
func ToComplex128E(value interface{}) (complex128, error) { return basic.ToComplex128E(value) }

// ToComplex64 converts any type to complex64.
func ToComplex64(value interface{}) complex64 { return basic.ToComplex64(value) }

// ToComplex64E converts any type to complex64 with error.
func ToComplex64E(value interface{}) (complex64, error) { return basic.ToComplex64E(value) }

// ToTime converts any type to time.Time.
func ToTime(value interface{}, formats ...string) time.Time { return basic.ToTime(value, formats...) }

// ToTimeE converts any type to time.Time with error.
func ToTimeE(value interface{}, formats ...string) (time.Time, error) {
	return basic.ToTimeE(value, formats...)
}

// ToDuration converts any type to time.Duration.
func ToDuration(value interface{}) time.Duration { return basic.ToDuration(value) }

// ToDurationE converts any type to time.Duration with error.
func ToDurationE(value interface{}) (time.Duration, error) { return basic.ToDurationE(value) }

// ToSlice converts any type to slice.
func ToSlice(value interface{}) []interface{} { return complex.ToSlice(value) }

// ToSliceE converts any type to slice with error.
func ToSliceE(value interface{}) ([]interface{}, error) { return complex.ToSliceE(value) }

// ToStringSlice converts any type to a string slice.
func ToStringSlice(value interface{}) []string { return complex.ToStringSlice(value) }

// ToStringSliceE converts any type to a string slice with error.
func ToStringSliceE(value interface{}) ([]string, error) { return complex.ToStringSliceE(value) }

// ToIntSlice converts any type to an int slice.
func ToIntSlice(value interface{}) []int { return complex.ToIntSlice(value) }

// ToIntSliceE converts any type to an int slice with error.
func ToIntSliceE(value interface{}) ([]int, error) { return complex.ToIntSliceE(value) }

// ToFloat64Slice converts any type to a float64 slice.
func ToFloat64Slice(value interface{}) []float64 { return complex.ToFloat64Slice(value) }

// ToFloat64SliceE converts any type to a float64 slice with error.
func ToFloat64SliceE(value interface{}) ([]float64, error) { return complex.ToFloat64SliceE(value) }

// ToMap converts any type to a map.
func ToMap(value interface{}) map[string]interface{} { return complex.ToMap(value) }

// ToMapE converts any type to a map with error.
func ToMapE(value interface{}) (map[string]interface{}, error) { return complex.ToMapE(value) }

// ToStringMap converts any type to a string map.
func ToStringMap(value interface{}) map[string]string { return complex.ToStringMap(value) }

// ToStringMapE converts any type to a string map with error.
func ToStringMapE(value interface{}) (map[string]string, error) { return complex.ToStringMapE(value) }

// ToIntMap converts any type to an int map.
func ToIntMap(value interface{}) map[string]int { return complex.ToIntMap(value) }

// ToIntMapE converts any type to an int map with error.
func ToIntMapE(value interface{}) (map[string]int, error) { return complex.ToIntMapE(value) }

// ToFloat64Map converts any type to a float64 map.
func ToFloat64Map(value interface{}) map[string]float64 { return complex.ToFloat64Map(value) }

// ToFloat64MapE converts any type to a float64 map with error.
func ToFloat64MapE(value interface{}) (map[string]float64, error) {
	return complex.ToFloat64MapE(value)
}

// ToJSON converts any type to JSON.
func ToJSON(value interface{}) string { return complex.ToJSON(value) }

// ToJSONE converts any type to JSON with error.
func ToJSONE(value interface{}) (string, error) { return complex.ToJSONE(value) }

// FromJSON converts JSON into target.
func FromJSON(jsonStr string, target interface{}) { complex.FromJSON(jsonStr, target) }

// FromJSONE converts JSON into target with error.
func FromJSONE(jsonStr string, target interface{}) error { return complex.FromJSONE(jsonStr, target) }

// ToMapFromJSON converts JSON to a map.
func ToMapFromJSON(jsonStr string) map[string]interface{} { return complex.ToMapFromJSON(jsonStr) }

// ToMapFromJSONE converts JSON to a map with error.
func ToMapFromJSONE(jsonStr string) (map[string]interface{}, error) {
	return complex.ToMapFromJSONE(jsonStr)
}

// ToSliceFromJSON converts JSON to a slice.
func ToSliceFromJSON(jsonStr string) []interface{} { return complex.ToSliceFromJSON(jsonStr) }

// ToSliceFromJSONE converts JSON to a slice with error.
func ToSliceFromJSONE(jsonStr string) ([]interface{}, error) {
	return complex.ToSliceFromJSONE(jsonStr)
}

// ToStruct converts a map or struct to a struct.
func ToStruct(source, pointer interface{}, hooks ...HookFunc) {
	complex.ToStruct(source, pointer, hooks...)
}

// ToStructE converts a map or struct to a struct with error.
func ToStructE(source, pointer interface{}, hooks ...HookFunc) error {
	return complex.ToStructE(source, pointer, hooks...)
}

// SetStringCacheSize sets the string conversion cache size.
func SetStringCacheSize(size int) { internal.SetStringCacheSize(size) }

// SetTimeCacheSize sets the time conversion cache size.
func SetTimeCacheSize(size int) { internal.SetTimeCacheSize(size) }

// ClearStringCache clears the string conversion cache.
func ClearStringCache() { internal.ClearStringCache() }

// ClearTimeCache clears the time conversion cache.
func ClearTimeCache() { internal.ClearTimeCache() }

// ClearAllCaches clears all conversion caches.
func ClearAllCaches() { internal.ClearAllCaches() }

// SetTypeInfoCacheSize sets the reflection type-info cache size.
func SetTypeInfoCacheSize(size int) { internal.SetTypeInfoCacheSize(size) }

// SetConversionCacheSize sets the reflection conversion cache size.
func SetConversionCacheSize(size int) { internal.SetConversionCacheSize(size) }

// ClearTypeInfoCache clears the reflection type-info cache.
func ClearTypeInfoCache() { internal.ClearTypeInfoCache() }

// ClearConversionCache clears the reflection conversion cache.
func ClearConversionCache() { internal.ClearConversionCache() }

// Note: This library also provides generic conversion functions, which need to
// be imported directly from the complex package.
//
// Generic slice conversion functions:
// - complex.ToSliceT[T any](value interface{}) []T
//   Convert any type to []T type slice
//   Example:
//     strSlice := complex.ToSliceT[string](value) // Convert to []string
//     intSlice := complex.ToSliceT[int](value)    // Convert to []int
//
// - complex.ToSliceTE[T any](value interface{}) ([]T, error)
//   Convert any type to []T type slice, and return the possible error
//   Example:
//     strSlice, err := complex.ToSliceTE[string](value) // Convert to []string
//     intSlice, err := complex.ToSliceTE[int](value)    // Convert to []int
//
// Generic map conversion functions:
// - complex.ToMapT[K comparable, V any](value interface{}) map[K]V
//   Convert any type to map[K]V type map
//   Example:
//     strMap := complex.ToMapT[string, string](value) // Convert to map[string]string
//     intMap := complex.ToMapT[string, int](value)    // Convert to map[string]int
//
// - complex.ToMapTE[K comparable, V any](value interface{}) (map[K]V, error)
//   Convert any type to map[K]V type map, and return the possible error
//   Example:
//     strMap, err := complex.ToMapTE[string, string](value) // Convert to map[string]string
//     intMap, err := complex.ToMapTE[string, int](value)    // Convert to map[string]int
