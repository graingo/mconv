package mconv_test

import (
	"testing"
	"time"

	"github.com/mingzaily/mconv"
)

type TestPerson struct {
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	IsActive bool      `json:"is_active"`
	Birthday time.Time `json:"birthday"`
}

func TestToJSON(t *testing.T) {
	// Create test data
	birthday, _ := time.Parse("2006-01-02", "1990-01-01")
	person := TestPerson{
		Name:     "John Doe",
		Age:      30,
		IsActive: true,
		Birthday: birthday,
	}

	// Test struct to JSON conversion
	json := mconv.ToJSON(person)
	expected := `{"name":"John Doe","age":30,"is_active":true,"birthday":"1990-01-01T00:00:00Z"}`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}

	// Test map to JSON conversion
	personMap := map[string]interface{}{
		"name":      "Jane Doe",
		"age":       25,
		"is_active": false,
		"birthday":  birthday,
	}
	json = mconv.ToJSON(personMap)
	expected = `{"age":25,"birthday":"1990-01-01T00:00:00Z","is_active":false,"name":"Jane Doe"}`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}

	// Test slice to JSON conversion
	slice := []interface{}{1, "two", true}
	json = mconv.ToJSON(slice)
	expected = `[1,"two",true]`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}

	// Test basic type to JSON conversion
	json = mconv.ToJSON(123)
	expected = `123`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}

	json = mconv.ToJSON("hello")
	expected = `"hello"`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}

	json = mconv.ToJSON(true)
	expected = `true`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}

	json = mconv.ToJSON(nil)
	expected = `null`
	if json != expected {
		t.Errorf("mconv.ToJSON() = %v; want %v", json, expected)
	}
}

func TestToJSONE(t *testing.T) {
	// Create test data
	birthday, _ := time.Parse("2006-01-02", "1990-01-01")
	person := TestPerson{
		Name:     "John Doe",
		Age:      30,
		IsActive: true,
		Birthday: birthday,
	}

	// Test struct to JSON conversion
	json, err := mconv.ToJSONE(person)
	if err != nil {
		t.Errorf("mconv.ToJSONE() unexpected error: %v", err)
	}
	expected := `{"name":"John Doe","age":30,"is_active":true,"birthday":"1990-01-01T00:00:00Z"}`
	if json != expected {
		t.Errorf("mconv.ToJSONE() = %v; want %v", json, expected)
	}

	// Test error cases
	// Create a value that cannot be serialized to JSON
	badValue := make(chan int)
	_, err = mconv.ToJSONE(badValue)
	if err == nil {
		t.Errorf("mconv.ToJSONE() with bad value expected error")
	}
}

func TestFromJSON(t *testing.T) {
	// Create test JSON
	json := `{"name":"John Doe","age":30,"is_active":true,"birthday":"1990-01-01T00:00:00Z"}`

	// Test JSON to struct conversion
	var person TestPerson
	mconv.FromJSON(json, &person)

	if person.Name != "John Doe" {
		t.Errorf("mconv.FromJSON() Name = %v; want %v", person.Name, "John Doe")
	}
	if person.Age != 30 {
		t.Errorf("mconv.FromJSON() Age = %v; want %v", person.Age, 30)
	}
	if person.IsActive != true {
		t.Errorf("mconv.FromJSON() IsActive = %v; want %v", person.IsActive, true)
	}

	expectedBirthday, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00Z")
	if !person.Birthday.Equal(expectedBirthday) {
		t.Errorf("mconv.FromJSON() Birthday = %v; want %v", person.Birthday, expectedBirthday)
	}
}

func TestFromJSONE(t *testing.T) {
	// Create test JSON
	json := `{"name":"John Doe","age":30,"is_active":true,"birthday":"1990-01-01T00:00:00Z"}`

	// Test JSON to struct conversion
	var person TestPerson
	err := mconv.FromJSONE(json, &person)
	if err != nil {
		t.Errorf("mconv.FromJSONE() unexpected error: %v", err)
	}

	if person.Name != "John Doe" {
		t.Errorf("mconv.FromJSONE() Name = %v; want %v", person.Name, "John Doe")
	}
	if person.Age != 30 {
		t.Errorf("mconv.FromJSONE() Age = %v; want %v", person.Age, 30)
	}
	if person.IsActive != true {
		t.Errorf("mconv.FromJSONE() IsActive = %v; want %v", person.IsActive, true)
	}

	expectedBirthday, _ := time.Parse(time.RFC3339, "1990-01-01T00:00:00Z")
	if !person.Birthday.Equal(expectedBirthday) {
		t.Errorf("mconv.FromJSONE() Birthday = %v; want %v", person.Birthday, expectedBirthday)
	}

	// Test error cases
	invalidJSON := `{"name":"John Doe",`
	err = mconv.FromJSONE(invalidJSON, &person)
	if err == nil {
		t.Errorf("mconv.FromJSONE() with invalid JSON expected error")
	}
}

