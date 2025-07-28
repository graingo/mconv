package complex

import (
	"reflect"

	"github.com/graingo/mconv/basic"
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
		var keyConverted K
		keyRv := reflect.ValueOf(k)

		var keyErr error
		if keyRv.Type().ConvertibleTo(kt) {
			keyConverted = keyRv.Convert(kt).Interface().(K)
		} else {
			// try to convert via basic types
			switch kt.Kind() {
			case reflect.String:
				keyConverted = any(basic.ToString(k)).(K)
			case reflect.Int:
				i, err := basic.ToIntE(k)
				if err != nil {
					keyErr = err
				} else {
					keyConverted = any(i).(K)
				}
			case reflect.Int64:
				i, err := basic.ToInt64E(k)
				if err != nil {
					keyErr = err
				} else {
					keyConverted = any(i).(K)
				}
			default:
				keyErr = internal.NewConversionError(k, "K", internal.ErrConversionFailed)
			}
		}
		if keyErr != nil {
			return nil, keyErr
		}

		// Convert value
		var valueConverted V
		var valueErr error
		if v == nil {
			result[keyConverted] = valueConverted
			continue
		}
		valueRv := reflect.ValueOf(v)
		if valueRv.Type().ConvertibleTo(vt) {
			valueConverted = valueRv.Convert(vt).Interface().(V)
		} else {
			// try to convert via basic types
			switch vt.Kind() {
			case reflect.String:
				valueConverted = any(basic.ToString(v)).(V)
			case reflect.Int:
				i, err := basic.ToIntE(v)
				if err != nil {
					valueErr = err
				} else {
					valueConverted = any(i).(V)
				}
			case reflect.Int64:
				i, err := basic.ToInt64E(v)
				if err != nil {
					valueErr = err
				} else {
					valueConverted = any(i).(V)
				}
			case reflect.Float64:
				f, err := basic.ToFloat64E(v)
				if err != nil {
					valueErr = err
				} else {
					valueConverted = any(f).(V)
				}
			case reflect.Bool:
				b, err := basic.ToBoolE(v)
				if err != nil {
					valueErr = err
				} else {
					valueConverted = any(b).(V)
				}
			default:
				valueErr = internal.NewConversionError(v, "V", internal.ErrConversionFailed)
			}
		}
		if valueErr != nil {
			return nil, valueErr
		}

		result[keyConverted] = valueConverted
	}

	return result, nil
}
