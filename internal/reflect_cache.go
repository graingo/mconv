package internal

import (
	"reflect"
	"sync"
)

// DecoderCacheKey identifies a decoder plan by destination type.
type DecoderCacheKey struct {
	DestType reflect.Type
}

// FieldDecoder holds the decoding plan for a single struct field.
type FieldDecoder struct {
	// Field represents the struct field to be decoded into.
	Field reflect.StructField
	// Index is the field's index in the struct.
	Index []int
	// Name is the key name in the source map to look for.
	Name string
	// Depth is the number of embedded fields traversed to reach the field.
	Depth int
	// Ambiguous marks fields that collide at the same embedding depth.
	Ambiguous bool
}

// Decoder holds the complete decoding plan for a struct type.
type Decoder struct {
	// Fields is a map from the source map key to the field decoder.
	// It's used for fast lookups.
	Fields map[string]*FieldDecoder
	// FieldArr is an array of all field decoders.
	FieldArr []*FieldDecoder
}

var (
	// structDecoderCache stores the cached decoders for struct types.
	structDecoderCache sync.Map
)

// GetDecoder gets a decoder for a given type from the cache.
func GetDecoder(key DecoderCacheKey) (*Decoder, bool) {
	if v, ok := structDecoderCache.Load(key); ok {
		return v.(*Decoder), true
	}
	return nil, false
}

// SetDecoder stores a decoder in the cache.
func SetDecoder(key DecoderCacheKey, decoder *Decoder) {
	structDecoderCache.Store(key, decoder)
}

// ClearDecoderCache clears the decoder cache.
func ClearDecoderCache() {
	clearSyncMap(&structDecoderCache)
}

func clearSyncMap(cache *sync.Map) {
	cache.Range(func(key, _ interface{}) bool {
		cache.Delete(key)
		return true
	})
}
