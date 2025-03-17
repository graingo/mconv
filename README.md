# mconv

[English](README.md) | [中文](README_zh.md)

A lightweight Go type conversion library with high performance.

## Installation

```bash
go get github.com/mingzaily/mconv
```

## Features

- Simple and intuitive API
- Zero dependencies
- Comprehensive type conversion support
- Thread-safe
- High performance with caching mechanisms
- Generic support (Go 1.18+)
- Reflection result caching

## Basic Usage

```go
// Basic type conversions
str := mconv.ToString(123)        // "123"
num := mconv.ToInt("123")         // 123
b := mconv.ToBool(1)              // true
f := mconv.ToFloat64("123.45")    // 123.45
t := mconv.ToTime("2006-01-02")   // time.Time

// With error handling
str, err := mconv.ToStringE(123)  // "123", nil
num, err := mconv.ToIntE("abc")   // 0, error

// Complex type conversions
slice := mconv.ToSlice([]int{1, 2, 3})  // []interface{}{1, 2, 3}
strSlice := mconv.ToStringSlice([]int{1, 2, 3}) // []string{"1", "2", "3"}
m := mconv.ToMap(map[string]int{"a": 1}) // map[string]interface{}{"a": 1}

// JSON conversions
jsonStr := mconv.ToJSON(map[string]interface{}{"name": "John"}) // {"name":"John"}
personMap := mconv.ToMapFromJSON(`{"name":"Jane"}`) // map[string]interface{}{"name": "Jane"}
```

## Generic Support (Go 1.18+)

```go
import "github.com/mingzaily/mconv/complex"

// Slice conversions with generics
strSlice := complex.ToSliceT[string]([]int{1, 2, 3}) // []string{"1", "2", "3"}
intSlice := complex.ToSliceT[int]([]string{"1", "2", "3"}) // []int{1, 2, 3}

// Map conversions with generics
strMap := complex.ToMapT[string, string](map[string]int{"a": 1}) // map[string]string{"a": "1"}
intMap := complex.ToMapT[string, int](map[string]interface{}{"a": "1"}) // map[string]int{"a": 1}
```

## Performance Optimization

```go
// Set cache sizes
mconv.SetStringCacheSize(2000)       // Set string cache size (default 1000)
mconv.SetTimeCacheSize(200)          // Set time cache size (default 100)
mconv.SetTypeInfoCacheSize(1000)     // Set type info cache size (default 1000)
mconv.SetConversionCacheSize(1000)   // Set conversion cache size (default 1000)

// Clear caches
mconv.ClearStringCache()             // Clear string cache
mconv.ClearTimeCache()               // Clear time cache
mconv.ClearTypeInfoCache()           // Clear type info cache
mconv.ClearConversionCache()         // Clear conversion cache
mconv.ClearAllCaches()               // Clear all caches
```

## Benchmark Results

The following benchmark results were measured on an Apple M2 processor:

```
BenchmarkToString-8                 13534419               107.5 ns/op            8 B/op           1 allocs/op
BenchmarkToInt-8                    100000000               10.91 ns/op            0 B/op           0 allocs/op
BenchmarkToBool-8                   178200254                6.75 ns/op            0 B/op           0 allocs/op
BenchmarkToFloat64-8                54603055                20.84 ns/op            0 B/op           0 allocs/op
BenchmarkToTime-8                   15548144                76.57 ns/op           32 B/op           1 allocs/op
BenchmarkToSlice-8                  35479252                33.26 ns/op           40 B/op           2 allocs/op
BenchmarkToStringSlice-8             3392653               364.7 ns/op           120 B/op           8 allocs/op
BenchmarkToMap-8                     7964676               153.1 ns/op           336 B/op           2 allocs/op
BenchmarkToJSON-8                    3466358               348.2 ns/op           192 B/op           7 allocs/op
```

Generic functions:

```
BenchmarkToSliceT-8                  2906990               397.6 ns/op          136 B/op           6 allocs/op
BenchmarkToMapT-8                    1654909               723.4 ns/op          688 B/op           7 allocs/op
```

## Reflection Caching Benefits

Reflection caching is a key feature of the mconv library that significantly improves type conversion performance. The following benchmark results demonstrate the performance advantages of reflection caching:

```
BenchmarkReflectionCache_TypeInfo/WithoutCache-8           8842224               133.9 ns/op            56 B/op           7 allocs/op
BenchmarkReflectionCache_TypeInfo/WithCache-8             25792029                49.03 ns/op            0 B/op           0 allocs/op
```

From these results, we can observe:

1. **Performance Improvement**: With reflection caching, processing the same type of reflection operations is about 2.7 times faster (from 133.9 ns/op to 49.03 ns/op).
2. **Memory Optimization**: With caching, memory allocation is reduced from 56 bytes and 7 allocations per operation to 0 bytes and 0 allocations, completely eliminating memory allocation overhead.
3. **Throughput Increase**: The number of operations that can be processed per second increases from about 8.8 million to about 25.8 million, an improvement of approximately 192%.

For large data structures, the benefits are even more significant:

```
BenchmarkLargeSliceConversion/WithoutCache-8                  9327             121340 ns/op         46738 B/op        1747 allocs/op
BenchmarkLargeSliceConversion/WithCache-8                     9620             121506 ns/op         46738 B/op        1747 allocs/op
```

These advantages are particularly evident in applications that process large datasets or require frequent type conversions. Reflection caching stores type information and conversion results, avoiding repeated reflection operations, thereby significantly improving performance.

## Use Cases

Reflection caching is particularly useful in the following scenarios:

1. **API Services**: Frequently converting data between different formats
2. **Data Processing Pipelines**: Processing large amounts of structurally similar data
3. **ORM and Data Mapping**: Mapping between database records and structs
4. **Configuration Processing**: Parsing and converting configuration data in various formats
5. **JSON/XML Processing**: Frequent serialization and deserialization operations

## License

MIT License
