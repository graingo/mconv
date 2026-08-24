package basic

import (
	"math"
	"strconv"
	"strings"

	"github.com/graingo/mconv/internal"
)

// ToUint converts value to uint and returns zero when conversion fails.
func ToUint(value interface{}) uint {
	result, _ := ToUintE(value)
	return result
}

// ToUintE converts value to uint.
func ToUintE(value interface{}) (uint, error) {
	converted, err := toUnsignedInteger(value, strconv.IntSize, "uint")
	return uint(converted), err
}

// ToUint64 converts value to uint64 and returns zero when conversion fails.
func ToUint64(value interface{}) uint64 {
	result, _ := ToUint64E(value)
	return result
}

// ToUint64E converts value to uint64.
func ToUint64E(value interface{}) (uint64, error) {
	return toUnsignedInteger(value, 64, "uint64")
}

// ToUint32 converts value to uint32 and returns zero when conversion fails.
func ToUint32(value interface{}) uint32 {
	result, _ := ToUint32E(value)
	return result
}

// ToUint32E converts value to uint32.
func ToUint32E(value interface{}) (uint32, error) {
	converted, err := toUnsignedInteger(value, 32, "uint32")
	return uint32(converted), err
}

// ToUint16 converts value to uint16 and returns zero when conversion fails.
func ToUint16(value interface{}) uint16 {
	result, _ := ToUint16E(value)
	return result
}

// ToUint16E converts value to uint16.
func ToUint16E(value interface{}) (uint16, error) {
	converted, err := toUnsignedInteger(value, 16, "uint16")
	return uint16(converted), err
}

// ToUint8 converts value to uint8 and returns zero when conversion fails.
func ToUint8(value interface{}) uint8 {
	result, _ := ToUint8E(value)
	return result
}

// ToUint8E converts value to uint8.
func ToUint8E(value interface{}) (uint8, error) {
	converted, err := toUnsignedInteger(value, 8, "uint8")
	return uint8(converted), err
}

// toUnsignedInteger applies the shared conversion and range rules for every
// unsigned integer width. Negative, non-finite, and rounded upper boundaries
// are rejected consistently before a result is returned.
func toUnsignedInteger(value interface{}, bits int, target string) (uint64, error) {
	normalized, valid := normalizeInput(value)
	if !valid {
		return 0, nil
	}
	if converted, handled, err := checkedUnsignedNumber(normalized, bits, target); handled {
		return converted, err
	}

	var converted uint64
	switch typed := normalized.(type) {
	case uint:
		converted = uint64(typed)
	case uint64:
		converted = typed
	case uint32:
		converted = uint64(typed)
	case uint16:
		converted = uint64(typed)
	case uint8:
		converted = uint64(typed)
	case int:
		return unsignedFromInt(int64(typed), normalized, bits, target)
	case int64:
		return unsignedFromInt(typed, normalized, bits, target)
	case int32:
		return unsignedFromInt(int64(typed), normalized, bits, target)
	case int16:
		return unsignedFromInt(int64(typed), normalized, bits, target)
	case int8:
		return unsignedFromInt(int64(typed), normalized, bits, target)
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case string:
		parsed, err := strconv.ParseUint(strings.TrimSpace(typed), 0, bits)
		if err != nil {
			return 0, internal.NewConversionError(normalized, target, err)
		}
		return parsed, nil
	default:
		return 0, internal.NewConversionError(normalized, target, internal.ErrUnsupportedType)
	}

	if converted > unsignedMaximum(bits) {
		return 0, internal.NewConversionError(normalized, target, internal.ErrOverflow)
	}
	return converted, nil
}

func unsignedFromInt(value int64, original interface{}, bits int, target string) (uint64, error) {
	if value < 0 {
		return 0, internal.NewConversionError(original, target, internal.ErrOverflow)
	}
	converted := uint64(value)
	if converted > unsignedMaximum(bits) {
		return 0, internal.NewConversionError(original, target, internal.ErrOverflow)
	}
	return converted, nil
}

func unsignedMaximum(bits int) uint64 {
	if bits == 64 {
		return math.MaxUint64
	}
	return uint64(1)<<bits - 1
}
