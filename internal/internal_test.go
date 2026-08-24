package internal_test

import (
	"errors"
	"sync"
	"testing"
	"time"

	"github.com/graingo/mconv/internal"
)

func TestNewConversionError(t *testing.T) {
	err := internal.NewConversionError("test", "string", errors.New("test error"))
	expected := `unable to convert "test" of type string to string: test error`
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}

	err = internal.NewConversionError(nil, "string", errors.New("nil value"))
	expected = `unable to convert <nil> of type nil to string: nil value`
	if err.Error() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, err.Error())
	}
}

func TestCachesAreSafeDuringConcurrentReconfiguration(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go func(index int) {
			defer wg.Done()
			internal.SetStringCacheSize(index % 4)
			internal.AddStringToCache(index, "value")
			_, _ = internal.GetStringFromCache(index)
		}(i)
		go func(index int) {
			defer wg.Done()
			internal.SetTimeCacheSize(index % 4)
			internal.AddTimeToCache(index, time.Now())
			_, _ = internal.GetTimeFromCache(index)
		}(i)
	}
	wg.Wait()
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

func TestCacheSize(t *testing.T) {
	// String cache
	internal.ClearAllCaches()
	internal.SetStringCacheSize(0)
	internal.AddStringToCache("a", "a")
	if _, ok := internal.GetStringFromCache("a"); ok {
		t.Error("string cache should be disabled when size is 0")
	}

	internal.SetStringCacheSize(1)
	internal.AddStringToCache("a", "a")
	internal.AddStringToCache("b", "b")
	if _, ok := internal.GetStringFromCache("a"); ok {
		t.Error("string cache should have evicted 'a'")
	}
	if _, ok := internal.GetStringFromCache("b"); !ok {
		t.Error("string cache should have 'b'")
	}

	// Time cache
	internal.ClearAllCaches()
	internal.SetTimeCacheSize(0)
	internal.AddTimeToCache("a", time.Now())
	if _, ok := internal.GetTimeFromCache("a"); ok {
		t.Error("time cache should be disabled when size is 0")
	}

	internal.SetTimeCacheSize(1)
	now := time.Now()
	internal.AddTimeToCache("a", now)
	internal.AddTimeToCache("b", now)
	if _, ok := internal.GetTimeFromCache("a"); ok {
		t.Error("time cache should have evicted 'a'")
	}
	if _, ok := internal.GetTimeFromCache("b"); !ok {
		t.Error("time cache should have 'b'")
	}
	// A negative size disables the cache consistently with the string cache.
	internal.SetTimeCacheSize(-1)
	if _, ok := internal.GetTimeFromCache("b"); ok {
		t.Error("time cache should be disabled by a negative size")
	}
}

func TestTimeCache(t *testing.T) {
	internal.ClearAllCaches()
	internal.SetTimeCacheSize(2)

	now := time.Now()
	internal.AddTimeToCache("key1", now)
	internal.AddTimeToCache("key2", now.Add(time.Second))

	val, ok := internal.GetTimeFromCache("key1")
	if !ok || !val.Equal(now) {
		t.Errorf("Expected to find key1 with value %v, got %v", now, val)
	}

	// This should evict key1
	internal.AddTimeToCache("key3", now.Add(2*time.Second))

	_, ok = internal.GetTimeFromCache("key1")
	if ok {
		t.Errorf("Expected key1 to be evicted")
	}

	internal.ClearTimeCache()
	_, ok = internal.GetTimeFromCache("key2")
	if ok {
		t.Errorf("Expected key2 to be cleared")
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
