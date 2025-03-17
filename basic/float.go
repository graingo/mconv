package basic

import (
	"strconv"

	"github.com/mingzaily/mconv/internal"
)

// ToFloat64 converts any type to float64.
func ToFloat64(value interface{}) float64 {
	result, _ := ToFloat64E(value)
	return result
}

// ToFloat64E converts any type to float64 with error.
func ToFloat64E(value interface{}) (float64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case complex64:
		return float64(real(v)), nil
	case complex128:
		return float64(real(v)), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, internal.NewConversionError(value, "float64", err)
		}
		return f, nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, internal.NewConversionError(value, "float64", internal.ErrUnsupportedType)
	}
}

// ToFloat32 converts any type to float32.
func ToFloat32(value interface{}) float32 {
	result, _ := ToFloat32E(value)
	return result
}

// ToFloat32E converts any type to float32 with error.
func ToFloat32E(value interface{}) (float32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case float64:
		return float32(v), nil
	case float32:
		return v, nil
	case int:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case complex64:
		return float32(real(v)), nil
	case complex128:
		return float32(real(v)), nil
	case string:
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return 0, internal.NewConversionError(value, "float32", err)
		}
		return float32(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return 0, internal.NewConversionError(value, "float32", internal.ErrUnsupportedType)
	}
}
