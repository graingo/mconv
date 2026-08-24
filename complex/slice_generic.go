package complex

import (
	"fmt"
	"reflect"

	"github.com/graingo/mconv/internal"
)

// ToSliceT converts any type to []T.
// Scalar targets use direct converters; defined and complex targets use the
// shared reflection decoder.
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
// Errors include the failing element index in ConversionError.Path.
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
	if isNilCollectionInput(value) {
		return nil, nil
	}

	if v, ok := value.([]T); ok {
		return v, nil
	}

	source, valid := indirectValue(value)
	if !valid {
		return nil, nil
	}

	var length int
	collection := source.Kind() == reflect.Slice || source.Kind() == reflect.Array
	if collection {
		length = source.Len()
	} else {
		length = 1
	}

	result := make([]T, length)
	convert := genericConverter[T]()

	for i := 0; i < length; i++ {
		item := source
		if collection {
			item = source.Index(i)
		}
		if isNilReflectValue(item) {
			continue
		}

		itemValue := item.Interface()
		converted, err := convertGenericValue(itemValue, convert)
		if err != nil {
			return nil, internal.PrependConversionPath(err, fmt.Sprintf("[%d]", i))
		}
		result[i] = converted
	}

	return result, nil
}

func isNilReflectValue(value reflect.Value) bool {
	if !value.IsValid() {
		return true
	}
	switch value.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return value.IsNil()
	default:
		return false
	}
}

func setGenericValue(target reflect.Value, value interface{}) error {
	return setFieldValue(target, value, defaultHooks...)
}
