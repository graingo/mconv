package complex

import (
	"reflect"
	"time"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/internal"
)

// ToT converts value to T and returns the zero value when conversion fails.
func ToT[T any](value interface{}, hooks ...HookFunc) T {
	result, err := ToTE[T](value, hooks...)
	if err != nil {
		var zero T
		return zero
	}
	return result
}

// ToTE converts value to T and returns a conversion error with path context.
func ToTE[T any](value interface{}, hooks ...HookFunc) (T, error) {
	if len(hooks) == 0 {
		return convertGenericValue(value, genericConverter[T]())
	}

	var result T
	allHooks := make([]HookFunc, 0, len(defaultHooks)+len(hooks))
	allHooks = append(allHooks, defaultHooks...)
	allHooks = append(allHooks, hooks...)
	err := setFieldValue(reflect.ValueOf(&result).Elem(), value, allHooks...)
	return result, err
}

// genericConverter selects the target-specific conversion once per container.
func genericConverter[T any]() func(interface{}) (T, error) {
	var target T
	switch any(&target).(type) {
	case *string:
		return typedConverter[T](basic.ToStringE)
	case *bool:
		return typedConverter[T](basic.ToBoolE)
	case *int:
		return typedConverter[T](basic.ToIntE)
	case *int8:
		return typedConverter[T](basic.ToInt8E)
	case *int16:
		return typedConverter[T](basic.ToInt16E)
	case *int32:
		return typedConverter[T](basic.ToInt32E)
	case *int64:
		return typedConverter[T](basic.ToInt64E)
	case *uint:
		return typedConverter[T](basic.ToUintE)
	case *uint8:
		return typedConverter[T](basic.ToUint8E)
	case *uint16:
		return typedConverter[T](basic.ToUint16E)
	case *uint32:
		return typedConverter[T](basic.ToUint32E)
	case *uint64:
		return typedConverter[T](basic.ToUint64E)
	case *uintptr:
		return func(value interface{}) (T, error) {
			converted, err := basic.ToUint64E(value)
			if err == nil && uint64(uintptr(converted)) != converted {
				return target, internal.NewConversionError(value, "uintptr", internal.ErrOverflow)
			}
			return any(uintptr(converted)).(T), err
		}
	case *float32:
		return typedConverter[T](basic.ToFloat32E)
	case *float64:
		return typedConverter[T](basic.ToFloat64E)
	case *complex64:
		return typedConverter[T](basic.ToComplex64E)
	case *complex128:
		return typedConverter[T](basic.ToComplex128E)
	case *time.Time:
		return func(value interface{}) (T, error) {
			converted, err := basic.ToTimeE(value)
			return any(converted).(T), err
		}
	case *time.Duration:
		return typedConverter[T](basic.ToDurationE)
	default:
		return func(value interface{}) (T, error) {
			var result T
			err := setGenericValue(reflect.ValueOf(&result).Elem(), value)
			return result, err
		}
	}
}

func typedConverter[T, U any](convert func(interface{}) (U, error)) func(interface{}) (T, error) {
	return func(value interface{}) (T, error) {
		converted, err := convert(value)
		return any(converted).(T), err
	}
}

func convertGenericValue[T any](value interface{}, convert func(interface{}) (T, error)) (T, error) {
	var zero T
	if value == nil {
		return zero, nil
	}
	if converted, ok := value.(T); ok {
		return converted, nil
	}
	return convert(value)
}
