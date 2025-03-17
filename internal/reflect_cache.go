package internal

import (
	"reflect"
	"sync"
)

// TypeInfo stores reflection information about a type
type TypeInfo struct {
	// Basic type information
	Type        reflect.Type
	Kind        reflect.Kind
	IsBasic     bool
	IsContainer bool

	// Struct field information
	Fields     map[string]reflect.StructField
	FieldNames []string

	// Method information
	Methods map[string]reflect.Method

	// Type conversion information
	ConvertibleTo map[reflect.Type]bool
	AssignableTo  map[reflect.Type]bool
}

// Type information cache
var (
	typeInfoCache     = sync.Map{}
	typeInfoCacheSize = 1000 // Maximum cache size
	typeInfoCacheLen  = 0    // Current cache length
	typeInfoCacheLock sync.Mutex
)

// GetTypeInfo gets reflection information about a type, preferably from cache
func GetTypeInfo(t reflect.Type) *TypeInfo {
	// Try to get from cache
	if info, ok := typeInfoCache.Load(t); ok {
		return info.(*TypeInfo)
	}

	// Create new type information
	info := &TypeInfo{
		Type:          t,
		Kind:          t.Kind(),
		IsBasic:       isBasicType(t),
		IsContainer:   isContainerType(t),
		Fields:        make(map[string]reflect.StructField),
		FieldNames:    make([]string, 0),
		Methods:       make(map[string]reflect.Method),
		ConvertibleTo: make(map[reflect.Type]bool),
		AssignableTo:  make(map[reflect.Type]bool),
	}

	// Fill struct field information
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.IsExported() {
				info.Fields[field.Name] = field
				info.FieldNames = append(info.FieldNames, field.Name)
			}
		}
	}

	// Fill method information
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		info.Methods[method.Name] = method
	}

	// Check cache size and store
	typeInfoCacheLock.Lock()
	if typeInfoCacheLen >= typeInfoCacheSize {
		// Cache is full, clear it
		typeInfoCache = sync.Map{}
		typeInfoCacheLen = 0
	}
	typeInfoCacheLen++
	typeInfoCacheLock.Unlock()

	typeInfoCache.Store(t, info)
	return info
}

// IsConvertibleTo checks if the type can be converted to the target type, result is cached
func (ti *TypeInfo) IsConvertibleTo(target reflect.Type) bool {
	// Check cache
	if result, ok := ti.ConvertibleTo[target]; ok {
		return result
	}

	// Calculate result and cache it
	result := ti.Type.ConvertibleTo(target)
	ti.ConvertibleTo[target] = result
	return result
}

// IsAssignableTo checks if the type can be assigned to the target type, result is cached
func (ti *TypeInfo) IsAssignableTo(target reflect.Type) bool {
	// Check cache
	if result, ok := ti.AssignableTo[target]; ok {
		return result
	}

	// Calculate result and cache it
	result := ti.Type.AssignableTo(target)
	ti.AssignableTo[target] = result
	return result
}

// Check if it's a basic type
func isBasicType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.String:
		return true
	default:
		return false
	}
}

// Check if it's a container type
func isContainerType(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map:
		return true
	default:
		return false
	}
}

// ClearTypeInfoCache clears the type information cache
func ClearTypeInfoCache() {
	typeInfoCache = sync.Map{}
	typeInfoCacheLock.Lock()
	typeInfoCacheLen = 0
	typeInfoCacheLock.Unlock()
}

// SetTypeInfoCacheSize sets the type information cache size
func SetTypeInfoCacheSize(size int) {
	if size <= 0 {
		return
	}

	typeInfoCacheLock.Lock()
	typeInfoCacheSize = size
	typeInfoCache = sync.Map{}
	typeInfoCacheLen = 0
	typeInfoCacheLock.Unlock()
}

// ConversionPair represents a source-target type pair for conversion caching
type ConversionPair struct {
	Source reflect.Type
	Target reflect.Type
}

// Type conversion cache
var (
	conversionCache     = sync.Map{}
	conversionCacheSize = 1000 // Maximum cache size
	conversionCacheLen  = 0    // Current cache length
	conversionCacheLock sync.Mutex
)

// CacheConversion caches a type conversion result
func CacheConversion(source, target reflect.Type, convertible bool) {
	pair := ConversionPair{Source: source, Target: target}

	// Check cache size
	conversionCacheLock.Lock()
	if conversionCacheLen >= conversionCacheSize {
		// Cache is full, clear it
		conversionCache = sync.Map{}
		conversionCacheLen = 0
	}
	conversionCacheLen++
	conversionCacheLock.Unlock()

	conversionCache.Store(pair, convertible)
}

// GetCachedConversion gets a cached type conversion result
func GetCachedConversion(source, target reflect.Type) (bool, bool) {
	pair := ConversionPair{Source: source, Target: target}
	if result, ok := conversionCache.Load(pair); ok {
		return result.(bool), true
	}
	return false, false
}

// ClearConversionCache clears the type conversion cache
func ClearConversionCache() {
	conversionCache = sync.Map{}
	conversionCacheLock.Lock()
	conversionCacheLen = 0
	conversionCacheLock.Unlock()
}

// SetConversionCacheSize sets the type conversion cache size
func SetConversionCacheSize(size int) {
	if size <= 0 {
		return
	}

	conversionCacheLock.Lock()
	conversionCacheSize = size
	conversionCache = sync.Map{}
	conversionCacheLen = 0
	conversionCacheLock.Unlock()
}

// ClearAllReflectCaches clears all reflection-related caches
func ClearAllReflectCaches() {
	ClearTypeInfoCache()
	ClearConversionCache()
}
