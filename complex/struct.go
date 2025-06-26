package complex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/internal"
)

// Struct converts a map or a struct to a struct.
// The `pointer` parameter should be a pointer to a struct.
// It supports `mconv` tag for custom field mapping.
func Struct(source, pointer interface{}) {
	_ = StructE(source, pointer)
}

// StructE is the same as Struct but returns an error.
func StructE(source, pointer interface{}) error {
	if pointer == nil {
		return errors.New("pointer cannot be nil")
	}

	// Get the reflect.Value of the pointer
	pointerRv := reflect.ValueOf(pointer)
	if pointerRv.Kind() != reflect.Ptr {
		return fmt.Errorf("pointer must be a pointer to a struct, but got %T", pointer)
	}

	// Get the reflect.Value of the struct
	structRv := pointerRv.Elem()
	if structRv.Kind() != reflect.Struct {
		return fmt.Errorf("pointer must be a pointer to a struct, but got a pointer to %s", structRv.Kind())
	}

	// Convert source to map[string]interface{}
	sourceMap, err := ToMapE(source)
	if err != nil {
		// If the source is not a map-like structure, we can't proceed.
		// However, gconv can handle struct-to-struct conversion directly.
		// For now, we'll stick to map-based conversion.
		return fmt.Errorf("source data cannot be converted to a map: %w", err)
	}

	if sourceMap == nil {
		return nil
	}

	// Iterate over struct fields
	for i := 0; i < structRv.NumField(); i++ {
		field := structRv.Field(i)
		fieldType := structRv.Type().Field(i)

		if !field.CanSet() {
			continue
		}

		// Get map key from tag or field name
		key := fieldType.Name
		tag := fieldType.Tag.Get("mconv")
		if tag != "" {
			if tag == "-" {
				continue
			}
			// honor ,omitempty
			parts := strings.Split(tag, ",")
			key = parts[0]
		}

		// Find value in map
		mapValue, ok := sourceMap[key]
		if !ok {
			// Case-insensitive key search
			for k, v := range sourceMap {
				if strings.EqualFold(k, key) {
					mapValue = v
					ok = true
					break
				}
			}
		}

		if !ok {
			continue
		}

		// Set field value
		if err := setFieldValue(field, mapValue); err != nil {
			return fmt.Errorf("failed to set field '%s': %w", fieldType.Name, err)
		}
	}

	return nil
}

// setFieldValue sets a reflect.Value with an interface{} value, performing necessary type conversions.
func setFieldValue(field reflect.Value, value interface{}) error {
	if !field.IsValid() {
		return errors.New("field is not valid")
	}

	if value == nil {
		return nil
	}

	valueRv := reflect.ValueOf(value)

	// If types are directly assignable
	if valueRv.IsValid() && valueRv.Type().AssignableTo(field.Type()) {
		field.Set(valueRv)
		return nil
	}

	// Handle pointer fields
	if field.Kind() == reflect.Ptr {
		if !valueRv.IsValid() {
			return nil // Don't set nil to a pointer field
		}
		// Create a new instance for the pointer if the field is nil
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		// Set the value of the element pointed to
		return setFieldValue(field.Elem(), value)
	}

	switch field.Kind() {
	case reflect.String:
		s, err := basic.ToStringE(value)
		if err != nil {
			return err
		}
		field.SetString(s)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := basic.ToInt64E(value)
		if err != nil {
			return err
		}
		if field.OverflowInt(i) {
			return fmt.Errorf("value %v overflows field of type %s", value, field.Type())
		}
		field.SetInt(i)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		u, err := basic.ToUint64E(value)
		if err != nil {
			return err
		}
		if field.OverflowUint(u) {
			return fmt.Errorf("value %v overflows field of type %s", value, field.Type())
		}
		field.SetUint(u)
	case reflect.Float32, reflect.Float64:
		f, err := basic.ToFloat64E(value)
		if err != nil {
			return err
		}
		if field.OverflowFloat(f) {
			return fmt.Errorf("value %v overflows field of type %s", value, field.Type())
		}
		field.SetFloat(f)
	case reflect.Bool:
		b, err := basic.ToBoolE(value)
		if err != nil {
			return err
		}
		field.SetBool(b)
	case reflect.Struct:
		// Recursive call for nested structs
		return StructE(value, field.Addr().Interface())
	case reflect.Slice:
		sliceData, err := ToSliceE(value)
		if err != nil {
			return err
		}
		newSlice := reflect.MakeSlice(field.Type(), len(sliceData), len(sliceData))
		for i, v := range sliceData {
			elem := newSlice.Index(i)
			if err := setFieldValue(elem, v); err != nil {
				return err
			}
		}
		field.Set(newSlice)
	case reflect.Map:
		mapData, err := ToMapE(value)
		if err != nil {
			return err
		}
		keyType := field.Type().Key()
		valType := field.Type().Elem()
		newMap := reflect.MakeMap(field.Type())
		for k, v := range mapData {
			newKey := reflect.New(keyType).Elem()
			if err := setFieldValue(newKey, k); err != nil {
				return fmt.Errorf("failed to convert map key: %w", err)
			}

			newVal := reflect.New(valType).Elem()
			if err := setFieldValue(newVal, v); err != nil {
				return fmt.Errorf("failed to convert map value for key '%s': %w", k, err)
			}
			newMap.SetMapIndex(newKey, newVal)
		}
		field.Set(newMap)
	default:
		// Try a final conversion attempt
		if valueRv.IsValid() && valueRv.Type().ConvertibleTo(field.Type()) {
			field.Set(valueRv.Convert(field.Type()))
			return nil
		}
		return internal.NewConversionError(value, field.Type().String(), internal.ErrUnsupportedType)
	}
	return nil
}

// Scan is an alias for Struct.
func Scan(source, pointer interface{}) error {
	return StructE(source, pointer)
}
