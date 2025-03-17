package internal

import (
	"errors"
	"fmt"
)

// define errors
var (
	ErrNilValue          = errors.New("input value is nil")
	ErrUnsupportedType   = errors.New("unsupported type")
	ErrConversionFailed  = errors.New("conversion failed")
	ErrOverflow          = errors.New("value overflow")
	ErrInvalidFormat     = errors.New("invalid format")
	ErrInvalidTimeFormat = errors.New("invalid time format")
	ErrInvalidJSONFormat = errors.New("invalid JSON format")
)

// ConversionError represents a conversion error.
type ConversionError struct {
	Value      interface{}
	Type       string
	TargetType string
	Err        error
}

// Error implements the error interface.
func (e *ConversionError) Error() string {
	return fmt.Sprintf("unable to convert %#v of type %s to %s: %v", e.Value, e.Type, e.TargetType, e.Err)
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
