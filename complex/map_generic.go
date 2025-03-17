package complex

import (
	"reflect"

	"github.com/mingzaily/mconv/basic"
	"github.com/mingzaily/mconv/internal"
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

	// Check if value is already a map[K]V
	if v, ok := value.(map[K]V); ok {
		return v, nil
	}

	// Convert to map[string]interface{} first
	m, err := ToMapE(value)
	if err != nil {
		return nil, internal.NewConversionError(value, "map[K]V", err)
	}

	// Get target key and value types
	kt := reflect.TypeOf((*K)(nil)).Elem()
	vt := reflect.TypeOf((*V)(nil)).Elem()

	// Create result map
	result := make(map[K]V)

	// Convert each key-value pair
	for k, v := range m {
		// Convert key
		var keyConverted K
		var keyErr error

		// Get key reflection value
		kValue := reflect.ValueOf(k)
		// Get cached type information to avoid repeated reflection operations
		kTypeInfo := internal.GetTypeInfo(kValue.Type())

		// Handle common key types
		switch any(keyConverted).(type) {
		case string:
			// If K is string, we already have string keys
			if kt == reflect.TypeOf("") {
				// Use cached assignability check for better performance
				if !kTypeInfo.IsAssignableTo(kt) {
					keyErr = internal.NewConversionError(k, "K", internal.ErrConversionFailed)
				}
			}
		case int:
			intKey, e := basic.ToIntE(k)
			if e != nil {
				keyErr = internal.NewConversionError(k, "K", e)
			} else {
				keyConverted = any(intKey).(K)
			}
		default:
			// For other types, try direct conversion using cached convertibility check
			// This avoids repeated calls to ConvertibleTo which is expensive
			if !kTypeInfo.IsConvertibleTo(kt) {
				keyErr = internal.NewConversionError(k, "K", internal.ErrConversionFailed)
			}
		}

		if keyErr != nil {
			return nil, keyErr
		}

		// Convert value
		var valueConverted V
		var valueErr error

		// Get value reflection value
		vValue := reflect.ValueOf(v)
		// Get cached type information to avoid repeated reflection operations
		vTypeInfo := internal.GetTypeInfo(vValue.Type())

		// Handle common value types
		switch any(valueConverted).(type) {
		case string:
			strVal, e := basic.ToStringE(v)
			if e != nil {
				valueErr = internal.NewConversionError(v, "V", e)
			} else {
				valueConverted = any(strVal).(V)
			}
		case int:
			intVal, e := basic.ToIntE(v)
			if e != nil {
				valueErr = internal.NewConversionError(v, "V", e)
			} else {
				valueConverted = any(intVal).(V)
			}
		case int64:
			int64Val, e := basic.ToInt64E(v)
			if e != nil {
				valueErr = internal.NewConversionError(v, "V", e)
			} else {
				valueConverted = any(int64Val).(V)
			}
		case float64:
			floatVal, e := basic.ToFloat64E(v)
			if e != nil {
				valueErr = internal.NewConversionError(v, "V", e)
			} else {
				valueConverted = any(floatVal).(V)
			}
		case bool:
			boolVal, e := basic.ToBoolE(v)
			if e != nil {
				valueErr = internal.NewConversionError(v, "V", e)
			} else {
				valueConverted = any(boolVal).(V)
			}
		default:
			// For other types, try direct conversion using cached type information
			// Use cached assignability and convertibility checks for better performance
			if vTypeInfo.IsAssignableTo(vt) {
				valueConverted = v.(V)
			} else if vTypeInfo.IsConvertibleTo(vt) {
				// Try to convert using reflection
				converted := vValue.Convert(vt)
				valueConverted = converted.Interface().(V)
			} else {
				valueErr = internal.NewConversionError(v, "V", internal.ErrConversionFailed)
			}
		}

		if valueErr != nil {
			return nil, valueErr
		}

		// Add to result map
		result[keyConverted] = valueConverted
	}

	return result, nil
}
