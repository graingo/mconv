package basic

import (
	"math"
	"strconv"
	"strings"

	"github.com/graingo/mconv/internal"
)

// ToInt converts value to int and returns zero when conversion fails.
func ToInt(value interface{}) int {
	result, _ := ToIntE(value)
	return result
}

// ToIntE converts value to int.
func ToIntE(value interface{}) (int, error) {
	converted, err := toSignedInteger(value, strconv.IntSize, "int")
	return int(converted), err
}

// ToInt64 converts value to int64 and returns zero when conversion fails.
func ToInt64(value interface{}) int64 {
	result, _ := ToInt64E(value)
	return result
}

// ToInt64E converts value to int64.
func ToInt64E(value interface{}) (int64, error) {
	return toSignedInteger(value, 64, "int64")
}

// ToInt32 converts value to int32 and returns zero when conversion fails.
func ToInt32(value interface{}) int32 {
	result, _ := ToInt32E(value)
	return result
}

// ToInt32E converts value to int32.
func ToInt32E(value interface{}) (int32, error) {
	converted, err := toSignedInteger(value, 32, "int32")
	return int32(converted), err
}

// ToInt16 converts value to int16 and returns zero when conversion fails.
func ToInt16(value interface{}) int16 {
	result, _ := ToInt16E(value)
	return result
}

// ToInt16E converts value to int16.
func ToInt16E(value interface{}) (int16, error) {
	converted, err := toSignedInteger(value, 16, "int16")
	return int16(converted), err
}

// ToInt8 converts value to int8 and returns zero when conversion fails.
func ToInt8(value interface{}) int8 {
	result, _ := ToInt8E(value)
	return result
}

// ToInt8E converts value to int8.
func ToInt8E(value interface{}) (int8, error) {
	converted, err := toSignedInteger(value, 8, "int8")
	return int8(converted), err
}

// toSignedInteger applies the shared conversion and range rules for every
// signed integer width. Floating-point and complex inputs are validated before
// the exact integer cases so non-finite values and rounded boundaries fail.
func toSignedInteger(value interface{}, bits int, target string) (int64, error) {
	normalized, valid := normalizeInput(value)
	if !valid {
		return 0, nil
	}
	if converted, handled, err := checkedSignedNumber(normalized, bits, target); handled {
		return converted, err
	}

	var converted int64
	switch typed := normalized.(type) {
	case int:
		converted = int64(typed)
	case int64:
		converted = typed
	case int32:
		converted = int64(typed)
	case int16:
		converted = int64(typed)
	case int8:
		converted = int64(typed)
	case uint:
		return signedFromUint(uint64(typed), normalized, bits, target)
	case uint64:
		return signedFromUint(typed, normalized, bits, target)
	case uint32:
		return signedFromUint(uint64(typed), normalized, bits, target)
	case uint16:
		return signedFromUint(uint64(typed), normalized, bits, target)
	case uint8:
		return signedFromUint(uint64(typed), normalized, bits, target)
	case bool:
		if typed {
			return 1, nil
		}
		return 0, nil
	case string:
		parsed, err := strconv.ParseInt(strings.TrimSpace(typed), 0, bits)
		if err != nil {
			return 0, internal.NewConversionError(normalized, target, err)
		}
		return parsed, nil
	default:
		return 0, internal.NewConversionError(normalized, target, internal.ErrUnsupportedType)
	}

	minimum, maximum := signedBounds(bits)
	if converted < minimum || converted > maximum {
		return 0, internal.NewConversionError(normalized, target, internal.ErrOverflow)
	}
	return converted, nil
}

func signedFromUint(value uint64, original interface{}, bits int, target string) (int64, error) {
	_, maximum := signedBounds(bits)
	if value > uint64(maximum) {
		return 0, internal.NewConversionError(original, target, internal.ErrOverflow)
	}
	return int64(value), nil
}

func signedBounds(bits int) (int64, int64) {
	if bits == 64 {
		return math.MinInt64, math.MaxInt64
	}
	maximum := int64(1)<<(bits-1) - 1
	return -maximum - 1, maximum
}
