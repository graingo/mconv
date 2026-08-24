package basic_test

import (
	"testing"
	"time"

	"github.com/graingo/mconv"
)

type (
	definedBool     bool
	definedInt      int64
	definedUint     uint32
	definedFloat    float32
	definedComplex  complex64
	definedString   string
	definedDuration time.Duration
)

func TestDefinedScalarConversions(t *testing.T) {
	tests := []struct {
		name    string
		convert func() (interface{}, error)
		want    interface{}
	}{
		{
			name: "bool",
			convert: func() (interface{}, error) {
				return mconv.ToBoolE(definedBool(true))
			},
			want: true,
		},
		{
			name: "int",
			convert: func() (interface{}, error) {
				return mconv.ToInt64E(definedInt(42))
			},
			want: int64(42),
		},
		{
			name: "uint",
			convert: func() (interface{}, error) {
				return mconv.ToUint32E(definedUint(42))
			},
			want: uint32(42),
		},
		{
			name: "float",
			convert: func() (interface{}, error) {
				return mconv.ToFloat32E(definedFloat(4.25))
			},
			want: float32(4.25),
		},
		{
			name: "complex",
			convert: func() (interface{}, error) {
				return mconv.ToComplex64E(definedComplex(complex(4, 2)))
			},
			want: complex64(complex(4, 2)),
		},
		{
			name: "string to int",
			convert: func() (interface{}, error) {
				return mconv.ToIntE(definedString("42"))
			},
			want: 42,
		},
		{
			name: "duration",
			convert: func() (interface{}, error) {
				return mconv.ToDurationE(definedDuration(time.Second))
			},
			want: time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.convert()
			if err != nil {
				t.Fatalf("conversion failed: %v", err)
			}
			if got != tt.want {
				t.Fatalf("got %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestPointerScalarConversions(t *testing.T) {
	input := definedInt(42)
	pointer := &input
	pointerToPointer := &pointer

	got, err := mconv.ToInt64E(pointerToPointer)
	if err != nil {
		t.Fatalf("pointer conversion failed: %v", err)
	}
	if got != 42 {
		t.Fatalf("got %d, want 42", got)
	}

	var nilPointer *definedInt
	got, err = mconv.ToInt64E(nilPointer)
	if err != nil {
		t.Fatalf("nil pointer conversion failed: %v", err)
	}
	if got != 0 {
		t.Fatalf("got %d, want zero", got)
	}
}

func TestToStringPreservesStringerSemantics(t *testing.T) {
	duration := time.Second
	if got := mconv.ToString(&duration); got != "1s" {
		t.Fatalf("got %q, want %q", got, "1s")
	}
}

type nilSafeStringer struct{}

func (*nilSafeStringer) String() string {
	panic("String must not be called for a nil pointer")
}

func TestToStringHandlesNilStringerPointer(t *testing.T) {
	var value *nilSafeStringer
	got, err := mconv.ToStringE(value)
	if err != nil {
		t.Fatalf("nil stringer conversion failed: %v", err)
	}
	if got != "" {
		t.Fatalf("got %q, want an empty string", got)
	}
}
