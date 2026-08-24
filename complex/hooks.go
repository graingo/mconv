package complex

import (
	"reflect"
	"time"

	"github.com/graingo/mconv/basic"
)

var defaultHooks = []HookFunc{
	stringToTimeHook,
	stringToDurationHook,
}

// stringToTimeHook converts string to time.Time.
func stringToTimeHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() != reflect.String || to != reflect.TypeOf(time.Time{}) {
		return data, nil
	}

	s, ok := data.(string)
	if !ok {
		return data, nil
	}

	return basic.ToTimeE(s)
}

// stringToDurationHook converts string to time.Duration.
func stringToDurationHook(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
	if from.Kind() != reflect.String || to != reflect.TypeOf(time.Duration(0)) {
		return data, nil
	}

	s, ok := data.(string)
	if !ok {
		return data, nil
	}

	return basic.ToDurationE(s)
}
