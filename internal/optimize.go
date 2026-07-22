package internal

import (
	"fmt"
	"sync"
	"time"
)

var stringCache = struct {
	sync.RWMutex
	values map[string]string
	size   int
}{
	values: make(map[string]string),
	size:   1000,
}

func basicCacheKey(value interface{}) (string, bool) {
	switch value.(type) {
	case string, int, int64, int32, int16, int8,
		uint, uint64, uint32, uint16, uint8,
		float64, float32, bool:
		return fmt.Sprintf("%T:%v", value, value), true
	default:
		return "", false
	}
}

// AddStringToCache adds a string conversion result to the bounded cache.
func AddStringToCache(value interface{}, result string) {
	key, ok := basicCacheKey(value)
	if !ok {
		return
	}
	stringCache.Lock()
	defer stringCache.Unlock()
	if stringCache.size <= 0 {
		return
	}
	if _, exists := stringCache.values[key]; !exists && len(stringCache.values) >= stringCache.size {
		stringCache.values = make(map[string]string)
	}
	stringCache.values[key] = result
}

// GetStringFromCache gets a string conversion result from the cache.
func GetStringFromCache(value interface{}) (string, bool) {
	key, ok := basicCacheKey(value)
	if !ok {
		return "", false
	}
	stringCache.RLock()
	defer stringCache.RUnlock()
	if stringCache.size <= 0 {
		return "", false
	}
	result, exists := stringCache.values[key]
	return result, exists
}

// ClearStringCache clears the string conversion cache.
func ClearStringCache() {
	stringCache.Lock()
	stringCache.values = make(map[string]string)
	stringCache.Unlock()
}

// SetStringCacheSize sets the string cache size. A non-positive size disables it.
func SetStringCacheSize(size int) {
	stringCache.Lock()
	stringCache.size = size
	if size < 0 {
		stringCache.size = 0
	}
	stringCache.values = make(map[string]string)
	stringCache.Unlock()
}

var timeCache = struct {
	sync.RWMutex
	values map[string]time.Time
	size   int
}{
	values: make(map[string]time.Time),
	size:   100,
}

func timeCacheKey(value interface{}, formats []string) (string, bool) {
	key, ok := basicCacheKey(value)
	if !ok {
		return "", false
	}
	return fmt.Sprintf("%s|formats:%q", key, formats), true
}

// AddTimeToCache adds a time conversion using the default parsing formats.
func AddTimeToCache(value interface{}, result time.Time) {
	AddTimeToCacheWithFormats(value, nil, result)
}

// AddTimeToCacheWithFormats adds a time conversion result with its parsing formats.
func AddTimeToCacheWithFormats(value interface{}, formats []string, result time.Time) {
	key, ok := timeCacheKey(value, formats)
	if !ok {
		return
	}
	timeCache.Lock()
	defer timeCache.Unlock()
	if timeCache.size <= 0 {
		return
	}
	if _, exists := timeCache.values[key]; !exists && len(timeCache.values) >= timeCache.size {
		timeCache.values = make(map[string]time.Time)
	}
	timeCache.values[key] = result
}

// GetTimeFromCache gets a time conversion using the default parsing formats.
func GetTimeFromCache(value interface{}) (time.Time, bool) {
	return GetTimeFromCacheWithFormats(value, nil)
}

// GetTimeFromCacheWithFormats gets a time conversion for the specified parsing formats.
func GetTimeFromCacheWithFormats(value interface{}, formats []string) (time.Time, bool) {
	key, ok := timeCacheKey(value, formats)
	if !ok {
		return time.Time{}, false
	}
	timeCache.RLock()
	defer timeCache.RUnlock()
	if timeCache.size <= 0 {
		return time.Time{}, false
	}
	result, exists := timeCache.values[key]
	return result, exists
}

// ClearTimeCache clears the time conversion cache.
func ClearTimeCache() {
	timeCache.Lock()
	timeCache.values = make(map[string]time.Time)
	timeCache.Unlock()
}

// SetTimeCacheSize sets the time cache size. A zero size disables it; negative values are ignored.
func SetTimeCacheSize(size int) {
	if size < 0 {
		return
	}
	timeCache.Lock()
	timeCache.size = size
	timeCache.values = make(map[string]time.Time)
	timeCache.Unlock()
}

// ClearAllCaches clears all conversion, reflection, and decoder caches.
func ClearAllCaches() {
	ClearStringCache()
	ClearTimeCache()
	ClearTypeInfoCache()
	ClearConversionCache()
	ClearDecoderCache()
}
