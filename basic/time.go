package basic

import (
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
