package mconv_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/graingo/mconv/complex"
)

// --- Test Structs ---

type SimpleUser struct {
	ID   int
	Name string
}

type UserWithTags struct {
	UserID   int    `mconv:"user_id"`
	UserName string `mconv:"user_name"`
	Email    string `mconv:"email,omitempty"`
}

type NestedUser struct {
	ID      int
	Profile struct {
		Age  int
		City string
	}
}

type UserWithTime struct {
	ID        int
	Name      string
	CreatedAt time.Time `mconv:"created_at"`
}

type UserWithEmbedded struct {
	ID int
	SimpleUser
	Email string
}

func TestStruct(t *testing.T) {
	t.Run("SimpleConversion", func(t *testing.T) {
		source := map[string]interface{}{"ID": 1, "Name": "Alice"}
		var target SimpleUser
		err := complex.StructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := SimpleUser{ID: 1, Name: "Alice"}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("WithTags", func(t *testing.T) {
		source := map[string]interface{}{"user_id": 2, "user_name": "Bob"}
		var target UserWithTags
		err := complex.StructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := UserWithTags{UserID: 2, UserName: "Bob"}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("CaseInsensitive", func(t *testing.T) {
		source := map[string]interface{}{"id": 3, "nAmE": "Charlie"}
		var target SimpleUser
		err := complex.StructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := SimpleUser{ID: 3, Name: "Charlie"}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("NestedStruct", func(t *testing.T) {
		source := map[string]interface{}{
			"ID": 4,
			"Profile": map[string]interface{}{
				"Age":  30,
				"City": "New York",
			},
		}
		var target NestedUser
		err := complex.StructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := NestedUser{ID: 4, Profile: struct {
			Age  int
			City string
		}{Age: 30, City: "New York"}}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("BuiltInTimeHook", func(t *testing.T) {
		timeStr := "2024-01-01T15:04:05Z"
		source := map[string]interface{}{
			"ID":         5,
			"Name":       "David",
			"created_at": timeStr,
		}
		var target UserWithTime
		err := complex.StructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		parsedTime, _ := time.Parse(time.RFC3339, timeStr)
		expected := UserWithTime{ID: 5, Name: "David", CreatedAt: parsedTime}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("CustomHook", func(t *testing.T) {
		type HookUser struct {
			ID     int
			IsCool string `mconv:"is_cool"`
		}

		intToStringHook := func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
			if from.Kind() == reflect.Int && to.Kind() == reflect.String {
				i, _ := data.(int)
				if i == 1 {
					return "Yes", nil
				}
				return "No", nil
			}
			return data, nil
		}

		source := map[string]interface{}{"ID": 6, "is_cool": 1}
		var target HookUser
		err := complex.StructE(source, &target, intToStringHook)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := HookUser{ID: 6, IsCool: "Yes"}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("EmbeddedStruct", func(t *testing.T) {
		source := map[string]interface{}{
			"ID":    7,
			"Name":  "Embed",
			"Email": "embed@example.com",
		}
		var target UserWithEmbedded
		err := complex.StructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := UserWithEmbedded{
			ID:         7,
			SimpleUser: SimpleUser{ID: 0, Name: "Embed"},
			Email:      "embed@example.com",
		}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})
}
