package complex

import (
	"reflect"

	"github.com/mingzaily/mconv/basic"
	"github.com/mingzaily/mconv/internal"
)

// ToSlice converts any type to []interface{}.
func ToSlice(value interface{}) []interface{} {
	result, _ := ToSliceE(value)
	return result
}

// ToSliceE converts any type to []interface{} with error.
func ToSliceE(value interface{}) ([]interface{}, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []interface{}:
		return v, nil
	case []string:
		result := make([]interface{}, len(v))
		for i, s := range v {
			result[i] = s
		}
		return result, nil
	case []int:
		result := make([]interface{}, len(v))
		for i, j := range v {
			result[i] = j
		}
		return result, nil
	case []int64:
		result := make([]interface{}, len(v))
		for i, j := range v {
			result[i] = j
		}
		return result, nil
	case []float64:
		result := make([]interface{}, len(v))
		for i, f := range v {
			result[i] = f
		}
		return result, nil
	case []complex64:
		result := make([]interface{}, len(v))
		for i, c := range v {
			result[i] = c
		}
		return result, nil
	case []complex128:
		result := make([]interface{}, len(v))
		for i, c := range v {
			result[i] = c
		}
		return result, nil
	case string:
		return []interface{}{v}, nil
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, complex64, complex128, bool:
		return []interface{}{v}, nil
	default:
		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.Slice {
			return []interface{}{value}, nil
		}

		sliceLen := rv.Len()
		result := make([]interface{}, sliceLen)
		for i := 0; i < sliceLen; i++ {
			result[i] = rv.Index(i).Interface()
		}
		return result, nil
	}
}

// ToStringSlice converts any type to []string
func ToStringSlice(value interface{}) []string {
	result, _ := ToStringSliceE(value)
	return result
}

// ToStringSliceE converts any type to []string with error
func ToStringSliceE(value interface{}) ([]string, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []string:
		return v, nil
	case []interface{}:
		result := make([]string, len(v))
		for i, val := range v {
			str, err := basic.ToStringE(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "slice", err)
			}
			result[i] = str
		}
		return result, nil
	case string:
		str := v
		return []string{str}, nil
	default:
		rv := reflect.ValueOf(value)
		if rv.Kind() != reflect.Slice {
			str, err := basic.ToStringE(value)
			if err != nil {
				return nil, internal.NewConversionError(value, "slice", err)
			}
			return []string{str}, nil
		}

		result := make([]string, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			str, err := basic.ToStringE(rv.Index(i).Interface())
			if err != nil {
				return nil, internal.NewConversionError(value, "slice", err)
			}
			result[i] = str
		}
		return result, nil
	}
}

// ToIntSlice converts any type to []int.
func ToIntSlice(value interface{}) []int {
	result, _ := ToIntSliceE(value)
	return result
}

// ToIntSliceE converts any type to []int with error.
func ToIntSliceE(value interface{}) ([]int, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []int:
		return v, nil
	case []interface{}:
		result := make([]int, len(v))
		for i, val := range v {
			intVal, err := basic.ToIntE(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "slice", err)
			}
			result[i] = intVal
		}
		return result, nil
	case []string:
		result := make([]int, len(v))
		for i, val := range v {
			intVal, err := basic.ToIntE(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "slice", err)
			}
			result[i] = intVal
		}
		return result, nil
	default:
		// try to convert other type to []interface{} and then handle
		slice, err := ToSliceE(value)
		if err != nil {
			return nil, internal.NewConversionError(value, "slice", err)
		}
		return ToIntSliceE(slice)
	}
}

// ToFloat64Slice converts any type to []float64.
func ToFloat64Slice(value interface{}) []float64 {
	result, _ := ToFloat64SliceE(value)
	return result
}

// ToFloat64SliceE converts any type to []float64 with error.
func ToFloat64SliceE(value interface{}) ([]float64, error) {
	if value == nil {
		return nil, nil
	}

	switch v := value.(type) {
	case []float64:
		return v, nil
	case []interface{}:
		result := make([]float64, len(v))
		for i, val := range v {
			floatVal, err := basic.ToFloat64E(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "slice", err)
			}
			result[i] = floatVal
		}
		return result, nil
	case []string:
		result := make([]float64, len(v))
		for i, val := range v {
			floatVal, err := basic.ToFloat64E(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "slice", err)
			}
			result[i] = floatVal
		}
		return result, nil
	default:
		// try to convert other type to []interface{} and then handle
		slice, err := ToSliceE(value)
		if err != nil {
			return nil, internal.NewConversionError(value, "slice", err)
		}
		return ToFloat64SliceE(slice)
	}
}
