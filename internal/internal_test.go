package internal_test

import (
	"errors"
	"testing"

	"github.com/graingo/mconv/internal"
)

func TestNewConversionError(t *testing.T) {
	err := internal.NewConversionError("test", "string", errors.New("test error"))
	expected := `unable to convert "test" of type string to string: test error`
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

func TestConversionError_Unwrap(t *testing.T) {
	originalErr := errors.New("original error")
	err := internal.NewConversionError("value", "target", originalErr)

	if !errors.Is(err, originalErr) {
		t.Errorf("Expected error to be unwrappable to the original error")
	}
}

func TestStringCache(t *testing.T) {
	internal.ClearAllCaches()
	internal.SetStringCacheSize(2)

	internal.AddStringToCache("key1", "value1")
	internal.AddStringToCache("key2", "value2")

	val, ok := internal.GetStringFromCache("key1")
	if !ok || val != "value1" {
		t.Errorf("Expected to find key1 with value 'value1', got %s", val)
	}

	// This should evict key1
	internal.AddStringToCache("key3", "value3")

	_, ok = internal.GetStringFromCache("key1")
	if ok {
		t.Errorf("Expected key1 to be evicted")
	}
}

func TestDecoderCache(t *testing.T) {
	internal.ClearAllCaches()
	key := internal.DecoderCacheKey{}
	decoder := &internal.Decoder{}
	internal.SetDecoder(key, decoder)
	retrieved, ok := internal.GetDecoder(key)
	if !ok || retrieved != decoder {
		t.Error("Decoder cache failed")
	}
	internal.ClearDecoderCache()
	_, ok = internal.GetDecoder(key)
	if ok {
		t.Error("ClearDecoderCache failed")
	}
}
