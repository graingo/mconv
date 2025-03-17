package complex

import (
	"encoding/json"

	"github.com/graingo/mconv/internal"
)

// ToJSONE converts any type to JSON string with error.
func ToJSONE(value interface{}) (string, error) {
	if value == nil {
		return "null", nil
	}

	// use json.Marshal to convert
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", internal.NewConversionError(value, "JSON", err)
	}

	return string(bytes), nil
}

// ToJSON converts any type to JSON string.
func ToJSON(value interface{}) string {
	result, _ := ToJSONE(value)
	return result
}

// FromJSONE converts JSON string to specified type with error.
func FromJSONE(jsonStr string, target interface{}) error {
	if jsonStr == "" {
		return nil
	}

	// use json.Unmarshal to convert
	err := json.Unmarshal([]byte(jsonStr), target)
	if err != nil {
		return internal.NewConversionError(jsonStr, "object", internal.ErrInvalidJSONFormat)
	}

	return nil
}

// FromJSON converts JSON string to specified type.
func FromJSON(jsonStr string, target interface{}) {
	_ = FromJSONE(jsonStr, target)
}

// ToMapFromJSONE converts JSON string to map[string]interface{} with error.
func ToMapFromJSONE(jsonStr string) (map[string]interface{}, error) {
	if jsonStr == "" {
		return nil, nil
	}

	// create a map
	result := make(map[string]interface{})

	// use json.Unmarshal to convert
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, internal.NewConversionError(jsonStr, "map", internal.ErrInvalidJSONFormat)
	}

	return result, nil
}

// ToMapFromJSON converts JSON string to map[string]interface{}.
func ToMapFromJSON(jsonStr string) map[string]interface{} {
	result, _ := ToMapFromJSONE(jsonStr)
	return result
}

// ToSliceFromJSONE converts JSON string to []interface{} with error.
func ToSliceFromJSONE(jsonStr string) ([]interface{}, error) {
	if jsonStr == "" {
		return nil, nil
	}

	// create a slice
	var result []interface{}

	// use json.Unmarshal to convert
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, internal.NewConversionError(jsonStr, "slice", internal.ErrInvalidJSONFormat)
	}

	return result, nil
}

// ToSliceFromJSON converts JSON string to []interface{}.
func ToSliceFromJSON(jsonStr string) []interface{} {
	result, _ := ToSliceFromJSONE(jsonStr)
	return result
}
