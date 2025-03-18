package basic

import (
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/graingo/mconv/internal"
)

// ToTimeE converts any type to time.Time with error.
// When value is a string, it will be parsed using the formats.
func ToTimeE(value interface{}, formats ...string) (time.Time, error) {
	if value == nil {
		return time.Time{}, nil
	}

	if cachedValue, ok := internal.GetTimeFromCache(value); ok {
		return cachedValue, nil
	}

	var result time.Time
	var err error

	switch v := value.(type) {
	case time.Time:
		return v, nil
	case string:
		result, err = parseTimeString(v, formats...)
	case int:
		result = time.Unix(int64(v), 0)
	case int64:
		result = time.Unix(v, 0)
	case int32:
		result = time.Unix(int64(v), 0)
	case uint:
		result = time.Unix(int64(v), 0)
	case uint64:
		result = time.Unix(int64(v), 0)
	case uint32:
		result = time.Unix(int64(v), 0)
	default:
		return time.Time{}, internal.NewConversionError(value, "time.Time", internal.ErrUnsupportedType)
	}

	if err == nil {
		internal.AddTimeToCache(value, result)
	}

	return result, err
}

// ToTime converts any type to time.Time.
// When value is a string, it will be parsed using the formats.
func ToTime(value interface{}, formats ...string) time.Time {
	result, _ := ToTimeE(value, formats...)
	return result
}

// parseTimeString parses a string to time.Time
func parseTimeString(s string, formats ...string) (time.Time, error) {
	if len(formats) != 0 {
		if t, err := time.Parse(formats[0], s); err != nil {
			return time.Time{}, internal.NewConversionError(s, "time.Time", err)
		} else {
			return t, nil
		}
	}

	// Try to parse as Unix timestamp
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return time.Unix(i, 0), nil
	}

	// Try common formats
	for _, layout := range []string{
		time.RFC3339,
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04",
		"2006-01-02",
		"01/02/2006",
		"01/02/2006 15:04:05",
		"01/02/2006 15:04",
		"2006/01/02",
		"2006/01/02 15:04:05",
		"2006/01/02 15:04",
	} {
		if t, err := time.Parse(layout, s); err == nil {
			return t, nil
		}
	}

	return time.Time{}, internal.NewConversionError(s, "time.Time", internal.ErrInvalidTimeFormat)
}

// ToDuration converts any type to time.Duration.
func ToDuration(value interface{}) time.Duration {
	result, _ := ToDurationE(value)
	return result
}

// ToDurationE converts any type to time.Duration with error.
func ToDurationE(value interface{}) (time.Duration, error) {
	if value == nil {
		return 0, nil
	}

	switch v := value.(type) {
	case time.Duration:
		return v, nil
	case int:
		return time.Duration(v), nil // Integers are treated as nanoseconds
	case int64:
		return time.Duration(v), nil // Integers are treated as nanoseconds
	case int32:
		return time.Duration(v), nil
	case int16:
		return time.Duration(v), nil
	case int8:
		return time.Duration(v), nil
	case uint:
		return time.Duration(v), nil
	case uint64:
		return time.Duration(v), nil
	case uint32:
		return time.Duration(v), nil
	case uint16:
		return time.Duration(v), nil
	case uint8:
		return time.Duration(v), nil
	case float64:
		return time.Duration(v), nil // Float values are treated as nanoseconds
	case float32:
		return time.Duration(v), nil
	case string:
		// Try to parse using time.ParseDuration first (e.g., "1h", "10m", "30s")
		d, err := time.ParseDuration(v)
		if err == nil {
			return d, nil
		}

		// Try to parse as integer (nanoseconds)
		if i, err := strconv.ParseInt(v, 10, 64); err == nil {
			return time.Duration(i), nil
		}

		// Try to parse as float (nanoseconds)
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return time.Duration(f), nil
		}

		return 0, internal.NewConversionError(value, "time.Duration", fmt.Errorf("cannot parse %q as duration", v))
	default:
		// Handle reflection cases
		rv := reflect.ValueOf(value)
		switch rv.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return time.Duration(rv.Int()), nil
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return time.Duration(rv.Uint()), nil
		case reflect.Float32, reflect.Float64:
			return time.Duration(rv.Float()), nil
		}
		return 0, internal.NewConversionError(value, "time.Duration", internal.ErrUnsupportedType)
	}
}
