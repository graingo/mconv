package mconv_test

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/graingo/mconv"
)

func TestGenericRootAPI(t *testing.T) {
	type User struct {
		ID   int
		Name string
	}

	user, err := mconv.ToE[User](map[string]interface{}{
		"ID":   "7",
		"Name": "maltose",
	})
	if err != nil {
		t.Fatalf("struct conversion failed: %v", err)
	}
	if user != (User{ID: 7, Name: "maltose"}) {
		t.Fatalf("unexpected user: %#v", user)
	}

	if got := mconv.To[int]("42"); got != 42 {
		t.Fatalf("got %d, want 42", got)
	}
	if got := mconv.ToSliceT[string]([2]int{1, 2}); !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Fatalf("unexpected slice: %#v", got)
	}
	if got := mconv.ToMapT[int, int](map[int]string{1: "2"}); !reflect.DeepEqual(got, map[int]int{1: 2}) {
		t.Fatalf("unexpected map: %#v", got)
	}
}

func TestGenericRootAPIAppliesCustomHooks(t *testing.T) {
	type Label string
	hook := func(from, to reflect.Type, data interface{}) (interface{}, error) {
		if from.Kind() == reflect.Int && to == reflect.TypeOf(Label("")) {
			return Label("hooked"), nil
		}
		return data, nil
	}

	got, err := mconv.ToE[Label](1, hook)
	if err != nil {
		t.Fatalf("hook conversion failed: %v", err)
	}
	if got != "hooked" {
		t.Fatalf("got %q, want hooked", got)
	}
}

func TestPublicConversionErrors(t *testing.T) {
	_, err := mconv.ToInt8E(128)
	if !errors.Is(err, mconv.ErrOverflow) {
		t.Fatalf("expected ErrOverflow, got %v", err)
	}

	var conversionErr *mconv.ConversionError
	if !errors.As(err, &conversionErr) {
		t.Fatalf("expected ConversionError, got %T", err)
	}
	if conversionErr.TargetType != "int8" {
		t.Fatalf("unexpected target type: %q", conversionErr.TargetType)
	}
}

func TestConversionErrorIncludesNestedPath(t *testing.T) {
	type User struct{ Age int }
	type Config struct{ Users []User }

	_, err := mconv.ToE[Config](map[string]interface{}{
		"Users": []interface{}{
			map[string]interface{}{"Age": "invalid"},
		},
	})
	if err == nil {
		t.Fatal("invalid nested value should fail")
	}

	var conversionErr *mconv.ConversionError
	if !errors.As(err, &conversionErr) {
		t.Fatalf("expected ConversionError, got %T: %v", err, err)
	}
	if conversionErr.Path != "Users[0].Age" {
		t.Fatalf("got path %q, want %q", conversionErr.Path, "Users[0].Age")
	}
}

func TestGenericMapPreservesComparableKeys(t *testing.T) {
	type Key struct{ ID int }
	source := map[Key]string{{ID: 1}: "2"}

	converted, err := mconv.ToE[map[Key]int](source)
	if err != nil {
		t.Fatalf("map conversion failed: %v", err)
	}
	if !reflect.DeepEqual(converted, map[Key]int{{ID: 1}: 2}) {
		t.Fatalf("unexpected map: %#v", converted)
	}
}

