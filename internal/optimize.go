package internal

import (
	"fmt"
	"sync"
	"time"
)

// String cache, used to cache common string conversion results
var (
	stringCache     = sync.Map{}
	stringCacheSize = 1000 // Maximum cache size
	stringCacheLen  = 0    // Current cache length
	stringCacheLock sync.Mutex
)

// Add string to cache
func addStringToCache(key interface{}, value string) {
	// Only cache basic types to avoid memory leaks from caching complex types
	var cacheKey string
	switch v := key.(type) {
	case string:
		cacheKey = v
	case bool:
		if v {
			cacheKey = "bool:true"
		} else {
			cacheKey = "bool:false"
		}
	case int:
		cacheKey = fmt.Sprintf("int:%d", v)
	case int64:
		cacheKey = fmt.Sprintf("int64:%d", v)
	case int32:
		cacheKey = fmt.Sprintf("int32:%d", v)
	case int16:
		cacheKey = fmt.Sprintf("int16:%d", v)
	case int8:
		cacheKey = fmt.Sprintf("int8:%d", v)
	case uint:
		cacheKey = fmt.Sprintf("uint:%d", v)
	case uint64:
		cacheKey = fmt.Sprintf("uint64:%d", v)
	case uint32:
		cacheKey = fmt.Sprintf("uint32:%d", v)
	case uint16:
		cacheKey = fmt.Sprintf("uint16:%d", v)
	case uint8:
		cacheKey = fmt.Sprintf("uint8:%d", v)
	case float64:
		cacheKey = fmt.Sprintf("float64:%g", v)
	case float32:
		cacheKey = fmt.Sprintf("float32:%g", v)
	default:
		// Do not cache other types
		return
	}

	// Check if we need to clear the cache
	stringCacheLock.Lock()
	if stringCacheLen >= stringCacheSize {
		stringCache = sync.Map{}
		stringCacheLen = 0
	}
	stringCacheLen++
	stringCacheLock.Unlock()

	// Store in cache
	stringCache.Store(cacheKey, value)
}

// Get string from cache
func getStringFromCache(key interface{}) (string, bool) {
	// Only get basic types from cache
	var cacheKey string
	switch v := key.(type) {
	case string:
		cacheKey = v
	case bool:
		if v {
			cacheKey = "bool:true"
		} else {
			cacheKey = "bool:false"
		}
	case int:
		cacheKey = fmt.Sprintf("int:%d", v)
	case int64:
		cacheKey = fmt.Sprintf("int64:%d", v)
	case int32:
		cacheKey = fmt.Sprintf("int32:%d", v)
	case int16:
		cacheKey = fmt.Sprintf("int16:%d", v)
	case int8:
		cacheKey = fmt.Sprintf("int8:%d", v)
	case uint:
		cacheKey = fmt.Sprintf("uint:%d", v)
	case uint64:
		cacheKey = fmt.Sprintf("uint64:%d", v)
	case uint32:
		cacheKey = fmt.Sprintf("uint32:%d", v)
	case uint16:
		cacheKey = fmt.Sprintf("uint16:%d", v)
	case uint8:
		cacheKey = fmt.Sprintf("uint8:%d", v)
	case float64:
		cacheKey = fmt.Sprintf("float64:%g", v)
	case float32:
		cacheKey = fmt.Sprintf("float32:%g", v)
	default:
		// Do not get other types from cache
		return "", false
	}

	// Get from cache
	if value, ok := stringCache.Load(cacheKey); ok {
		return value.(string), true
	}
	return "", false
}

// Clear string cache
func ClearStringCache() {
	stringCache = sync.Map{}
	stringCacheLock.Lock()
	stringCacheLen = 0
	stringCacheLock.Unlock()
}

// Set string cache size
func SetStringCacheSize(size int) {
	if size <= 0 {
		return
	}

	stringCacheLock.Lock()
	stringCacheSize = size
	stringCache = sync.Map{}
	stringCacheLen = 0
	stringCacheLock.Unlock()
}

