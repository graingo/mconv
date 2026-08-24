package complex

import "reflect"

func indirectValue(value interface{}) (reflect.Value, bool) {
	if value == nil {
		return reflect.Value{}, false
	}

	rv := reflect.ValueOf(value)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		if rv.IsNil() {
			return reflect.Value{}, false
		}
		rv = rv.Elem()
	}
	return rv, rv.IsValid()
}

// isNilCollectionInput reports whether a collection input resolves to nil.
func isNilCollectionInput(value interface{}) bool {
	if value == nil {
		return true
	}

	rv, valid := indirectValue(value)
	if !valid {
		return true
	}

	return (rv.Kind() == reflect.Map || rv.Kind() == reflect.Slice) && rv.IsNil()
}
