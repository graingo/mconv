package mconv_test

import (
	"reflect"
	"testing"

	"github.com/graingo/mconv"
)

type StructBasic struct {
	Name   string
	Age    int
	Score  float64
	Active bool
}

type StructWithTags struct {
	FirstName string `mconv:"name"`
	UserAge   int    `mconv:"age"`
	Omit      string `mconv:"-"`
	NoTag     string
}

type StructNested struct {
	ID    int
	User  StructBasic
	Extra map[string]string
}

type StructWithPointers struct {
	Name  *string
	Age   *int
	Child *StructBasic
}

func TestStruct(t *testing.T) {
	t.Run("BasicMapping", func(t *testing.T) {
		source := map[string]interface{}{
			"Name":   "John",
			"Age":    30,
			"Score":  99.5,
			"Active": true,
		}
		var dest StructBasic
		err := mconv.StructE(source, &dest)
		if err != nil {
			t.Fatalf("StructE() returned an unexpected error: %v", err)
		}
		if dest.Name != "John" {
			t.Errorf("got Name %v, want %v", dest.Name, "John")
		}
		if dest.Age != 30 {
			t.Errorf("got Age %v, want %v", dest.Age, 30)
		}
		if dest.Score != 99.5 {
			t.Errorf("got Score %v, want %v", dest.Score, 99.5)
		}
		if !dest.Active {
			t.Errorf("got Active %v, want %v", dest.Active, true)
		}
	})

	t.Run("CaseInsensitiveMapping", func(t *testing.T) {
		source := map[string]interface{}{
			"name": "Jane",
			"age":  25,
		}
		var dest StructBasic
		err := mconv.StructE(source, &dest)
		if err != nil {
			t.Fatalf("StructE() returned an unexpected error: %v", err)
		}
		if dest.Name != "Jane" {
			t.Errorf("got Name %v, want %v", dest.Name, "Jane")
		}
		if dest.Age != 25 {
			t.Errorf("got Age %v, want %v", dest.Age, 25)
		}
	})

	t.Run("WithTags", func(t *testing.T) {
		source := map[string]interface{}{
			"name":  "Mike",
			"age":   40,
			"Omit":  "should-be-omitted",
			"NoTag": "has-no-tag",
		}
		var dest StructWithTags
		err := mconv.StructE(source, &dest)
		if err != nil {
			t.Fatalf("StructE() returned an unexpected error: %v", err)
		}
		if dest.FirstName != "Mike" {
			t.Errorf("got FirstName %v, want %v", dest.FirstName, "Mike")
		}
		if dest.UserAge != 40 {
			t.Errorf("got UserAge %v, want %v", dest.UserAge, 40)
		}
		if dest.Omit != "" {
			t.Errorf("got Omit %q, want %q", dest.Omit, "")
		}
		if dest.NoTag != "has-no-tag" {
			t.Errorf("got NoTag %q, want %q", dest.NoTag, "has-no-tag")
		}
	})

	t.Run("NestedStruct", func(t *testing.T) {
		source := map[string]interface{}{
			"ID": 1,
			"User": map[string]interface{}{
				"Name": "SubUser",
				"Age":  10,
			},
			"Extra": map[string]interface{}{
				"key1": "val1",
			},
		}
		var dest StructNested
		err := mconv.StructE(source, &dest)
		if err != nil {
			t.Fatalf("StructE() returned an unexpected error: %v", err)
		}
		if dest.ID != 1 {
			t.Errorf("got ID %v, want %v", dest.ID, 1)
		}
		if dest.User.Name != "SubUser" {
			t.Errorf("got User.Name %v, want %v", dest.User.Name, "SubUser")
		}
		if dest.User.Age != 10 {
			t.Errorf("got User.Age %v, want %v", dest.User.Age, 10)
		}
		expectedExtra := map[string]string{"key1": "val1"}
		if !reflect.DeepEqual(dest.Extra, expectedExtra) {
			t.Errorf("got Extra map %v, want %v", dest.Extra, expectedExtra)
		}
	})

	t.Run("PointerFields", func(t *testing.T) {
		name := "PointerMan"
		age := 55
		source := map[string]interface{}{
			"Name": name,
			"Age":  age,
			"Child": map[string]interface{}{
				"Name": "ChildName",
				"Age":  15,
			},
		}

		var dest StructWithPointers
		err := mconv.StructE(source, &dest)
		if err != nil {
			t.Fatalf("StructE() returned an unexpected error: %v", err)
		}

		if dest.Name == nil {
			t.Fatal("expected Name to be non-nil")
		}
		if *dest.Name != name {
			t.Errorf("got Name %v, want %v", *dest.Name, name)
		}

		if dest.Age == nil {
			t.Fatal("expected Age to be non-nil")
		}
		if *dest.Age != age {
			t.Errorf("got Age %v, want %v", *dest.Age, age)
		}

		if dest.Child == nil {
			t.Fatal("expected Child to be non-nil")
		}
		if dest.Child.Name != "ChildName" {
			t.Errorf("got Child.Name %v, want %v", dest.Child.Name, "ChildName")
		}
		if dest.Child.Age != 15 {
			t.Errorf("got Child.Age %v, want %v", dest.Child.Age, 15)
		}
	})

	t.Run("NilSourceValueForPointer", func(t *testing.T) {
		source := map[string]interface{}{
			"Name":  "NotNil",
			"Age":   nil,
			"Child": nil,
		}
		var dest StructWithPointers
		err := mconv.StructE(source, &dest)
		if err != nil {
			t.Fatalf("StructE() returned an unexpected error: %v", err)
		}
		if dest.Name == nil {
			t.Fatal("expected Name to be non-nil")
		}
		if *dest.Name != "NotNil" {
			t.Errorf("got *dest.Name %q, want %q", *dest.Name, "NotNil")
		}
		if dest.Age != nil {
			t.Errorf("expected Age to be nil, but it was %d", *dest.Age)
		}
		if dest.Child != nil {
			t.Errorf("expected Child to be nil, but it was %v", *dest.Child)
		}
	})

	t.Run("TargetNotPointer", func(t *testing.T) {
		source := map[string]interface{}{}
		var dest StructBasic
		err := mconv.StructE(source, dest)
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
	})

	t.Run("TargetNotStructPointer", func(t *testing.T) {
		source := map[string]interface{}{}
		var i int
		err := mconv.StructE(source, &i)
		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
	})
}

func TestScan(t *testing.T) {
	source := map[string]interface{}{
		"Name": "ScanUser",
		"Age":  99,
	}
	var dest StructBasic
	err := mconv.Scan(source, &dest)
	if err != nil {
		t.Fatalf("Scan() returned an unexpected error: %v", err)
	}
	if dest.Name != "ScanUser" {
		t.Errorf("got Name %v, want %v", dest.Name, "ScanUser")
	}
	if dest.Age != 99 {
		t.Errorf("got Age %v, want %v", dest.Age, 99)
	}
}
