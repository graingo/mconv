package basic_test

import (
	"testing"

	"encoding/json"
	"html/template"
	"time"

	"github.com/graingo/mconv"
)

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{123, "123"},
		{int64(123), "123"},
		{123.45, "123.45"},
		{float32(123.45), "123.45"},
		{true, "true"},
		{[]byte("test"), "test"},
		{nil, ""},
	}

	for _, test := range tests {
		if got := mconv.ToString(test.input); got != test.expected {
			t.Errorf("mconv.ToString(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToStringE(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		isErr    bool
	}{
		{123, "123", false},
		{int8(123), "123", false},
		{int16(123), "123", false},
		{int32(123), "123", false},
		{int64(123), "123", false},
		{uint(123), "123", false},
		{uint8(123), "123", false},
		{uint16(123), "123", false},
		{uint32(123), "123", false},
		{uint64(123), "123", false},
		{123.45, "123.45", false},
		{true, "true", false},
		{[]byte("test"), "test", false},
		{nil, "", false},
		{complex(1, 2), "(1+2i)", false},
		{complex64(complex(1, 2)), "(1+2i)", false},
		{template.HTML("<b>Hello</b>"), "<b>Hello</b>", false},
		{template.URL("http://a.b"), "http://a.b", false},
		{template.JS("alert(1)"), "alert(1)", false},
		{template.CSS("a{}"), "a{}", false},
		{template.HTMLAttr("a=b"), "a=b", false},
		{json.Number("123"), "123", false},
		{time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), "2020-01-01T00:00:00Z", false},
	}

	for _, test := range tests {
		got, err := mconv.ToStringE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToStringE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToStringE(%v) unexpected error: %v", test.input, err)
		}
		if got != test.expected {
			t.Errorf("mconv.ToStringE(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
