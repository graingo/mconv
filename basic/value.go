package basic

import "reflect"

// isNilInput reports whether value contains a nil pointer or interface.
func isNilInput(value interface{}) bool {
	if value == nil {
		return true
	}
	rv := reflect.ValueOf(value)
	switch rv.Kind() {
	case reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	default:
		return false
	}
}

// normalizeInput dereferences pointers and interfaces, then converts scalar
// values to their built-in representation. This gives every converter the
// same behavior for pointers and user-defined scalar types.
func normalizeInput(value interface{}) (interface{}, bool) {
	if value == nil {
		return nil, false
	}

	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return nil, false
		}
		rv = rv.Elem()
	}

	switch rv.Kind() {
	case reflect.Bool:
		return rv.Bool(), true
	case reflect.Int:
		return int(rv.Int()), true
	case reflect.Int8:
		return int8(rv.Int()), true
	case reflect.Int16:
		return int16(rv.Int()), true
	case reflect.Int32:
		return int32(rv.Int()), true
	case reflect.Int64:
		return rv.Int(), true
	case reflect.Uint:
		return uint(rv.Uint()), true
	case reflect.Uint8:
		return uint8(rv.Uint()), true
	case reflect.Uint16:
		return uint16(rv.Uint()), true
	case reflect.Uint32:
		return uint32(rv.Uint()), true
	case reflect.Uint64:
		return rv.Uint(), true
	case reflect.Uintptr:
		return rv.Uint(), true
	case reflect.Float32:
		return float32(rv.Float()), true
	case reflect.Float64:
		return rv.Float(), true
	case reflect.Complex64:
		return complex64(rv.Complex()), true
	case reflect.Complex128:
		return rv.Complex(), true
	case reflect.String:
		return rv.String(), true
	default:
		return rv.Interface(), true
	}
}
