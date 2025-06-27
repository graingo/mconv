package basic_test

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

func TestToDuration(t *testing.T) {
	tests := []struct {
		input interface{}
		want  time.Duration
	}{
		// Standard time.ParseDuration compatible strings
		{
			input: "1h",
			want:  time.Hour,
		},
		{
			input: "1h30m",
			want:  time.Hour + 30*time.Minute,
		},
		{
			input: "6s",
			want:  6 * time.Second,
		},
		{
			input: "8ms",
			want:  8 * time.Millisecond,
		},
		{
			input: "9m",
			want:  9 * time.Minute,
		},
		{
			input: "2.5s",
			want:  2*time.Second + 500*time.Millisecond,
		},

		// Numeric values (interpreted as nanoseconds)
		{
			input: int64(3600000000000), // 3.6 billion nanoseconds = 1 hour
			want:  time.Hour,
		},
		{
			input: 3600, // 3600 nanoseconds
			want:  3600 * time.Nanosecond,
		},
		{
			input: float64(1e9), // 1 billion nanoseconds = 1 second
			want:  time.Second,
		},

		// String numbers (interpreted as nanoseconds)
		{
			input: "3600", // 3600 nanoseconds as string
			want:  3600 * time.Nanosecond,
		},
		{
			input: "1000000000", // 1 billion nanoseconds = 1 second
			want:  time.Second,
		},

		// Edge cases
		{
			input: nil,
			want:  0,
		},
	}

	for _, tc := range tests {
		got := mconv.ToDuration(tc.input)
		if got != tc.want {
			t.Errorf("ToDuration(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}

func TestToDurationE(t *testing.T) {
	tests := []struct {
		input   interface{}
		want    time.Duration
		wantErr bool
	}{
		// Standard time.ParseDuration compatible strings
		{
			input:   "1h",
			want:    time.Hour,
			wantErr: false,
		},
		{
			input:   "1h30m",
			want:    time.Hour + 30*time.Minute,
			wantErr: false,
		},
		{
			input:   "6s",
			want:    6 * time.Second,
			wantErr: false,
		},
		{
			input:   "8ms",
			want:    8 * time.Millisecond,
			wantErr: false,
		},
		{
			input:   "9m",
			want:    9 * time.Minute,
			wantErr: false,
		},
		{
			input:   "2.5s",
			want:    2*time.Second + 500*time.Millisecond,
			wantErr: false,
		},

		// Numeric values (interpreted as nanoseconds)
		{
			input:   int64(3600000000000), // 3.6 billion nanoseconds = 1 hour
			want:    time.Hour,
			wantErr: false,
		},
		{
			input:   3600, // 3600 nanoseconds
			want:    3600 * time.Nanosecond,
			wantErr: false,
		},
		{
			input:   float64(1e9), // 1 billion nanoseconds = 1 second
			want:    time.Second,
			wantErr: false,
		},

		// String numbers (interpreted as nanoseconds)
		{
			input:   "3600", // 3600 nanoseconds as string
			want:    3600 * time.Nanosecond,
			wantErr: false,
		},
		{
			input:   "1000000000", // 1 billion nanoseconds = 1 second
			want:    time.Second,
			wantErr: false,
		},

		// Error cases
		{
			input:   "invalid",
			want:    0,
			wantErr: true,
		},

		// Edge cases
		{
			input:   nil,
			want:    0,
			wantErr: false,
		},
	}

	for _, tc := range tests {
		got, err := mconv.ToDurationE(tc.input)
		if (err != nil) != tc.wantErr {
			t.Errorf("ToDurationE(%v) error = %v, wantErr %v", tc.input, err, tc.wantErr)
			continue
		}

		if !tc.wantErr && got != tc.want {
			t.Errorf("ToDurationE(%v) = %v, want %v", tc.input, got, tc.want)
		}
	}
}
