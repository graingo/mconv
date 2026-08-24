package internal

import (
	"math"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var stringCache = struct {
	sync.RWMutex
	values map[scalarCacheKey]string
}{
	values: make(map[scalarCacheKey]string),
}

var stringCacheSize int64

type scalarCacheKey struct {
	kind  reflect.Kind
	text  string
	first uint64
}

func basicCacheKey(value interface{}) (scalarCacheKey, bool) {
	switch v := value.(type) {
	case string:
		return scalarCacheKey{kind: reflect.String, text: v}, true
	case bool:
		var encoded uint64
		if v {
			encoded = 1
		}
		return scalarCacheKey{kind: reflect.Bool, first: encoded}, true
	case int:
		return signedCacheKey(int64(v)), true
	case int64:
		return signedCacheKey(v), true
	case int32:
		return signedCacheKey(int64(v)), true
	case int16:
		return signedCacheKey(int64(v)), true
	case int8:
		return signedCacheKey(int64(v)), true
	case uint:
		return unsignedCacheKey(uint64(v)), true
	case uint64:
		return unsignedCacheKey(v), true
	case uint32:
		return unsignedCacheKey(uint64(v)), true
	case uint16:
		return unsignedCacheKey(uint64(v)), true
	case uint8:
		return unsignedCacheKey(uint64(v)), true
	case float64:
		return scalarCacheKey{kind: reflect.Float64, first: math.Float64bits(v)}, true
	case float32:
		return scalarCacheKey{kind: reflect.Float32, first: uint64(math.Float32bits(v))}, true
	default:
		return scalarCacheKey{}, false
	}
}

func signedCacheKey(value int64) scalarCacheKey {
	return scalarCacheKey{kind: reflect.Int64, first: uint64(value)}
}

func unsignedCacheKey(value uint64) scalarCacheKey {
	return scalarCacheKey{kind: reflect.Uint64, first: value}
}

// AddStringToCache adds a string conversion result to the bounded cache.
func AddStringToCache(value interface{}, result string) {
	if atomic.LoadInt64(&stringCacheSize) <= 0 {
		return
	}
	key, ok := basicCacheKey(value)
	if !ok {
		return
	}
	stringCache.Lock()
	defer stringCache.Unlock()
	size := atomic.LoadInt64(&stringCacheSize)
	if size <= 0 {
		return
	}
	if _, exists := stringCache.values[key]; !exists && int64(len(stringCache.values)) >= size {
		stringCache.values = make(map[scalarCacheKey]string)
	}
	stringCache.values[key] = result
}

// GetStringFromCache gets a string conversion result from the cache.
func GetStringFromCache(value interface{}) (string, bool) {
	if atomic.LoadInt64(&stringCacheSize) <= 0 {
		return "", false
	}
	key, ok := basicCacheKey(value)
	if !ok {
		return "", false
	}
	stringCache.RLock()
	defer stringCache.RUnlock()
	if atomic.LoadInt64(&stringCacheSize) <= 0 {
		return "", false
	}
	result, exists := stringCache.values[key]
	return result, exists
}

// ClearStringCache clears the string conversion cache.
func ClearStringCache() {
	stringCache.Lock()
	stringCache.values = make(map[scalarCacheKey]string)
	stringCache.Unlock()
}

// SetStringCacheSize sets the string cache size. A non-positive size disables it.
func SetStringCacheSize(size int) {
	if size < 0 {
		size = 0
	}
	stringCache.Lock()
	atomic.StoreInt64(&stringCacheSize, int64(size))
	stringCache.values = make(map[scalarCacheKey]string)
	stringCache.Unlock()
}

var timeCache = struct {
	sync.RWMutex
	values map[timeConversionCacheKey]time.Time
}{
	values: make(map[timeConversionCacheKey]time.Time),
}

var timeCacheSize int64

type timeConversionCacheKey struct {
	value        scalarCacheKey
	formatCount  int
	singleFormat string
	formats      string
}

func timeCacheKey(value interface{}, formats []string) (timeConversionCacheKey, bool) {
	key, ok := basicCacheKey(value)
	if !ok {
		return timeConversionCacheKey{}, false
	}
	result := timeConversionCacheKey{value: key, formatCount: len(formats)}
	switch len(formats) {
	case 0:
		return result, true
	case 1:
		result.singleFormat = formats[0]
		return result, true
	default:
		var builder strings.Builder
		for _, format := range formats {
			builder.WriteString(strconv.Itoa(len(format)))
			builder.WriteByte(':')
			builder.WriteString(format)
		}
		result.formats = builder.String()
		return result, true
	}
}

// AddTimeToCache adds a time conversion using the default parsing formats.
func AddTimeToCache(value interface{}, result time.Time) {
	AddTimeToCacheWithFormats(value, nil, result)
}

// AddTimeToCacheWithFormats adds a time conversion result with its parsing formats.
func AddTimeToCacheWithFormats(value interface{}, formats []string, result time.Time) {
	if atomic.LoadInt64(&timeCacheSize) <= 0 {
		return
	}
	key, ok := timeCacheKey(value, formats)
	if !ok {
		return
	}
	timeCache.Lock()
	defer timeCache.Unlock()
	size := atomic.LoadInt64(&timeCacheSize)
	if size <= 0 {
		return
	}
	if _, exists := timeCache.values[key]; !exists && int64(len(timeCache.values)) >= size {
		timeCache.values = make(map[timeConversionCacheKey]time.Time)
	}
	timeCache.values[key] = result
}

// GetTimeFromCache gets a time conversion using the default parsing formats.
func GetTimeFromCache(value interface{}) (time.Time, bool) {
	return GetTimeFromCacheWithFormats(value, nil)
}

// GetTimeFromCacheWithFormats gets a time conversion for the specified parsing formats.
func GetTimeFromCacheWithFormats(value interface{}, formats []string) (time.Time, bool) {
	if atomic.LoadInt64(&timeCacheSize) <= 0 {
		return time.Time{}, false
	}
	key, ok := timeCacheKey(value, formats)
	if !ok {
		return time.Time{}, false
	}
	timeCache.RLock()
	defer timeCache.RUnlock()
	if atomic.LoadInt64(&timeCacheSize) <= 0 {
		return time.Time{}, false
	}
	result, exists := timeCache.values[key]
	return result, exists
}

// ClearTimeCache clears the time conversion cache.
func ClearTimeCache() {
	timeCache.Lock()
	timeCache.values = make(map[timeConversionCacheKey]time.Time)
	timeCache.Unlock()
}

// SetTimeCacheSize sets the time cache size. A non-positive size disables it.
func SetTimeCacheSize(size int) {
	if size < 0 {
		size = 0
	}
	timeCache.Lock()
	atomic.StoreInt64(&timeCacheSize, int64(size))
	timeCache.values = make(map[timeConversionCacheKey]time.Time)
	timeCache.Unlock()
}

// ClearAllCaches clears all active conversion and decoder caches.
func ClearAllCaches() {
	ClearStringCache()
	ClearTimeCache()
	ClearDecoderCache()
}
