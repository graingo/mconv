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

// ToStructE converts source into pointer and reports conversion errors.
func ToStructE(source, pointer interface{}, hooks ...HookFunc) error {
	if pointer == nil {
		return errors.New("pointer cannot be nil")
	}

	pointerRv := reflect.ValueOf(pointer)
	if pointerRv.Kind() != reflect.Ptr || pointerRv.IsNil() {
		return fmt.Errorf("pointer must be a pointer to a struct, but got %T", pointer)
	}
	structRv := pointerRv.Elem()
	if structRv.Kind() != reflect.Struct {
		return fmt.Errorf("pointer must be a pointer to a struct, but got a pointer to %s", structRv.Kind())
	}

	allHooks := make([]HookFunc, 0, len(defaultHooks)+len(hooks))
	allHooks = append(allHooks, defaultHooks...)
	allHooks = append(allHooks, hooks...)

	converted, err := decodeStruct(source, structRv, allHooks)
	if err != nil {
		return err
	}
	structRv.Set(converted)
	return nil
}

// decodeStruct converts source into a new value based on current. The caller's
// value remains unchanged until the returned value is committed.
func decodeStruct(source interface{}, current reflect.Value, hooks []HookFunc) (reflect.Value, error) {
	decoder := getDecoder(current.Type())
	sourceMap, err := ToMapE(source)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("source data cannot be converted to a map: %w", err)
	}
	if sourceMap == nil {
		result := reflect.New(current.Type()).Elem()
		result.Set(current)
		return result, nil
	}

	lowerCaseKeyMap := make(map[string]string, len(sourceMap))
	ambiguousKeys := make(map[string]struct{})
	for k := range sourceMap {
		lowerKey := strings.ToLower(k)
		if existing, ok := lowerCaseKeyMap[lowerKey]; ok && existing != k {
			ambiguousKeys[lowerKey] = struct{}{}
			continue
		}
		lowerCaseKeyMap[lowerKey] = k
	}

	foldedDestinationCount := make(map[string]int, len(decoder.FieldArr))
	result := reflect.New(current.Type()).Elem()
	result.Set(current)

	for _, fieldDecoder := range decoder.FieldArr {
		foldedDestinationCount[strings.ToLower(fieldDecoder.Name)]++
	}

	for _, fieldDecoder := range decoder.FieldArr {
		var (
			mapValue any
			ok       bool
		)

		mapValue, ok = sourceMap[fieldDecoder.Name]

		if !ok {
			lowerKey := strings.ToLower(fieldDecoder.Name)
			if _, ambiguous := ambiguousKeys[lowerKey]; ambiguous {
				return reflect.Value{}, fmt.Errorf("source contains ambiguous keys for field '%s'", fieldDecoder.Field.Name)
			}
			if originalKey, found := lowerCaseKeyMap[lowerKey]; found {
				if _, exactDestination := decoder.Fields[originalKey]; exactDestination {
					continue
				}
				if foldedDestinationCount[lowerKey] > 1 {
					return reflect.Value{}, fmt.Errorf("destination contains ambiguous fields for source key %q", originalKey)
				}
				mapValue = sourceMap[originalKey]
				ok = true
			}
		}

		if !ok {
			continue
		}
		if fieldDecoder.Ambiguous {
			return reflect.Value{}, fmt.Errorf("destination contains ambiguous fields for source key %q", fieldDecoder.Name)
		}
		if mapValue == nil {
			continue
		}

		fieldVal, err := fieldByIndexCopyAlloc(result, fieldDecoder.Index)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("failed to access field '%s': %w", fieldDecoder.Field.Name, err)
		}

		if err := setFieldValue(fieldVal, mapValue, hooks...); err != nil {
			return reflect.Value{}, internal.PrependConversionPath(err, fieldDecoder.Name)
		}
	}
	return result, nil
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
		if hook == nil || fromType == nil {
			continue
		}
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
		converted := reflect.New(field.Type().Elem())
		if err := setFieldValue(converted.Elem(), value, hooks...); err != nil {
			return err
		}
		field.Set(converted)
		return nil
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
		converted, err := decodeStruct(value, field, hooks)
		if err != nil {
			return err
		}
		field.Set(converted)
		return nil
	case reflect.Slice:
		sliceData, err := ToSliceE(value)
		if err != nil {
			return err
		}
		newSlice := reflect.MakeSlice(field.Type(), len(sliceData), len(sliceData))
		for i, v := range sliceData {
			elem := newSlice.Index(i)
			if err := setFieldValue(elem, v, hooks...); err != nil {
				return internal.PrependConversionPath(err, fmt.Sprintf("[%d]", i))
			}
		}
		field.Set(newSlice)
	case reflect.Map:
		return setMapValue(field, value, hooks)
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

