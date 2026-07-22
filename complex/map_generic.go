package complex

import (
	"reflect"

	"github.com/graingo/mconv/internal"
)

// ToMapT converts any type to map[K]V.
// This is a generic version of ToMap that returns a map with key type K and value type V.
// It uses reflection caching to improve performance for repeated conversions.
//
// Examples:
//
//	// Convert to map[string]string
//	strMap := ToMapT[string, string](value)
//
//	// Convert to map[string]int
//	intMap := ToMapT[string, int](value)
//
//	// Convert to map[int]float64
//	floatMap := ToMapT[int, float64](value)
func ToMapT[K comparable, V any](value interface{}) map[K]V {
	result, _ := ToMapTE[K, V](value)
	return result
}

// ToMapTE converts any type to map[K]V with error.
// This is a generic version of ToMapE that returns a map with key type K and value type V.
// It uses reflection caching to improve performance for repeated conversions.
//
// Examples:
//
//	// Convert to map[string]string with error handling
//	strMap, err := ToMapTE[string, string](value)
//
//	// Convert to map[string]int with error handling
//	intMap, err := ToMapTE[string, int](value)
//
//	// Convert to map[int]float64 with error handling
//	floatMap, err := ToMapTE[int, float64](value)
func ToMapTE[K comparable, V any](value interface{}) (map[K]V, error) {
	if value == nil {
		return nil, nil
	}

	if v, ok := value.(map[K]V); ok {
		return v, nil
	}

	m, err := ToMapE(value)
	if err != nil {
		return nil, internal.NewConversionError(value, "map[K]V", err)
	}

	kt := reflect.TypeOf((*K)(nil)).Elem()
	vt := reflect.TypeOf((*V)(nil)).Elem()

	result := make(map[K]V)

	for k, v := range m {
		keyValue := reflect.New(kt).Elem()
		if err := setGenericValue(keyValue, k); err != nil {
			return nil, internal.NewConversionError(k, kt.String(), err)
		}
		keyConverted := keyValue.Interface().(K)

		valueTarget := reflect.New(vt).Elem()
		if v == nil {
			result[keyConverted] = valueTarget.Interface().(V)
			continue
		}
		if err := setGenericValue(valueTarget, v); err != nil {
			return nil, internal.NewConversionError(v, vt.String(), err)
		}
		result[keyConverted] = valueTarget.Interface().(V)
	}

	return result, nil
}
