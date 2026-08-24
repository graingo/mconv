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

func TestTypedContainerConversionsShareScalarSemantics(t *testing.T) {
	type definedInt int

	slice, err := complex.ToStringSliceE([2]definedInt{1, 2})
	if err != nil {
		t.Fatalf("array conversion failed: %v", err)
	}
	if !reflect.DeepEqual(slice, []string{"1", "2"}) {
		t.Fatalf("unexpected string slice: %#v", slice)
	}

	convertedMap, err := complex.ToStringMapE(map[string]definedInt{"answer": 42})
	if err != nil {
		t.Fatalf("map conversion failed: %v", err)
	}
	if !reflect.DeepEqual(convertedMap, map[string]string{"answer": "42"}) {
		t.Fatalf("unexpected string map: %#v", convertedMap)
	}
}

func TestSliceConversionDereferencesPointersAndPreservesNil(t *testing.T) {
	type definedSlice []int

	input := definedSlice{1, 2}
	pointer := &input
	converted, err := complex.ToSliceE(&pointer)
	if err != nil {
		t.Fatalf("pointer conversion failed: %v", err)
	}
	if !reflect.DeepEqual(converted, []interface{}{1, 2}) {
		t.Fatalf("unexpected slice: %#v", converted)
	}

	var nilSlice definedSlice
	converted, err = complex.ToSliceE(nilSlice)
	if err != nil {
		t.Fatalf("nil slice conversion failed: %v", err)
	}
	if converted != nil {
		t.Fatalf("got %#v, want nil", converted)
	}
}

func TestMapConversionPreservesDefinedNilMap(t *testing.T) {
	type definedMap map[string]int
	var input definedMap

	converted, err := complex.ToMapE(input)
	if err != nil {
		t.Fatalf("nil map conversion failed: %v", err)
	}
	if converted != nil {
		t.Fatalf("got %#v, want nil", converted)
	}
}

func TestTypedConvertersPreserveNilCollections(t *testing.T) {
	var (
		slice []interface{}
		m     map[string]interface{}
	)

	stringSlice, err := complex.ToStringSliceE(slice)
	if err != nil || stringSlice != nil {
		t.Fatalf("nil slice conversion returned %#v, %v", stringSlice, err)
	}
	intSlice, err := complex.ToIntSliceE(slice)
	if err != nil || intSlice != nil {
		t.Fatalf("nil slice conversion returned %#v, %v", intSlice, err)
	}
	floatSlice, err := complex.ToFloat64SliceE(slice)
	if err != nil || floatSlice != nil {
		t.Fatalf("nil slice conversion returned %#v, %v", floatSlice, err)
	}

	stringMap, err := complex.ToStringMapE(m)
	if err != nil || stringMap != nil {
		t.Fatalf("nil map conversion returned %#v, %v", stringMap, err)
	}
	intMap, err := complex.ToIntMapE(m)
	if err != nil || intMap != nil {
		t.Fatalf("nil map conversion returned %#v, %v", intMap, err)
	}
	floatMap, err := complex.ToFloat64MapE(m)
	if err != nil || floatMap != nil {
		t.Fatalf("nil map conversion returned %#v, %v", floatMap, err)
	}
}

func TestMapConversionRejectsStringifiedKeyCollisions(t *testing.T) {
	_, err := complex.ToMapE(map[interface{}]interface{}{1: "number", "1": "string"})
	if err == nil {
		t.Fatal("colliding map keys should not silently overwrite data")
	}
}

