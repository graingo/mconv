package mconv_test

import (
	"testing"
	"time"

	"github.com/graingo/mconv"
)

func TestToTime(t *testing.T) {
	now := time.Now()
	tests := []struct {
		input    interface{}
		expected time.Time
	}{
		{now, now},
		{now.Unix(), time.Unix(now.Unix(), 0)},
		{nil, time.Time{}},
	}

	for _, test := range tests {
		got := mconv.ToTime(test.input)
		if !got.Equal(test.expected) {
			t.Errorf("mconv.ToTime(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}

func TestToTimeE(t *testing.T) {
	now := time.Now()
	rfc3339 := now.Format(time.RFC3339)
	tests := []struct {
		input    interface{}
		expected time.Time
		isErr    bool
	}{
		{now, now, false},
		{now.Unix(), time.Unix(now.Unix(), 0), false},
		{rfc3339, now.Truncate(time.Second), false},
		{"invalid", time.Time{}, true},
		{nil, time.Time{}, false},
	}

	for _, test := range tests {
		got, err := mconv.ToTimeE(test.input)
		if test.isErr && err == nil {
			t.Errorf("mconv.ToTimeE(%v) expected error", test.input)
		}
		if !test.isErr && err != nil {
			t.Errorf("mconv.ToTimeE(%v) unexpected error: %v", test.input, err)
		}
		if !test.isErr && !got.Equal(test.expected) {
			t.Errorf("mconv.ToTimeE(%v) = %v; want %v", test.input, got, test.expected)
		}
	}
}
