package complex

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/graingo/mconv/basic"
	"github.com/graingo/mconv/internal"
)

// HookFunc is a function type for decoding hooks.
// It's used to add custom conversion logic to the Struct function.
// The function takes a source type, a destination type, and the data to be converted.
// If it can handle the conversion, it should return the converted value and nil error.
// If it cannot handle the conversion, it should return the original data and nil error.
// If an error occurs during conversion, it should return nil and the error.
type HookFunc func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error)

// ToStruct converts a map or a struct to a struct.
// The `pointer` parameter should be a pointer to a struct.
// It supports `mconv` tag for custom field mapping.
// It accepts optional HookFuncs to provide custom conversion logic.
func ToStruct(source, pointer interface{}, hooks ...HookFunc) {
	_ = ToStructE(source, pointer, hooks...)
}

// StructE is the same as Struct but returns an error.
func ToStructE(source, pointer interface{}, hooks ...HookFunc) error {
	if pointer == nil {
		return errors.New("pointer cannot be nil")
	}

	// Prepend default hooks
	allHooks := []HookFunc{stringToTimeHookFunc(), stringToDurationHookFunc(), intToBoolHookFunc()}
	allHooks = append(allHooks, hooks...)

	// Get the reflect.Value of the pointer and the struct
	pointerRv := reflect.ValueOf(pointer)
	if pointerRv.Kind() != reflect.Ptr {
		return fmt.Errorf("pointer must be a pointer to a struct, but got %T", pointer)
	}
	structRv := pointerRv.Elem()
	if structRv.Kind() != reflect.Struct {
		return fmt.Errorf("pointer must be a pointer to a struct, but got a pointer to %s", structRv.Kind())
	}

	// Get the decoder plan from cache or build a new one.
	decoder, err := getDecoder(structRv.Type(), allHooks...)
	if err != nil {
		return err
	}

	// Convert source to map[string]interface{}
	sourceMap, err := ToMapE(source)
	if err != nil {
		return fmt.Errorf("source data cannot be converted to a map: %w", err)
	}
	if sourceMap == nil {
		return nil
	}

	// For case-insensitive key matching, create a lookup map from lower-case key to original key.
	lowerCaseKeyMap := make(map[string]string, len(sourceMap))
	for k := range sourceMap {
		lowerCaseKeyMap[strings.ToLower(k)] = k
	}

	// Iterate over the fields in the decoder plan, not the struct fields directly.
	for _, fieldDecoder := range decoder.FieldArr {
		var (
			mapValue any
			ok       bool
		)

		// 1. Try case-sensitive match.
		mapValue, ok = sourceMap[fieldDecoder.Name]

		// 2. If not found, try case-insensitive match.
		if !ok {
			if originalKey, found := lowerCaseKeyMap[strings.ToLower(fieldDecoder.Name)]; found {
				mapValue = sourceMap[originalKey]
				ok = true
			}
		}

		if !ok {
			continue
		}

		// Get the field value by its cached index.
		fieldVal := structRv.FieldByIndex(fieldDecoder.Index)

		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			continue
		}

		// Set the field value.
		if err := setFieldValue(fieldVal, mapValue, allHooks...); err != nil {
			return fmt.Errorf("failed to set field '%s': %w", fieldDecoder.Field.Name, err)
		}
	}

	return nil
}

// setFieldValue sets a reflect.Value with an interface{} value, performing necessary type conversions.
func setFieldValue(field reflect.Value, value interface{}, hooks ...HookFunc) error {
	if !field.IsValid() {
		return errors.New("field is not valid")
	}

	if value == nil {
		return nil
	}

	// Apply hooks first.
	// Hooks can be chained, with the result of one hook being the input to the next.
	var (
		fromType = reflect.TypeOf(value)
		err      error
	)
	for _, hook := range hooks {
		value, err = hook(fromType, field.Type(), value)
		if err != nil {
			return fmt.Errorf("hook function error: %w", err)
		}
		if value != nil {
			fromType = reflect.TypeOf(value)
		} else {
			fromType = nil // Value is nil, so there's no type.
		}
	}

	// After hooks, value might become nil.
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
		return setFieldValue(field.Elem(), value, hooks...)
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
		return ToStructE(value, field.Addr().Interface(), hooks...)
	case reflect.Slice:
		sliceData, err := ToSliceE(value)
		if err != nil {
			return err
		}
		newSlice := reflect.MakeSlice(field.Type(), len(sliceData), len(sliceData))
		for i, v := range sliceData {
			elem := newSlice.Index(i)
			if err := setFieldValue(elem, v, hooks...); err != nil {
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
			if err := setFieldValue(newKey, k, hooks...); err != nil {
				return fmt.Errorf("failed to convert map key: %w", err)
			}

			newVal := reflect.New(valType).Elem()
			if err := setFieldValue(newVal, v, hooks...); err != nil {
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

// getDecoder retrieves a decoder for a given struct type from the cache.
// If the decoder is not found in the cache, it builds a new one, caches it, and returns it.
func getDecoder(destType reflect.Type, hooks ...HookFunc) (*internal.Decoder, error) {
	cacheKey := internal.DecoderCacheKey{DestType: destType, NumHooks: len(hooks)}
	if decoder, ok := internal.GetDecoder(cacheKey); ok {
		return decoder, nil
	}

	// Slow path: build a new decoder.
	decoder := &internal.Decoder{
		Fields:   make(map[string]*internal.FieldDecoder),
		FieldArr: make([]*internal.FieldDecoder, 0),
	}

	buildDecoderFields(destType, []int{}, decoder)

	// Cache the new decoder.
	internal.SetDecoder(cacheKey, decoder)
	return decoder, nil
}

// buildDecoderFields recursively traverses a struct type and populates the decoder with field information.
func buildDecoderFields(t reflect.Type, indexPrefix []int, decoder *internal.Decoder) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Recurse into anonymous embedded structs.
		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			buildDecoderFields(field.Type, append(append([]int(nil), indexPrefix...), i), decoder)
			continue
		}

		// Skip unexported fields.
		if isUnexportedField(field) {
			continue
		}

		// Parse the tag.
		tag := field.Tag.Get("mconv")
		if tag == "-" {
			continue
		}

		key := field.Name
		parts := strings.Split(tag, ",")
		if len(parts) > 0 && parts[0] != "" {
			key = parts[0]
		}

		// If a field with the same name already exists in a shallower layer, skip this one.
		if _, ok := decoder.Fields[key]; ok {
			continue
		}

		fieldDecoder := &internal.FieldDecoder{
			Field: field,
			Index: append(append([]int(nil), indexPrefix...), i), // Must be a copy
			Name:  key,
		}

		decoder.FieldArr = append(decoder.FieldArr, fieldDecoder)
		decoder.Fields[key] = fieldDecoder
	}
}

// isUnexportedField checks if a struct field is unexported.
func isUnexportedField(field reflect.StructField) bool {
	return field.PkgPath != ""
}
