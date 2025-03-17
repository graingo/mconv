package basic

import (
	"strconv"

	"github.com/mingzaily/mconv/internal"
)

// ToComplex128E converts any type to complex128 with error.
func ToComplex128E(value interface{}) (complex128, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case complex128:
		return v, nil
	case complex64:
		return complex128(v), nil
	case int:
		return complex(float64(v), 0), nil
	case int64:
		return complex(float64(v), 0), nil
	case int32:
		return complex(float64(v), 0), nil
	case int16:
		return complex(float64(v), 0), nil
	case int8:
		return complex(float64(v), 0), nil
	case uint:
		return complex(float64(v), 0), nil
	case uint64:
		return complex(float64(v), 0), nil
	case uint32:
		return complex(float64(v), 0), nil
	case uint16:
		return complex(float64(v), 0), nil
	case uint8:
		return complex(float64(v), 0), nil
	case float64:
		return complex(v, 0), nil
	case float32:
		return complex(float64(v), 0), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		// Try to parse complex number
		c, err := strconv.ParseComplex(v, 128)
		if err != nil {
			// Try to parse as float
			f, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0, internal.NewConversionError(value, "complex128", err)
			}
			return complex(f, 0), nil
		}
		return c, nil
	default:
		return 0, internal.NewConversionError(value, "complex128", internal.ErrUnsupportedType)
	}
}

// ToComplex64E converts any type to complex64 with error.
func ToComplex64E(value interface{}) (complex64, error) {
	c, err := ToComplex128E(value)
	if err != nil {
		return 0, err
	}

	// check overflow
	if real(c) > float64(float32(real(c))) || imag(c) > float64(float32(imag(c))) {
		return 0, internal.NewConversionError(value, "complex64", internal.ErrOverflow)
	}

	return complex64(c), nil
}

// ToComplex128 converts any type to complex128.
func ToComplex128(value interface{}) complex128 {
	result, _ := ToComplex128E(value)
	return result
}

// ToComplex64 converts any type to complex64.
func ToComplex64(value interface{}) complex64 {
	result, _ := ToComplex64E(value)
	return result
}
