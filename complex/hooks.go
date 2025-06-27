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
