package basic

import (
	"strconv"
	"strings"

	"github.com/graingo/mconv/internal"
)

// ToBool converts any type to bool.
func ToBool(value interface{}) bool {
	result, _ := ToBoolE(value)
	return result
}

// ToBoolE converts any type to bool with error.
func ToBoolE(value interface{}) (bool, error) {
	if value == nil {
		return false, nil
	}

	switch v := value.(type) {
	case bool:
		return v, nil
	case int:
		return v != 0, nil
	case int64:
		return v != 0, nil
	case int32:
		return v != 0, nil
	case int16:
		return v != 0, nil
	case int8:
		return v != 0, nil
	case uint:
		return v != 0, nil
	case uint64:
		return v != 0, nil
	case uint32:
		return v != 0, nil
	case uint16:
		return v != 0, nil
	case uint8:
		return v != 0, nil
	case float64:
		return v != 0, nil
	case float32:
		return v != 0, nil
	case complex64:
		return real(v) != 0 || imag(v) != 0, nil
	case complex128:
		return real(v) != 0 || imag(v) != 0, nil
	case string:
		s := strings.ToLower(v)
		if s == "true" || s == "yes" || s == "y" || s == "1" {
			return true, nil
		} else if s == "false" || s == "no" || s == "n" || s == "0" {
			return false, nil
		}
		b, err := strconv.ParseBool(s)
		if err != nil {
			return false, internal.NewConversionError(value, "bool", err)
		}
		return b, nil
	default:
		return false, internal.NewConversionError(value, "bool", internal.ErrUnsupportedType)
	}
}
