package internal

import (
	"errors"
	"fmt"
	"strings"
)

// define errors
var (
	ErrUnsupportedType   = errors.New("unsupported type")
	ErrConversionFailed  = errors.New("conversion failed")
	ErrOverflow          = errors.New("value overflow")
	ErrInvalidTimeFormat = errors.New("invalid time format")
	ErrInvalidJSONFormat = errors.New("invalid JSON format")
)

// ConversionError represents a conversion error.
type ConversionError struct {
	Value      interface{}
	Type       string
	TargetType string
	Path       string
	Err        error
}

// Error implements the error interface.
func (e *ConversionError) Error() string {
	location := ""
	if e.Path != "" {
		location = " at " + e.Path
	}
	return fmt.Sprintf("unable to convert %#v of type %s%s to %s: %v", e.Value, e.Type, location, e.TargetType, e.Err)
}

// PrependConversionPath adds a field or collection segment to an error path.
func PrependConversionPath(err error, segment string) error {
	if err == nil || segment == "" {
		return err
	}

	var conversionErr *ConversionError
	if errors.As(err, &conversionErr) {
		cloned := *conversionErr
		cloned.Path = joinConversionPath(segment, cloned.Path)
		return &cloned
	}
	return fmt.Errorf("conversion failed at %s: %w", segment, err)
}

func joinConversionPath(prefix, path string) string {
	if path == "" {
		return prefix
	}
	if strings.HasPrefix(path, "[") {
		return prefix + path
	}
	return prefix + "." + path
}

// Unwrap returns the underlying error.
func (e *ConversionError) Unwrap() error {
	return e.Err
}

// NewConversionError creates a new conversion error.
func NewConversionError(value interface{}, targetType string, err error) *ConversionError {
	var valueType string
	if value != nil {
		valueType = fmt.Sprintf("%T", value)
	} else {
		valueType = "nil"
	}
	return &ConversionError{
		Value:      value,
		Type:       valueType,
		TargetType: targetType,
		Err:        err,
	}
}
