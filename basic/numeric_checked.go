package basic

import (
	"math"

	"github.com/graingo/mconv/internal"
)

func checkedSignedNumber(value interface{}, bits int, target string) (int64, bool, error) {
	var number float64
	switch typed := value.(type) {
	case float64:
		number = typed
	case float32:
		number = float64(typed)
	case complex64:
		if imag(typed) != 0 {
			return 0, true, internal.NewConversionError(value, target, internal.ErrConversionFailed)
		}
		number = float64(real(typed))
	case complex128:
		if imag(typed) != 0 {
			return 0, true, internal.NewConversionError(value, target, internal.ErrConversionFailed)
		}
		number = real(typed)
	default:
		return 0, false, nil
	}
	limit := math.Ldexp(1, bits-1)
	if math.IsNaN(number) || math.IsInf(number, 0) || number < -limit || number >= limit {
		return 0, true, internal.NewConversionError(value, target, internal.ErrOverflow)
	}
	return int64(number), true, nil
}

func checkedUnsignedNumber(value interface{}, bits int, target string) (uint64, bool, error) {
	var number float64
	switch typed := value.(type) {
	case float64:
		number = typed
	case float32:
		number = float64(typed)
	case complex64:
		if imag(typed) != 0 {
			return 0, true, internal.NewConversionError(value, target, internal.ErrConversionFailed)
		}
		number = float64(real(typed))
	case complex128:
		if imag(typed) != 0 {
			return 0, true, internal.NewConversionError(value, target, internal.ErrConversionFailed)
		}
		number = real(typed)
	default:
		return 0, false, nil
	}
	limit := math.Ldexp(1, bits)
	if math.IsNaN(number) || math.IsInf(number, 0) || number < 0 || number >= limit {
		return 0, true, internal.NewConversionError(value, target, internal.ErrOverflow)
	}
	return uint64(number), true, nil
}
