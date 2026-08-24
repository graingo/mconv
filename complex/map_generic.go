package complex

import (
	"fmt"
	"reflect"

	"github.com/graingo/mconv/internal"
)

// ToMapT converts any type to map[K]V.
// Source keys are converted directly to K without a string intermediate.
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
// It rejects key collisions and includes the failing key in ConversionError.Path.
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
	if isNilCollectionInput(value) {
		return nil, nil
	}

	if v, ok := value.(map[K]V); ok {
		return v, nil
	}

	source, valid := indirectValue(value)
	if !valid {
		return nil, nil
	}
	if source.Kind() == reflect.Struct {
		converted, err := ToMapE(source.Interface())
		if err != nil {
			return nil, internal.NewConversionError(value, "map[K]V", err)
		}
		source = reflect.ValueOf(converted)
	}
	if source.Kind() != reflect.Map {
		return nil, internal.NewConversionError(value, "map[K]V", internal.ErrUnsupportedType)
	}

	kt := reflect.TypeOf((*K)(nil)).Elem()

	result := make(map[K]V, source.Len())
	keysRemainUnique := source.Type().Key().AssignableTo(kt)
	convertKey := genericConverter[K]()
	convertValue := genericConverter[V]()
	iterator := source.MapRange()
	for iterator.Next() {
		sourceKey := iterator.Key().Interface()
		key, err := convertGenericValue(sourceKey, convertKey)
		if err != nil {
			return nil, internal.PrependConversionPath(err, fmt.Sprintf("[%v]", sourceKey))
		}
		if !keysRemainUnique {
			if _, exists := result[key]; exists {
				collisionErr := internal.NewConversionError(sourceKey, kt.String(), internal.ErrConversionFailed)
				return nil, internal.PrependConversionPath(collisionErr, fmt.Sprintf("[%v]", sourceKey))
			}
		}

		sourceValue := iterator.Value()
		var convertedValue V
		if !isNilReflectValue(sourceValue) {
			item := sourceValue.Interface()
			convertedValue, err = convertGenericValue(item, convertValue)
			if err != nil {
				return nil, internal.PrependConversionPath(err, fmt.Sprintf("[%v]", sourceKey))
			}
		}
		result[key] = convertedValue
	}

	return result, nil
}
