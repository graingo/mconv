package basic

import (
	"math"
	"strconv"
	"strings"

	"github.com/mingzaily/mconv/internal"
)

// ToUint converts any type to uint
func ToUint(value interface{}) uint {
	result, _ := ToUintE(value)
	return result
}

// ToUintE converts any type to uint with error
func ToUintE(value interface{}) (uint, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint:
		return v, nil
	case uint64:
		if v > uint64(^uint(0)) {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	case int:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case int32:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case int16:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case int8:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case float64:
		if v < 0 || v > float64(^uint(0)) {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(v), nil
	case float32:
		f := float64(v)
		if f < 0 || f > float64(^uint(0)) {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(f), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrConversionFailed)
		}
		f := float64(real(v))
		if f < 0 || f > float64(^uint(0)) {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint", internal.ErrConversionFailed)
		}
		f := real(v)
		if f < 0 || f > float64(^uint(0)) {
			return 0, internal.NewConversionError(value, "uint", internal.ErrOverflow)
		}
		return uint(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		u, err := strconv.ParseUint(v, 0, strconv.IntSize)
		if err != nil {
			return 0, internal.NewConversionError(value, "uint", err)
		}
		return uint(u), nil
	default:
		return 0, internal.NewConversionError(value, "uint", internal.ErrUnsupportedType)
	}
}

// ToUint64 converts any type to uint64
func ToUint64(value interface{}) uint64 {
	result, _ := ToUint64E(value)
	return result
}

// ToUint64E converts any type to uint64 with error
func ToUint64E(value interface{}) (uint64, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint:
		return uint64(v), nil
	case uint64:
		return v, nil
	case uint32:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case int:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(v), nil
	case float64:
		if v < 0 || v > float64(math.MaxUint64) {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(v), nil
	case float32:
		f := float64(v)
		if f < 0 || f > float64(math.MaxUint64) {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(f), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrConversionFailed)
		}
		f := float64(real(v))
		if f < 0 || f > float64(math.MaxUint64) {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrConversionFailed)
		}
		f := real(v)
		if f < 0 || f > float64(math.MaxUint64) {
			return 0, internal.NewConversionError(value, "uint64", internal.ErrOverflow)
		}
		return uint64(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		u, err := strconv.ParseUint(v, 0, 64)
		if err != nil {
			return 0, internal.NewConversionError(value, "uint64", err)
		}
		return u, nil
	default:
		return 0, internal.NewConversionError(value, "uint64", internal.ErrUnsupportedType)
	}
}

// ToUint32 converts any type to uint32
func ToUint32(value interface{}) uint32 {
	result, _ := ToUint32E(value)
	return result
}

// ToUint32E converts any type to uint32 with error
func ToUint32E(value interface{}) (uint32, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint:
		if v > uint(^uint32(0)) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case uint64:
		if v > uint64(^uint32(0)) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint16:
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case int:
		if v < 0 || v > int(^uint32(0)) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case int64:
		if v < 0 || v > int64(^uint32(0)) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case float64:
		if v < 0 || v > float64(math.MaxUint32) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(v), nil
	case float32:
		f := float64(v)
		if f < 0 || f > float64(math.MaxUint32) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(f), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrConversionFailed)
		}
		f := float64(real(v))
		if f < 0 || f > float64(math.MaxUint32) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrConversionFailed)
		}
		f := real(v)
		if f < 0 || f > float64(math.MaxUint32) {
			return 0, internal.NewConversionError(value, "uint32", internal.ErrOverflow)
		}
		return uint32(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		u, err := strconv.ParseUint(v, 0, 32)
		if err != nil {
			return 0, internal.NewConversionError(value, "uint32", err)
		}
		return uint32(u), nil
	default:
		return 0, internal.NewConversionError(value, "uint32", internal.ErrUnsupportedType)
	}
}

// ToUint16 converts any type to uint16
func ToUint16(value interface{}) uint16 {
	result, _ := ToUint16E(value)
	return result
}

// ToUint16E converts any type to uint16 with error
func ToUint16E(value interface{}) (uint16, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint:
		if v > uint(^uint16(0)) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case uint64:
		if v > uint64(^uint16(0)) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case uint32:
		if v > uint32(^uint16(0)) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint8:
		return uint16(v), nil
	case int:
		if v < 0 || v > int(^uint16(0)) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case int64:
		if v < 0 || v > int64(^uint16(0)) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case int32:
		if v < 0 || v > int32(^uint16(0)) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case float64:
		if v < 0 || v > float64(math.MaxUint16) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(v), nil
	case float32:
		f := float64(v)
		if f < 0 || f > float64(math.MaxUint16) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(f), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrConversionFailed)
		}
		f := float64(real(v))
		if f < 0 || f > float64(math.MaxUint16) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrConversionFailed)
		}
		f := real(v)
		if f < 0 || f > float64(math.MaxUint16) {
			return 0, internal.NewConversionError(value, "uint16", internal.ErrOverflow)
		}
		return uint16(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		u, err := strconv.ParseUint(v, 0, 16)
		if err != nil {
			return 0, internal.NewConversionError(value, "uint16", err)
		}
		return uint16(u), nil
	default:
		return 0, internal.NewConversionError(value, "uint16", internal.ErrUnsupportedType)
	}
}

// ToUint8 converts any type to uint8
func ToUint8(value interface{}) uint8 {
	result, _ := ToUint8E(value)
	return result
}

// ToUint8E converts any type to uint8 with error
func ToUint8E(value interface{}) (uint8, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case uint:
		if v > uint(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case uint64:
		if v > uint64(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case uint32:
		if v > uint32(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case uint16:
		if v > uint16(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case uint8:
		return v, nil
	case int:
		if v < 0 || v > int(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case int64:
		if v < 0 || v > int64(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case int32:
		if v < 0 || v > int32(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case int16:
		if v < 0 || v > int16(^uint8(0)) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case float64:
		if v < 0 || v > float64(math.MaxUint8) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(v), nil
	case float32:
		f := float64(v)
		if f < 0 || f > float64(math.MaxUint8) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(f), nil
	case complex64:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrConversionFailed)
		}
		f := float64(real(v))
		if f < 0 || f > float64(math.MaxUint8) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(f), nil
	case complex128:
		if imag(v) != 0 {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrConversionFailed)
		}
		f := real(v)
		if f < 0 || f > float64(math.MaxUint8) {
			return 0, internal.NewConversionError(value, "uint8", internal.ErrOverflow)
		}
		return uint8(f), nil
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case string:
		v = strings.TrimSpace(v)
		u, err := strconv.ParseUint(v, 0, 8)
		if err != nil {
			return 0, internal.NewConversionError(value, "uint8", err)
		}
		return uint8(u), nil
	default:
		return 0, internal.NewConversionError(value, "uint8", internal.ErrUnsupportedType)
	}
}
