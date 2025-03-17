package basic

import (
	"strconv"
	"strings"

	"github.com/mingzaily/mconv/internal"
)

// ToInt converts any type to int.
func ToInt(value interface{}) int {
	result, _ := ToIntE(value)
	return result
}

// ToIntE converts any type to int with error.
func ToIntE(value interface{}) (int, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case int64:
		if v > int64(^uint(0)>>1) || v < -int64(^uint(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(v), nil
	case int32:
		return int(v), nil
	case int16:
		return int(v), nil
	case int8:
		return int(v), nil
	case uint:
		if v > uint(^uint(0)>>1) {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(v), nil
	case uint64:
		if v > uint64(^uint(0)>>1) {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(v), nil
	case uint32:
		if strconv.IntSize == 32 && v > 2147483647 {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint8:
		return int(v), nil
	case float64:
		if v > float64(^uint(0)>>1) || v < -float64(^uint(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(v), nil
	case float32:
		if v > float32(^uint(0)>>1) || v < -float32(^uint(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(v), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int", internal.ErrConversionFailed)
		}
		f := float32(real(v))
		if f > float32(^uint(0)>>1) || f < -float32(^uint(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int", internal.ErrConversionFailed)
		}
		f := real(v)
		if f > float64(^uint(0)>>1) || f < -float64(^uint(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int", internal.ErrOverflow)
		}
		return int(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, strconv.IntSize)
		if err != nil {
			return 0, internal.NewConversionError(value, "int", err)
		}
		return int(i), nil
	default:
		return 0, internal.NewConversionError(value, "int", internal.ErrUnsupportedType)
	}
}

// ToInt64 converts any type to int64.
func ToInt64(value interface{}) int64 {
	result, _ := ToInt64E(value)
	return result
}

// ToInt64E converts any type to int64 with error.
func ToInt64E(value interface{}) (int64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint64:
		if v > uint64(^uint64(0)>>1) {
			return 0, internal.NewConversionError(value, "int64", internal.ErrOverflow)
		}
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case float64:
		if v > float64(^uint64(0)>>1) || v < -float64(^uint64(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int64", internal.ErrOverflow)
		}
		return int64(v), nil
	case float32:
		if v > float32(^uint64(0)>>1) || v < -float32(^uint64(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int64", internal.ErrOverflow)
		}
		return int64(v), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int64", internal.ErrConversionFailed)
		}
		f := float32(real(v))
		if f > float32(^uint64(0)>>1) || f < -float32(^uint64(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int64", internal.ErrOverflow)
		}
		return int64(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int64", internal.ErrConversionFailed)
		}
		f := real(v)
		if f > float64(^uint64(0)>>1) || f < -float64(^uint64(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int64", internal.ErrOverflow)
		}
		return int64(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		i, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			return 0, internal.NewConversionError(value, "int64", err)
		}
		return i, nil
	default:
		return 0, internal.NewConversionError(value, "int64", internal.ErrUnsupportedType)
	}
}

// ToInt32 converts any type to int32
func ToInt32(value interface{}) int32 {
	result, _ := ToInt32E(value)
	return result
}

// ToInt32E converts any type to int32 with error
func ToInt32E(value interface{}) (int32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		if v > int(^uint32(0)>>1) || v < -int(^uint32(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case int64:
		if v > int64(^uint32(0)>>1) || v < -int64(^uint32(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case int32:
		return v, nil
	case int16:
		return int32(v), nil
	case int8:
		return int32(v), nil
	case uint:
		if v > uint(^uint32(0)>>1) {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case uint64:
		if v > uint64(^uint32(0)>>1) {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case uint32:
		if v > uint32(^uint32(0)>>1) {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case float64:
		if v > float64(^uint32(0)>>1) || v < -float64(^uint32(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case float32:
		if v > float32(^uint32(0)>>1) || v < -float32(^uint32(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(v), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrConversionFailed)
		}
		f := float32(real(v))
		if f > float32(^uint32(0)>>1) || f < -float32(^uint32(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrConversionFailed)
		}
		f := real(v)
		if f > float64(^uint32(0)>>1) || f < -float64(^uint32(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int32", internal.ErrOverflow)
		}
		return int32(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		i, err := strconv.ParseInt(v, 0, 32)
		if err != nil {
			return 0, internal.NewConversionError(value, "int32", err)
		}
		return int32(i), nil
	default:
		return 0, internal.NewConversionError(value, "int32", internal.ErrUnsupportedType)
	}
}

// ToInt16 converts any type to int16
func ToInt16(value interface{}) int16 {
	result, _ := ToInt16E(value)
	return result
}

// ToInt16E converts any type to int16 with error
func ToInt16E(value interface{}) (int16, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		if v > int(^uint16(0)>>1) || v < -int(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case int64:
		if v > int64(^uint16(0)>>1) || v < -int64(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case int32:
		if v > int32(^uint16(0)>>1) || v < -int32(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case int16:
		return v, nil
	case int8:
		return int16(v), nil
	case uint:
		if v > uint(^uint16(0)>>1) {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case uint64:
		if v > uint64(^uint16(0)>>1) {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case uint32:
		if v > uint32(^uint16(0)>>1) {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case uint16:
		if v > uint16(^uint16(0)>>1) {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case float64:
		if v > float64(^uint16(0)>>1) || v < -float64(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case float32:
		if v > float32(^uint16(0)>>1) || v < -float32(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(v), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrConversionFailed)
		}
		f := float32(real(v))
		if f > float32(^uint16(0)>>1) || f < -float32(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrConversionFailed)
		}
		f := real(v)
		if f > float64(^uint16(0)>>1) || f < -float64(^uint16(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int16", internal.ErrOverflow)
		}
		return int16(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		i, err := strconv.ParseInt(v, 0, 16)
		if err != nil {
			return 0, internal.NewConversionError(value, "int16", err)
		}
		return int16(i), nil
	default:
		return 0, internal.NewConversionError(value, "int16", internal.ErrUnsupportedType)
	}
}

// ToInt8 converts any type to int8
func ToInt8(value interface{}) int8 {
	result, _ := ToInt8E(value)
	return result
}

// ToInt8E converts any type to int8 with error
func ToInt8E(value interface{}) (int8, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case int:
		if v > int(^uint8(0)>>1) || v < -int(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case int64:
		if v > int64(^uint8(0)>>1) || v < -int64(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case int32:
		if v > int32(^uint8(0)>>1) || v < -int32(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case int16:
		if v > int16(^uint8(0)>>1) || v < -int16(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case int8:
		return v, nil
	case uint:
		if v > uint(^uint8(0)>>1) {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case uint64:
		if v > uint64(^uint8(0)>>1) {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case uint32:
		if v > uint32(^uint8(0)>>1) {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case uint16:
		if v > uint16(^uint8(0)>>1) {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case uint8:
		if v > uint8(^uint8(0)>>1) {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case float64:
		if v > float64(^uint8(0)>>1) || v < -float64(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case float32:
		if v > float32(^uint8(0)>>1) || v < -float32(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(v), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrConversionFailed)
		}
		f := float32(real(v))
		if f > float32(^uint8(0)>>1) || f < -float32(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrConversionFailed)
		}
		f := real(v)
		if f > float64(^uint8(0)>>1) || f < -float64(^uint8(0)>>1)-1 {
			return 0, internal.NewConversionError(value, "int8", internal.ErrOverflow)
		}
		return int8(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		i, err := strconv.ParseInt(v, 0, 8)
		if err != nil {
			return 0, internal.NewConversionError(value, "int8", err)
		}
		return int8(i), nil
	default:
		return 0, internal.NewConversionError(value, "int8", internal.ErrUnsupportedType)
	}
}
