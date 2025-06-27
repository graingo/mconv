package mconv

import (
	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/complex"
	"github.com/graingo/mconv/internal"
)

// HookFunc is an alias of complex.HookFunc.
type HookFunc = complex.HookFunc

var (
	// ToString convert any type to string.
	ToString = basic.ToString
	// ToStringE convert any type to string with error.
	ToStringE = basic.ToStringE

	// ToInt convert any type to int.
	ToInt = basic.ToInt
	// ToIntE convert any type to int with error.
	ToIntE = basic.ToIntE
	// ToInt64 convert any type to int64.
	ToInt64 = basic.ToInt64
	// ToInt64E convert any type to int64 with error.
	ToInt64E = basic.ToInt64E
	// ToInt32 convert any type to int32.
	ToInt32 = basic.ToInt32
	// ToInt32E convert any type to int32 with error.
	ToInt32E = basic.ToInt32E
	// ToInt16 convert any type to int16.
	ToInt16 = basic.ToInt16
	// ToInt16E convert any type to int16 with error.
	ToInt16E = basic.ToInt16E
	// ToInt8 convert any type to int8.
	ToInt8 = basic.ToInt8
	// ToInt8E convert any type to int8 with error.
	ToInt8E = basic.ToInt8E

	// ToUint convert any type to uint.
	ToUint = basic.ToUint
	// ToUintE convert any type to uint with error.
	ToUintE = basic.ToUintE
	// ToUint64 convert any type to uint64.
	ToUint64 = basic.ToUint64
	// ToUint64E convert any type to uint64 with error.
	ToUint64E = basic.ToUint64E
	// ToUint32 convert any type to uint32.
	ToUint32 = basic.ToUint32
	// ToUint32E convert any type to uint32 with error.
	ToUint32E = basic.ToUint32E
	// ToUint16 convert any type to uint16.
	ToUint16 = basic.ToUint16
	// ToUint16E convert any type to uint16 with error.
	ToUint16E = basic.ToUint16E
	// ToUint8 convert any type to uint8.
	ToUint8 = basic.ToUint8
	// ToUint8E convert any type to uint8 with error.
	ToUint8E = basic.ToUint8E

	// ToFloat64 convert any type to float64.
	ToFloat64 = basic.ToFloat64
	// ToFloat64E convert any type to float64 with error.
	ToFloat64E = basic.ToFloat64E
	// ToFloat32 convert any type to float32.
	ToFloat32 = basic.ToFloat32
	// ToFloat32E convert any type to float32 with error.
	ToFloat32E = basic.ToFloat32E

	// ToBool convert any type to bool.
	ToBool = basic.ToBool
	// ToBoolE convert any type to bool with error.
	ToBoolE = basic.ToBoolE

	// ToComplex128 convert any type to complex128.
	ToComplex128 = basic.ToComplex128
	// ToComplex128E convert any type to complex128 with error.
	ToComplex128E = basic.ToComplex128E
	// ToComplex64 convert any type to complex64.
	ToComplex64 = basic.ToComplex64
	// ToComplex64E convert any type to complex64 with error.
	ToComplex64E = basic.ToComplex64E
	// ToTime convert any type to time.Time.
	ToTime = basic.ToTime
	// ToTimeE convert any type to time.Time with error.
	ToTimeE = basic.ToTimeE
	// ToDuration convert any type to time.Duration.
	ToDuration = basic.ToDuration
	// ToDurationE convert any type to time.Duration with error.
	ToDurationE = basic.ToDurationE
)

var (
	// ToSlice convert any type to slice.
	ToSlice = complex.ToSlice
	// ToSliceE convert any type to slice with error.
	ToSliceE = complex.ToSliceE
	// ToStringSlice convert any type to slice of string.
	ToStringSlice = complex.ToStringSlice
	// ToStringSliceE convert any type to slice of string with error.
	ToStringSliceE = complex.ToStringSliceE
	// ToIntSlice convert any type to slice of int.
	ToIntSlice = complex.ToIntSlice
	// ToIntSliceE convert any type to slice of int with error.
	ToIntSliceE = complex.ToIntSliceE
	// ToFloat64Slice convert any type to slice of float64.
	ToFloat64Slice = complex.ToFloat64Slice
	// ToFloat64SliceE convert any type to slice of float64 with error.
	ToFloat64SliceE = complex.ToFloat64SliceE

	// ToMap convert any type to map.
	ToMap = complex.ToMap
	// ToMapE convert any type to map with error.
	ToMapE = complex.ToMapE
	// ToStringMap convert any type to map of string.
	ToStringMap = complex.ToStringMap
	// ToStringMapE convert any type to map of string with error.
	ToStringMapE = complex.ToStringMapE
	// ToIntMap convert any type to map of int.
	ToIntMap = complex.ToIntMap
	// ToIntMapE convert any type to map of int with error.
	ToIntMapE = complex.ToIntMapE
	// ToFloat64Map convert any type to map of float64.
	ToFloat64Map = complex.ToFloat64Map
	// ToFloat64MapE convert any type to map of float64 with error.
	ToFloat64MapE = complex.ToFloat64MapE

	// ToJSON convert any type to json.
	ToJSON = complex.ToJSON
	// ToJSONE convert any type to json with error.
	ToJSONE = complex.ToJSONE
	// FromJSON convert json to any type.
	FromJSON = complex.FromJSON
	// FromJSONE convert json to any type with error.
	FromJSONE = complex.FromJSONE
	// ToMapFromJSON convert json to map.
	ToMapFromJSON = complex.ToMapFromJSON
	// ToMapFromJSONE convert json to map with error.
	ToMapFromJSONE = complex.ToMapFromJSONE
	// ToSliceFromJSON convert json to slice.
	ToSliceFromJSON = complex.ToSliceFromJSON
	// ToSliceFromJSONE convert json to slice with error.
	ToSliceFromJSONE = complex.ToSliceFromJSONE

	// ToStruct convert map or struct to struct.
	ToStruct = complex.ToStruct
	// ToStructE convert map or struct to struct with error.
	ToStructE = complex.ToStructE
)

// Export cache related functions
var (
	SetStringCacheSize = internal.SetStringCacheSize
	SetTimeCacheSize   = internal.SetTimeCacheSize
	ClearStringCache   = internal.ClearStringCache
	ClearTimeCache     = internal.ClearTimeCache
	ClearAllCaches     = internal.ClearAllCaches

	// Reflection cache
	SetTypeInfoCacheSize   = internal.SetTypeInfoCacheSize
	SetConversionCacheSize = internal.SetConversionCacheSize
	ClearTypeInfoCache     = internal.ClearTypeInfoCache
	ClearConversionCache   = internal.ClearConversionCache
)

// Note: This library also provides generic conversion functions, which need to be imported directly from the complex package.
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
