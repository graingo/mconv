package complex

import (
	"fmt"
	"reflect"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/internal"
)

// ToMap converts any type to map[string]interface{}.
func ToMap(value interface{}) map[string]interface{} {
	result, _ := ToMapE(value)
	return result
}

// ToMapE converts any type to map[string]interface{} with error.
func ToMapE(value interface{}) (map[string]interface{}, error) {
	if isNilCollectionInput(value) {
		return nil, nil
	}

	switch v := value.(type) {
	case map[string]interface{}:
		return v, nil
	case map[interface{}]interface{}:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			key, err := basic.ToStringE(k)
			if err != nil {
				return nil, internal.NewConversionError(k, "map", err)
			}
			if _, exists := result[key]; exists {
				return nil, internal.NewConversionError(k, "map", internal.ErrConversionFailed)
			}
			result[key] = val
		}
		return result, nil
	case map[string]string:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]int:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]int64:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]float64:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]float32:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]complex64:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]complex128:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	case map[string]bool:
		result := make(map[string]interface{}, len(v))
		for k, val := range v {
			result[k] = val
		}
		return result, nil
	default:
		rv := reflect.ValueOf(value)
		for rv.IsValid() && (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) {
			if rv.IsNil() {
				return nil, nil
			}
			rv = rv.Elem()
		}
		if !rv.IsValid() {
			return nil, nil
		}
		if rv.Kind() == reflect.Struct {
			return structToMap(rv)
		}
		if rv.Kind() != reflect.Map {
			return nil, internal.NewConversionError(value, "map", internal.ErrUnsupportedType)
		}
		if rv.IsNil() {
			return nil, nil
		}

		result := make(map[string]interface{}, rv.Len())
		for _, key := range rv.MapKeys() {
			keyStr, err := basic.ToStringE(key.Interface())
			if err != nil {
				return nil, internal.NewConversionError(key.Interface(), "map", err)
			}
			if _, exists := result[keyStr]; exists {
				return nil, internal.NewConversionError(key.Interface(), "map", internal.ErrConversionFailed)
			}
			result[keyStr] = rv.MapIndex(key).Interface()
		}
		return result, nil
	}
}

func structToMap(value reflect.Value) (map[string]interface{}, error) {
	decoder := getDecoder(value.Type())
	result := make(map[string]interface{}, len(decoder.FieldArr))
	for _, fieldDecoder := range decoder.FieldArr {
		if fieldDecoder.Ambiguous {
			return nil, fmt.Errorf("struct contains ambiguous map key %q", fieldDecoder.Name)
		}
		fieldValue := fieldByIndex(value, fieldDecoder.Index)
		if fieldValue.IsValid() {
			result[fieldDecoder.Name] = fieldValue.Interface()
		}
	}
	return result, nil
}

// ToStringMap converts any type to map[string]string
func ToStringMap(value interface{}) map[string]string {
	result, _ := ToStringMapE(value)
	return result
}

// ToStringMapE converts any type to map[string]string with error
func ToStringMapE(value interface{}) (map[string]string, error) {
	if isNilCollectionInput(value) {
		return nil, nil
	}

	if v, ok := value.(map[string]string); ok {
		return v, nil
	}

	source, err := ToMapE(value)
	if err != nil {
		return nil, internal.NewConversionError(value, "map[string]string", err)
	}
	if source == nil {
		return nil, nil
	}
	result := make(map[string]string, len(source))
	for key, item := range source {
		converted, err := basic.ToStringE(item)
		if err != nil {
			return nil, internal.NewConversionError(item, "string", err)
		}
		result[key] = converted
	}
	return result, nil
}

// ToIntMap converts any type to map[string]int.
func ToIntMap(value interface{}) map[string]int {
	result, _ := ToIntMapE(value)
	return result
}

// ToIntMapE converts any type to map[string]int with error.
func ToIntMapE(value interface{}) (map[string]int, error) {
	if isNilCollectionInput(value) {
		return nil, nil
	}

	switch v := value.(type) {
	case map[string]int:
		return v, nil
	case map[string]interface{}:
		result := make(map[string]int, len(v))
		for k, val := range v {
			intVal, err := basic.ToIntE(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "map", err)
			}
			result[k] = intVal
		}
		return result, nil
	case map[string]string:
		result := make(map[string]int, len(v))
		for k, val := range v {
			intVal, err := basic.ToIntE(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "map", err)
			}
			result[k] = intVal
		}
		return result, nil
	default:
		// try to convert other type to map[string]interface{} and then handle
		m, err := ToMapE(value)
		if err != nil {
			return nil, internal.NewConversionError(value, "map", err)
		}
		return ToIntMapE(m)
	}
}

// ToFloat64Map converts any type to map[string]float64.
func ToFloat64Map(value interface{}) map[string]float64 {
	result, _ := ToFloat64MapE(value)
	return result
}

// ToFloat64MapE converts any type to map[string]float64 with error.
func ToFloat64MapE(value interface{}) (map[string]float64, error) {
	if isNilCollectionInput(value) {
		return nil, nil
	}

	switch v := value.(type) {
	case map[string]float64:
		return v, nil
	case map[string]interface{}:
		result := make(map[string]float64, len(v))
		for k, val := range v {
			floatVal, err := basic.ToFloat64E(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "map", err)
			}
			result[k] = floatVal
		}
		return result, nil
	case map[string]string:
		result := make(map[string]float64, len(v))
		for k, val := range v {
			floatVal, err := basic.ToFloat64E(val)
			if err != nil {
				return nil, internal.NewConversionError(val, "map", err)
			}
			result[k] = floatVal
		}
		return result, nil
	default:
		// try to convert other type to map[string]interface{} and then handle
		m, err := ToMapE(value)
		if err != nil {
			return nil, internal.NewConversionError(value, "map", err)
		}
		return ToFloat64MapE(m)
	}
}
