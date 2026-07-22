package complex

import (
	"reflect"

	"github.com/graingo/mconv/internal"
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

		converted := reflect.New(targetType).Elem()
		if err := setGenericValue(converted, v); err != nil {
			return nil, internal.NewConversionError(v, targetType.String(), err)
		}
		result[i] = converted.Interface().(T)
	}

	return result, nil
}

func setGenericValue(target reflect.Value, value interface{}) error {
	return setFieldValue(target, value, stringToTimeHookFunc(), stringToDurationHookFunc(), intToBoolHookFunc())
}
