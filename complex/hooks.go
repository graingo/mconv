package complex

import (
	"reflect"
	"time"

	"github.com/graingo/mconv/basic"
)

// stringToTimeHookFunc returns a HookFunc that converts string to time.Time.
func stringToTimeHookFunc() HookFunc {
	return func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		// We only care about converting string to time.Time.
		if from.Kind() != reflect.String || to != reflect.TypeOf(time.Time{}) {
			return data, nil
		}

		s, ok := data.(string)
		if !ok {
			return data, nil
		}

		// Use the existing ToTimeE for conversion.
		return basic.ToTimeE(s)
	}
}

// intToBoolHookFunc returns a HookFunc that converts integer to bool.
func intToBoolHookFunc() HookFunc {
	return func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if to.Kind() != reflect.Bool {
			return data, nil
		}

		// We handle various integer types.
		switch from.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			// For signed integers
			val := reflect.ValueOf(data).Int()
			return val != 0, nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			// For unsigned integers
			val := reflect.ValueOf(data).Uint()
			return val != 0, nil
		}

		return data, nil
	}
}