// Time format cache, used to cache common time format parsing results
var (
	timeCache     = sync.Map{}
	timeCacheSize = 100 // Maximum cache size
	timeCacheLen  = 0   // Current cache length
	timeCacheLock sync.Mutex
)

// Add time to cache
func addTimeToCache(key string, value time.Time) {
	// Check if we need to clear the cache
	timeCacheLock.Lock()
	if timeCacheLen >= timeCacheSize {
		timeCache = sync.Map{}
		timeCacheLen = 0
	}
	timeCacheLen++
	timeCacheLock.Unlock()

	// Store in cache
	timeCache.Store(key, value)
}

// Get time from cache
func getTimeFromCache(key string) (time.Time, bool) {
	// Get from cache
	if value, ok := timeCache.Load(key); ok {
		return value.(time.Time), true
	}
	return time.Time{}, false
}

// Clear time cache
func ClearTimeCache() {
	timeCache = sync.Map{}
	timeCacheLock.Lock()
	timeCacheLen = 0
	timeCacheLock.Unlock()
}

// Set time cache size
func SetTimeCacheSize(size int) {
	if size <= 0 {
		return
	}

	timeCacheLock.Lock()
	timeCacheSize = size
	timeCache = sync.Map{}
	timeCacheLen = 0
	timeCacheLock.Unlock()
}

// ClearAllCaches clears all caches
func ClearAllCaches() {
	ClearStringCache()
	ClearTimeCache()
	ClearTypeInfoCache()
	ClearConversionCache()
}

// AddStringToCache adds string to cache.
func AddStringToCache(value interface{}, result string) {
	// Only cache basic types to avoid memory leaks
	switch value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool:
		// These types can be safely cached
	default:
		// Do not cache other types
		return
	}

	// Use type+value as cache key
	cacheKey := fmt.Sprintf("%T:%v", value, value)

	// Check if we need to clear the cache
	stringCacheLock.Lock()
	if stringCacheLen >= stringCacheSize {
		stringCache = sync.Map{}
		stringCacheLen = 0
	}
	stringCacheLen++
	stringCacheLock.Unlock()

	// Store in cache
	stringCache.Store(cacheKey, result)
}

// GetStringFromCache gets string from cache.
func GetStringFromCache(value interface{}) (string, bool) {
	// Only get basic types from cache
	switch value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32, bool:
		// These types can be safely cached
	default:
		// Do not get other types from cache
		return "", false
	}

	// Use type+value as cache key
	cacheKey := fmt.Sprintf("%T:%v", value, value)

	// Get from cache
	if result, ok := stringCache.Load(cacheKey); ok {
		return result.(string), true
	}
	return "", false
}

// AddTimeToCache adds time to cache.
func AddTimeToCache(value interface{}, result time.Time) {
	// Only cache basic types to avoid memory leaks
	switch value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32:
		// These types can be safely cached
	default:
		// Do not cache other types
		return
	}

	// Use type+value as cache key
	cacheKey := fmt.Sprintf("%T:%v", value, value)

	// Check if we need to clear the cache
	timeCacheLock.Lock()
	if timeCacheLen >= timeCacheSize {
		timeCache = sync.Map{}
		timeCacheLen = 0
	}
	timeCacheLen++
	timeCacheLock.Unlock()

	// Store in cache
	timeCache.Store(cacheKey, result)
}

// GetTimeFromCache gets time from cache.
func GetTimeFromCache(value interface{}) (time.Time, bool) {
	// Only get basic types from cache
	switch value.(type) {
	case string, int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float64, float32:
		// These types can be safely cached
	default:
		// Do not get other types from cache
		return time.Time{}, false
	}

	// Use type+value as cache key
	cacheKey := fmt.Sprintf("%T:%v", value, value)

	// Get from cache
	if result, ok := timeCache.Load(cacheKey); ok {
		return result.(time.Time), true
	}
	return time.Time{}, false
}