func setMapValue(field reflect.Value, value interface{}, hooks []HookFunc) error {
	source, valid := indirectValue(value)
	if !valid {
		field.Set(reflect.Zero(field.Type()))
		return nil
	}
	if source.Kind() == reflect.Struct {
		converted, err := ToMapE(source.Interface())
		if err != nil {
			return err
		}
		source = reflect.ValueOf(converted)
	}
	if source.Kind() != reflect.Map {
		return internal.NewConversionError(value, field.Type().String(), internal.ErrUnsupportedType)
	}
	if source.IsNil() {
		field.Set(reflect.Zero(field.Type()))
		return nil
	}

	newMap := reflect.MakeMapWithSize(field.Type(), source.Len())
	iterator := source.MapRange()
	for iterator.Next() {
		sourceKey := iterator.Key().Interface()
		path := fmt.Sprintf("[%v]", sourceKey)

		newKey := reflect.New(field.Type().Key()).Elem()
		if err := setFieldValue(newKey, sourceKey, hooks...); err != nil {
			return internal.PrependConversionPath(err, path)
		}
		if newMap.MapIndex(newKey).IsValid() {
			collisionErr := internal.NewConversionError(sourceKey, field.Type().Key().String(), internal.ErrConversionFailed)
			return internal.PrependConversionPath(collisionErr, path)
		}

		newValue := reflect.New(field.Type().Elem()).Elem()
		if err := setFieldValue(newValue, iterator.Value().Interface(), hooks...); err != nil {
			return internal.PrependConversionPath(err, path)
		}
		newMap.SetMapIndex(newKey, newValue)
	}
	field.Set(newMap)
	return nil
}

// getDecoder retrieves a decoder for a given struct type from the cache.
// If the decoder is not found in the cache, it builds a new one, caches it, and returns it.
func getDecoder(destType reflect.Type) *internal.Decoder {
	cacheKey := internal.DecoderCacheKey{DestType: destType}
	if decoder, ok := internal.GetDecoder(cacheKey); ok {
		return decoder
	}

	// Slow path: build a new decoder.
	decoder := &internal.Decoder{
		Fields:   make(map[string]*internal.FieldDecoder),
		FieldArr: make([]*internal.FieldDecoder, 0),
	}

	buildDecoderFields(destType, nil, 0, make(map[reflect.Type]struct{}), decoder)

	// Cache the new decoder.
	internal.SetDecoder(cacheKey, decoder)
	return decoder
}

// buildDecoderFields recursively traverses a struct type and populates the decoder with field information.
func buildDecoderFields(
	t reflect.Type,
	indexPrefix []int,
	depth int,
	ancestors map[reflect.Type]struct{},
	decoder *internal.Decoder,
) {
	if _, exists := ancestors[t]; exists {
		return
	}
	ancestors[t] = struct{}{}
	defer delete(ancestors, t)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip unexported fields.
		if isUnexportedField(field) {
			continue
		}

		tag := field.Tag.Get("mconv")
		if tag == "" {
			tag = field.Tag.Get("json")
		}
		if tag == "" {
			tag = field.Tag.Get("yaml")
		}

		if tag == "-" {
			continue
		}

		key := field.Name
		parts := strings.Split(tag, ",")
		if parts[0] == "-" {
			continue
		}
		if len(parts) > 0 && parts[0] != "" {
			key = parts[0]
		}

		if field.Anonymous && parts[0] == "" {
			embeddedType := field.Type
			if embeddedType.Kind() == reflect.Ptr {
				embeddedType = embeddedType.Elem()
			}
			if embeddedType.Kind() == reflect.Struct {
				buildDecoderFields(
					embeddedType,
					append(append([]int(nil), indexPrefix...), i),
					depth+1,
					ancestors,
					decoder,
				)
				continue
			}
		}

		fieldDecoder := &internal.FieldDecoder{
			Field: field,
			Index: append(append([]int(nil), indexPrefix...), i),
			Name:  key,
			Depth: depth,
		}

		existing, exists := decoder.Fields[key]
		switch {
		case !exists:
			decoder.FieldArr = append(decoder.FieldArr, fieldDecoder)
			decoder.Fields[key] = fieldDecoder
		case depth < existing.Depth:
			*existing = *fieldDecoder
		case depth == existing.Depth:
			existing.Ambiguous = true
		}
	}
}

func fieldByIndex(value reflect.Value, indexes []int) reflect.Value {
	current := value
	for _, index := range indexes {
		for current.Kind() == reflect.Ptr {
			if current.IsNil() {
				return reflect.Value{}
			}
			current = current.Elem()
		}
		if current.Kind() != reflect.Struct || index >= current.NumField() {
			return reflect.Value{}
		}
		current = current.Field(index)
	}
	return current
}

func fieldByIndexCopyAlloc(value reflect.Value, indexes []int) (reflect.Value, error) {
	current := value
	for _, index := range indexes {
		for current.Kind() == reflect.Ptr {
			if !current.CanSet() {
				return reflect.Value{}, errors.New("embedded pointer cannot be set")
			}
			cloned := reflect.New(current.Type().Elem())
			if !current.IsNil() {
				cloned.Elem().Set(current.Elem())
			}
			current.Set(cloned)
			current = cloned.Elem()
		}
		if current.Kind() != reflect.Struct || index >= current.NumField() {
			return reflect.Value{}, errors.New("invalid field index")
		}
		current = current.Field(index)
	}
	return current, nil
}

// isUnexportedField checks if a struct field is unexported.
func isUnexportedField(field reflect.StructField) bool {
	return field.PkgPath != ""
}
