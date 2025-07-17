package complex_test

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

type UserWithBoolFloat struct {
	ID       int
	IsActive bool    `mconv:"is_active"`
	Score    float64 `mconv:"score"`
	Balance  float32
}

type TestDuration struct {
	ID       int
	Duration time.Duration `mconv:"duration"`
}

type UserWithAllTags struct {
	ID       int    `mconv:"mconv_id" json:"json_id" yaml:"yaml_id"`
	Name     string `json:"json_name" yaml:"yaml_name"`
	Address  string `yaml:"yaml_address"`
	Untagged string
}

func TestStruct(t *testing.T) {
	t.Run("SimpleConversion", func(t *testing.T) {
		source := map[string]interface{}{"ID": 1, "Name": "Alice"}
		var target SimpleUser
		err := complex.ToStructE(source, &target)
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
		err := complex.ToStructE(source, &target)
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
		err := complex.ToStructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		expected := SimpleUser{ID: 3, Name: "Charlie"}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("BoolAndFloatConversion", func(t *testing.T) {
		source := map[string]interface{}{
			"ID":        8,
			"is_active": true,
			"score":     99.9,
			"Balance":   123.45,
		}
		var target UserWithBoolFloat
		err := complex.ToStructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := UserWithBoolFloat{
			ID:       8,
			IsActive: true,
			Score:    99.9,
			Balance:  123.45,
		}
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
		err := complex.ToStructE(source, &target)
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
		err := complex.ToStructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		parsedTime, _ := time.Parse(time.RFC3339, timeStr)
		expected := UserWithTime{ID: 5, Name: "David", CreatedAt: parsedTime}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("BuiltInIntToBoolHook", func(t *testing.T) {
		source := map[string]interface{}{
			"is_active": 1,
		}
		var target UserWithBoolFloat
		err := complex.ToStructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := UserWithBoolFloat{
			IsActive: true,
		}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}
	})

	t.Run("BuiltInStringToDurationHook", func(t *testing.T) {
		source := map[string]interface{}{
			"duration": "30s",
		}
		var target TestDuration
		err := complex.ToStructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := TestDuration{
			Duration: 30 * time.Second,
		}
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
		err := complex.ToStructE(source, &target, intToStringHook)
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
		err := complex.ToStructE(source, &target)
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

	t.Run("TagPriority", func(t *testing.T) {
		source := map[string]interface{}{
			"mconv_id":     1,
			"json_name":    "From JSON",
			"yaml_address": "From YAML",
			"Untagged":     "No tag",
		}
		var target UserWithAllTags
		err := complex.ToStructE(source, &target)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := UserWithAllTags{
			ID:       1,
			Name:     "From JSON",
			Address:  "From YAML",
			Untagged: "No tag",
		}
		if !reflect.DeepEqual(target, expected) {
			t.Errorf("expected %+v, got %+v", expected, target)
		}

		// Test json overriding yaml
		source = map[string]interface{}{
			"yaml_name": "From YAML",
			"json_name": "From JSON",
		}
		var target2 UserWithAllTags
		err = complex.ToStructE(source, &target2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if target2.Name != "From JSON" {
			t.Errorf("expected Name to be 'From JSON', but got '%s'", target2.Name)
		}

		// Test mconv overriding json and yaml
		source = map[string]interface{}{
			"mconv_id": 99,
			"json_id":  -1,
			"yaml_id":  -2,
		}
		var target3 UserWithAllTags
		err = complex.ToStructE(source, &target3)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if target3.ID != 99 {
			t.Errorf("expected ID to be 99, but got %d", target3.ID)
		}
	})

}
