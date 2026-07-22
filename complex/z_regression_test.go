package complex_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/graingo/mconv/complex"
)

func TestStructSourceAndEmbeddedPointerDestination(t *testing.T) {
	type source struct {
		Name string `json:"name"`
	}
	type Embedded struct {
		Name string `mconv:"name"`
	}
	type destination struct{ *Embedded }

	var result destination
	if err := complex.ToStructE(source{Name: "maltose"}, &result); err != nil {
		t.Fatalf("struct conversion failed: %v", err)
	}
	if result.Embedded == nil || result.Name != "maltose" {
		t.Fatalf("embedded pointer was not populated: %#v", result)
	}
}

func TestMapAndSliceSupportStructTagsAndArrays(t *testing.T) {
	type item struct {
		ID     int    `json:"id"`
		Secret string `json:"-"`
	}
	converted, err := complex.ToMapE(&item{ID: 7, Secret: "hidden"})
	if err != nil {
		t.Fatal(err)
	}
	if converted["id"] != 7 {
		t.Fatalf("unexpected map: %#v", converted)
	}
	if _, exists := converted["Secret"]; exists {
		t.Fatalf("ignored field leaked into map: %#v", converted)
	}

	slice, err := complex.ToSliceE([2]int{1, 2})
	if err != nil || len(slice) != 2 || slice[1] != 2 {
		t.Fatalf("array conversion failed: %#v, %v", slice, err)
	}
}

func TestMapConversionRejectsStringifiedKeyCollisions(t *testing.T) {
	_, err := complex.ToMapE(map[interface{}]interface{}{1: "number", "1": "string"})
	if err == nil {
		t.Fatal("colliding map keys should not silently overwrite data")
	}
}

func TestMapConversionRejectsEmbeddedFieldCollisions(t *testing.T) {
	type Embedded struct {
		Name string `json:"name"`
	}
	type source struct {
		Embedded
		DisplayName string `json:"name"`
	}

	if _, err := complex.ToMapE(source{Embedded: Embedded{Name: "first"}, DisplayName: "second"}); err == nil {
		t.Fatal("embedded field collisions should not silently overwrite data")
	}
}

func TestGenericMapUsesSemanticStringConversion(t *testing.T) {
	converted := complex.ToMapT[string, string](map[string]int{"a": 1})
	if converted["a"] != "1" {
		t.Fatalf("expected decimal string, got %q", converted["a"])
	}
	type namedString string
	named := complex.ToSliceT[namedString]([]int{65})
	if len(named) != 1 || named[0] != "65" {
		t.Fatalf("expected semantic conversion for named string, got %#v", named)
	}
}

func TestNilReturningHookDoesNotPanicLaterHooks(t *testing.T) {
	type target struct{ Value string }
	result := target{}
	err := complex.ToStructE(
		map[string]interface{}{"Value": "input"},
		&result,
		func(_, _ reflect.Type, data interface{}) (interface{}, error) { return nil, nil },
		func(_, _ reflect.Type, data interface{}) (interface{}, error) { return data, nil },
	)
	if err != nil {
		t.Fatalf("nil-returning hook should stop later conversions without panicking: %v", err)
	}
}

func TestPointerFieldIsUnchangedWhenConversionFails(t *testing.T) {
	type target struct{ Value *int }
	original := 7
	result := target{Value: &original}
	err := complex.ToStructE(map[string]interface{}{"Value": "invalid"}, &result)
	if err == nil {
		t.Fatal("invalid integer should fail")
	}
	if result.Value == nil || *result.Value != 7 {
		t.Fatalf("failed conversion mutated destination: %#v", result)
	}
}

func TestToStructEAmbiguousCaseInsensitiveKeys(t *testing.T) {
	type target struct {
		Name string `mconv:"display_name"`
	}

	var result target
	err := complex.ToStructE(map[string]interface{}{
		"DISPLAY_NAME": "first",
		"Display_Name": "second",
	}, &result)
	if err == nil || !strings.Contains(err.Error(), "ambiguous keys") {
		t.Fatalf("expected an ambiguous key error, got %v", err)
	}

	// An exact key remains deterministic even if another key differs only by case.
	err = complex.ToStructE(map[string]interface{}{
		"display_name": "exact",
		"DISPLAY_NAME": "fallback",
	}, &result)
	if err != nil {
		t.Fatalf("exact key should win: %v", err)
	}
	if result.Name != "exact" {
		t.Fatalf("unexpected exact-match result: %q", result.Name)
	}
}