func TestToMapFromJSON(t *testing.T) {
	// Create test JSON
	json := `{"name":"John Doe","age":30,"is_active":true,"birthday":"1990-01-01T00:00:00Z"}`

	// Test JSON to map conversion
	result := mconv.ToMapFromJSON(json)

	if result["name"] != "John Doe" {
		t.Errorf("ToMapFromJSON() name = %v; want %v", result["name"], "John Doe")
	}
	if result["age"] != float64(30) { // JSON numbers are parsed as float64
		t.Errorf("ToMapFromJSON() age = %v; want %v", result["age"], float64(30))
	}
	if result["is_active"] != true {
		t.Errorf("ToMapFromJSON() is_active = %v; want %v", result["is_active"], true)
	}
	if result["birthday"] != "1990-01-01T00:00:00Z" {
		t.Errorf("ToMapFromJSON() birthday = %v; want %v", result["birthday"], "1990-01-01T00:00:00Z")
	}
}

func TestToMapFromJSONE(t *testing.T) {
	// Create test JSON
	json := `{"name":"John Doe","age":30,"is_active":true,"birthday":"1990-01-01T00:00:00Z"}`

	// Test JSON to map conversion
	result, err := mconv.ToMapFromJSONE(json)
	if err != nil {
		t.Errorf("ToMapFromJSONE() error = %v", err)
		return
	}

	if result["name"] != "John Doe" {
		t.Errorf("ToMapFromJSONE() name = %v; want %v", result["name"], "John Doe")
	}
	if result["age"] != float64(30) { // JSON numbers are parsed as float64
		t.Errorf("ToMapFromJSONE() age = %v; want %v", result["age"], float64(30))
	}
	if result["is_active"] != true {
		t.Errorf("ToMapFromJSONE() is_active = %v; want %v", result["is_active"], true)
	}
	if result["birthday"] != "1990-01-01T00:00:00Z" {
		t.Errorf("ToMapFromJSONE() birthday = %v; want %v", result["birthday"], "1990-01-01T00:00:00Z")
	}
}

func TestToSliceFromJSON(t *testing.T) {
	// Create test JSON
	json := `["John","Jane","Bob"]`

	// Test JSON to slice conversion
	result := mconv.ToSliceFromJSON(json)

	if len(result) != 3 {
		t.Errorf("ToSliceFromJSON() length = %v; want %v", len(result), 3)
	}
	if result[0] != "John" {
		t.Errorf("ToSliceFromJSON()[0] = %v; want %v", result[0], "John")
	}
	if result[1] != "Jane" {
		t.Errorf("ToSliceFromJSON()[1] = %v; want %v", result[1], "Jane")
	}
	if result[2] != "Bob" {
		t.Errorf("ToSliceFromJSON()[2] = %v; want %v", result[2], "Bob")
	}
}

func TestToSliceFromJSONE(t *testing.T) {
	// Create test JSON
	json := `["John","Jane","Bob"]`

	// Test JSON to slice conversion
	result, err := mconv.ToSliceFromJSONE(json)
	if err != nil {
		t.Errorf("ToSliceFromJSONE() error = %v", err)
		return
	}

	if len(result) != 3 {
		t.Errorf("ToSliceFromJSONE() length = %v; want %v", len(result), 3)
	}
	if result[0] != "John" {
		t.Errorf("ToSliceFromJSONE()[0] = %v; want %v", result[0], "John")
	}
	if result[1] != "Jane" {
		t.Errorf("ToSliceFromJSONE()[1] = %v; want %v", result[1], "Jane")
	}
	if result[2] != "Bob" {
		t.Errorf("ToSliceFromJSONE()[2] = %v; want %v", result[2], "Bob")
	}
}