func TestMapConversionUsesShallowTaggedField(t *testing.T) {
	type Embedded struct {
		Name string `json:"name"`
	}
	type source struct {
		Embedded
		DisplayName string `json:"name"`
	}

	converted, err := complex.ToMapE(source{Embedded: Embedded{Name: "first"}, DisplayName: "second"})
	if err != nil {
		t.Fatalf("conversion failed: %v", err)
	}
	if !reflect.DeepEqual(converted, map[string]interface{}{"name": "second"}) {
		t.Fatalf("unexpected map: %#v", converted)
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

func TestToStructECommitsAtomically(t *testing.T) {
	type target struct {
		Name string
		Age  int
	}

	result := target{Name: "before", Age: 7}
	err := complex.ToStructE(map[string]interface{}{
		"Name": "after",
		"Age":  "invalid",
	}, &result)
	if err == nil {
		t.Fatal("invalid field should fail")
	}
	if result != (target{Name: "before", Age: 7}) {
		t.Fatalf("failed conversion mutated destination: %#v", result)
	}
}

func TestToStructECommitsEmbeddedPointersAtomically(t *testing.T) {
	type Embedded struct {
		Name string
		Age  int
	}
	type target struct{ *Embedded }

	original := &Embedded{Name: "before", Age: 7}
	result := target{Embedded: original}
	err := complex.ToStructE(map[string]interface{}{
		"Name": "after",
		"Age":  "invalid",
	}, &result)
	if err == nil {
		t.Fatal("invalid field should fail")
	}
	if result.Embedded != original || *original != (Embedded{Name: "before", Age: 7}) {
		t.Fatalf("failed conversion mutated embedded destination: %#v", result)
	}
}

func TestToStructEShallowerFieldsTakePriority(t *testing.T) {
	type Embedded struct{ Name string }
	type target struct {
		Embedded
		Name string
	}

	result := target{Embedded: Embedded{Name: "embedded"}}
	if err := complex.ToStructE(map[string]interface{}{"Name": "outer"}, &result); err != nil {
		t.Fatalf("conversion failed: %v", err)
	}
	if result.Name != "outer" || result.Embedded.Name != "embedded" {
		t.Fatalf("unexpected field selection: %#v", result)
	}
}

func TestToStructERejectsSameDepthDestinationAmbiguity(t *testing.T) {
	type First struct{ Name string }
	type Second struct{ Name string }
	type target struct {
		First
		Second
	}

	var result target
	if err := complex.ToStructE(map[string]interface{}{"Name": "value"}, &result); err == nil {
		t.Fatal("same-depth destination fields should be ambiguous")
	}
}

func TestToStructETaggedAnonymousFieldRemainsNested(t *testing.T) {
	type Embedded struct{ Name string }
	type target struct {
		Embedded `json:"profile"`
	}

	var result target
	if err := complex.ToStructE(
		map[string]interface{}{"profile": map[string]interface{}{"Name": "maltose"}},
		&result,
	); err != nil {
		t.Fatalf("conversion failed: %v", err)
	}
	if result.Name != "maltose" {
		t.Fatalf("unexpected nested field: %#v", result)
	}
}

func TestToStructERejectsCaseFoldedDestinationAmbiguity(t *testing.T) {
	type target struct {
		Upper string `mconv:"NAME"`
		Title string `mconv:"Name"`
	}

	var result target
	if err := complex.ToStructE(map[string]interface{}{"name": "value"}, &result); err == nil {
		t.Fatal("case-folded destination fields should be ambiguous")
	}

	if err := complex.ToStructE(map[string]interface{}{"NAME": "exact"}, &result); err != nil {
		t.Fatalf("exact destination field should win: %v", err)
	}
	if result.Upper != "exact" || result.Title != "" {
		t.Fatalf("unexpected exact-match result: %#v", result)
	}
}

func TestStructFieldPlanHandlesRecursiveEmbedding(t *testing.T) {
	type Recursive struct {
		*Recursive
		Value string
	}

	var result Recursive
	if err := complex.ToStructE(map[string]interface{}{"Value": "decoded"}, &result); err != nil {
		t.Fatalf("recursive type conversion failed: %v", err)
	}
	if result.Value != "decoded" {
		t.Fatalf("unexpected decoded value: %#v", result)
	}

	result.Recursive = &result
	converted, err := complex.ToMapE(result)
	if err != nil {
		t.Fatalf("recursive value conversion failed: %v", err)
	}
	if !reflect.DeepEqual(converted, map[string]interface{}{"Value": "decoded"}) {
		t.Fatalf("unexpected recursive map: %#v", converted)
	}
}

func TestStructToMapUsesShallowFieldPriority(t *testing.T) {
	type Embedded struct{ Name string }
	type Source struct {
		Embedded
		Name string
	}

	converted, err := complex.ToMapE(Source{
		Embedded: Embedded{Name: "embedded"},
		Name:     "outer",
	})
	if err != nil {
		t.Fatalf("conversion failed: %v", err)
	}
	if !reflect.DeepEqual(converted, map[string]interface{}{"Name": "outer"}) {
		t.Fatalf("unexpected map: %#v", converted)
	}
}

func TestStructFieldPlanHonorsIgnoredTagOptions(t *testing.T) {
	type target struct {
		Visible string
		Hidden  string `json:"-,omitempty"`
	}

	result := target{Hidden: "preserved"}
	if err := complex.ToStructE(map[string]interface{}{
		"Visible": "decoded",
		"-":       "hidden",
	}, &result); err != nil {
		t.Fatalf("conversion failed: %v", err)
	}
	if result.Visible != "decoded" || result.Hidden != "preserved" {
		t.Fatalf("unexpected result: %#v", result)
	}
}