func TestRootFacadeContracts(t *testing.T) {
	if got := mconv.ToString(1); got != "1" {
		t.Fatalf("unexpected string: %q", got)
	}
	if got, err := mconv.ToStringE(1); err != nil || got != "1" {
		t.Fatalf("unexpected string result: %q, %v", got, err)
	}

	if mconv.ToInt("1") != 1 || mconv.ToInt64("1") != 1 || mconv.ToInt32("1") != 1 ||
		mconv.ToInt16("1") != 1 || mconv.ToInt8("1") != 1 {
		t.Fatal("signed integer facade conversion failed")
	}
	if value, err := mconv.ToIntE("1"); err != nil || value != 1 {
		t.Fatal("ToIntE facade conversion failed")
	}
	if value, err := mconv.ToInt64E("1"); err != nil || value != 1 {
		t.Fatal("ToInt64E facade conversion failed")
	}
	if value, err := mconv.ToInt32E("1"); err != nil || value != 1 {
		t.Fatal("ToInt32E facade conversion failed")
	}
	if value, err := mconv.ToInt16E("1"); err != nil || value != 1 {
		t.Fatal("ToInt16E facade conversion failed")
	}
	if value, err := mconv.ToInt8E("1"); err != nil || value != 1 {
		t.Fatal("ToInt8E facade conversion failed")
	}

	if mconv.ToUint("1") != 1 || mconv.ToUint64("1") != 1 || mconv.ToUint32("1") != 1 ||
		mconv.ToUint16("1") != 1 || mconv.ToUint8("1") != 1 {
		t.Fatal("unsigned integer facade conversion failed")
	}
	if value, err := mconv.ToUintE("1"); err != nil || value != 1 {
		t.Fatal("ToUintE facade conversion failed")
	}
	if value, err := mconv.ToUint64E("1"); err != nil || value != 1 {
		t.Fatal("ToUint64E facade conversion failed")
	}
	if value, err := mconv.ToUint32E("1"); err != nil || value != 1 {
		t.Fatal("ToUint32E facade conversion failed")
	}
	if value, err := mconv.ToUint16E("1"); err != nil || value != 1 {
		t.Fatal("ToUint16E facade conversion failed")
	}
	if value, err := mconv.ToUint8E("1"); err != nil || value != 1 {
		t.Fatal("ToUint8E facade conversion failed")
	}

	if mconv.ToFloat64("1.5") != 1.5 || mconv.ToFloat32("1.5") != 1.5 {
		t.Fatal("float facade conversion failed")
	}
	if value, err := mconv.ToFloat64E("1.5"); err != nil || value != 1.5 {
		t.Fatal("ToFloat64E facade conversion failed")
	}
	if value, err := mconv.ToFloat32E("1.5"); err != nil || value != 1.5 {
		t.Fatal("ToFloat32E facade conversion failed")
	}
	if !mconv.ToBool("true") {
		t.Fatal("bool facade conversion failed")
	}
	if value, err := mconv.ToBoolE("true"); err != nil || !value {
		t.Fatal("ToBoolE facade conversion failed")
	}

	if mconv.ToComplex128("1+2i") != complex(1, 2) || mconv.ToComplex64("1+2i") != complex(1, 2) {
		t.Fatal("complex facade conversion failed")
	}
	if value, err := mconv.ToComplex128E("1+2i"); err != nil || value != complex(1, 2) {
		t.Fatal("ToComplex128E facade conversion failed")
	}
	if value, err := mconv.ToComplex64E("1+2i"); err != nil || value != complex(1, 2) {
		t.Fatal("ToComplex64E facade conversion failed")
	}

	if mconv.ToTime("2026-08-24").IsZero() {
		t.Fatal("time facade conversion failed")
	}
	if value, err := mconv.ToTimeE("2026-08-24"); err != nil || value.IsZero() {
		t.Fatal("ToTimeE facade conversion failed")
	}
	if mconv.ToDuration("1s") != time.Second {
		t.Fatal("duration facade conversion failed")
	}
	if value, err := mconv.ToDurationE("1s"); err != nil || value != time.Second {
		t.Fatal("ToDurationE facade conversion failed")
	}

	if len(mconv.ToSlice([2]int{1, 2})) != 2 {
		t.Fatal("slice facade conversion failed")
	}
	if value, err := mconv.ToSliceE([2]int{1, 2}); err != nil || len(value) != 2 {
		t.Fatal("ToSliceE facade conversion failed")
	}
	if !reflect.DeepEqual(mconv.ToStringSlice([]int{1, 2}), []string{"1", "2"}) {
		t.Fatal("string slice facade conversion failed")
	}
	if value, err := mconv.ToStringSliceE([]int{1, 2}); err != nil || !reflect.DeepEqual(value, []string{"1", "2"}) {
		t.Fatal("ToStringSliceE facade conversion failed")
	}
	if !reflect.DeepEqual(mconv.ToIntSlice([]string{"1", "2"}), []int{1, 2}) {
		t.Fatal("int slice facade conversion failed")
	}
	if value, err := mconv.ToIntSliceE([]string{"1", "2"}); err != nil || !reflect.DeepEqual(value, []int{1, 2}) {
		t.Fatal("ToIntSliceE facade conversion failed")
	}
	if !reflect.DeepEqual(mconv.ToFloat64Slice([]string{"1", "2"}), []float64{1, 2}) {
		t.Fatal("float slice facade conversion failed")
	}
	if value, err := mconv.ToFloat64SliceE([]string{"1", "2"}); err != nil || !reflect.DeepEqual(value, []float64{1, 2}) {
		t.Fatal("ToFloat64SliceE facade conversion failed")
	}

	sourceMap := map[string]interface{}{"a": "1"}
	if len(mconv.ToMap(sourceMap)) != 1 {
		t.Fatal("map facade conversion failed")
	}
	if value, err := mconv.ToMapE(sourceMap); err != nil || len(value) != 1 {
		t.Fatal("ToMapE facade conversion failed")
	}
	if !reflect.DeepEqual(mconv.ToStringMap(sourceMap), map[string]string{"a": "1"}) {
		t.Fatal("string map facade conversion failed")
	}
	if value, err := mconv.ToStringMapE(sourceMap); err != nil || !reflect.DeepEqual(value, map[string]string{"a": "1"}) {
		t.Fatal("ToStringMapE facade conversion failed")
	}
	if !reflect.DeepEqual(mconv.ToIntMap(sourceMap), map[string]int{"a": 1}) {
		t.Fatal("int map facade conversion failed")
	}
	if value, err := mconv.ToIntMapE(sourceMap); err != nil || !reflect.DeepEqual(value, map[string]int{"a": 1}) {
		t.Fatal("ToIntMapE facade conversion failed")
	}
	if !reflect.DeepEqual(mconv.ToFloat64Map(sourceMap), map[string]float64{"a": 1}) {
		t.Fatal("float map facade conversion failed")
	}
	if value, err := mconv.ToFloat64MapE(sourceMap); err != nil || !reflect.DeepEqual(value, map[string]float64{"a": 1}) {
		t.Fatal("ToFloat64MapE facade conversion failed")
	}

	jsonText := mconv.ToJSON(sourceMap)
	if value, err := mconv.ToJSONE(sourceMap); err != nil || value != jsonText {
		t.Fatal("JSON facade conversion failed")
	}
	var decoded map[string]interface{}
	mconv.FromJSON(jsonText, &decoded)
	if err := mconv.FromJSONE(jsonText, &decoded); err != nil {
		t.Fatal("FromJSONE facade conversion failed")
	}
	if len(mconv.ToMapFromJSON(jsonText)) != 1 {
		t.Fatal("map JSON facade conversion failed")
	}
	if value, err := mconv.ToMapFromJSONE(jsonText); err != nil || len(value) != 1 {
		t.Fatal("ToMapFromJSONE facade conversion failed")
	}
	jsonSlice := `[1,2]`
	if len(mconv.ToSliceFromJSON(jsonSlice)) != 2 {
		t.Fatal("slice JSON facade conversion failed")
	}
	if value, err := mconv.ToSliceFromJSONE(jsonSlice); err != nil || len(value) != 2 {
		t.Fatal("ToSliceFromJSONE facade conversion failed")
	}

	type target struct{ Value int }
	var destination target
	mconv.ToStruct(map[string]interface{}{"Value": "1"}, &destination)
	if err := mconv.ToStructE(map[string]interface{}{"Value": "2"}, &destination); err != nil || destination.Value != 2 {
		t.Fatal("struct facade conversion failed")
	}

	mconv.SetStringCacheSize(1)
	mconv.SetTimeCacheSize(1)
	mconv.ClearStringCache()
	mconv.ClearTimeCache()
	mconv.ClearAllCaches()
	mconv.SetStringCacheSize(0)
	mconv.SetTimeCacheSize(0)
	mconv.SetTypeInfoCacheSize(1)
	mconv.SetConversionCacheSize(1)
	mconv.ClearTypeInfoCache()
	mconv.ClearConversionCache()
}
