package complex

import (
	"reflect"

	"github.com/mingzaily/mconv/basic"
	"github.com/mingzaily/mconv/internal"
)

// ToSliceT converts any type to []T.
// This is a generic version of ToSlice that returns a slice of type T.
// It uses reflection caching to improve performance for repeated conversions.
//
// Examples:
//
//	// Convert to []string
//	strSlice := ToSliceT[string](value)
//
//	// Convert to []int
//	intSlice := ToSliceT[int](value)
//
//	// Convert to []float64
//	floatSlice := ToSliceT[float64](value)
func ToSliceT[T any](value interface{}) []T {
	result, _ := ToSliceTE[T](value)
	return result
}

// ToSliceTE converts any type to []T with error.
// This is a generic version of ToSliceE that returns a slice of type T.
// It uses reflection caching to improve performance for repeated conversions.
//
// Examples:
//
//	// Convert to []string with error handling
//	strSlice, err := ToSliceTE[string](value)
//
//	// Convert to []int with error handling
//	intSlice, err := ToSliceTE[int](value)
//
//	// Convert to []float64 with error handling
//	floatSlice, err := ToSliceTE[float64](value)
func ToSliceTE[T any](value interface{}) ([]T, error) {
	if value == nil {
		return nil, nil
	}

	// Check if value is already a []T
	if v, ok := value.([]T); ok {
		return v, nil
	}

	// Get target type
	targetType := reflect.TypeOf((*T)(nil)).Elem()

	// Convert to []interface{} first
	s, err := ToSliceE(value)
	if err != nil {
		return nil, internal.NewConversionError(value, "[]T", err)
	}

	// Create result slice
	result := make([]T, len(s))

	// Convert each element
	for i, v := range s {
		if v == nil {
			continue
		}

		// Get value reflection value
		vValue := reflect.ValueOf(v)
		// Get cached type information to avoid repeated reflection operations
		vTypeInfo := internal.GetTypeInfo(vValue.Type())

		// Handle common types
		switch any(result[0]).(type) {
		case string:
			strVal, e := basic.ToStringE(v)
			if e != nil {
				return nil, internal.NewConversionError(v, "T", e)
			}
			result[i] = any(strVal).(T)
		case int:
			intVal, e := basic.ToIntE(v)
			if e != nil {
				return nil, internal.NewConversionError(v, "T", e)
			}
			result[i] = any(intVal).(T)
		case int64:
			int64Val, e := basic.ToInt64E(v)
			if e != nil {
				return nil, internal.NewConversionError(v, "T", e)
			}
			result[i] = any(int64Val).(T)
		case float64:
			floatVal, e := basic.ToFloat64E(v)
			if e != nil {
				return nil, internal.NewConversionError(v, "T", e)
			}
			result[i] = any(floatVal).(T)
		case bool:
			boolVal, e := basic.ToBoolE(v)
			if e != nil {
				return nil, internal.NewConversionError(v, "T", e)
			}
			result[i] = any(boolVal).(T)
		default:
			// For other types, try direct conversion using cached type information
			// Use cached assignability and convertibility checks for better performance
			if vTypeInfo.IsAssignableTo(targetType) {
				result[i] = v.(T)
			} else if vTypeInfo.IsConvertibleTo(targetType) {
				// Try to convert using reflection
				converted := vValue.Convert(targetType)
				result[i] = converted.Interface().(T)
			} else {
				return nil, internal.NewConversionError(v, "T", internal.ErrConversionFailed)
			}
		}
	}

	return result, nil
}
