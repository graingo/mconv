# mconv

[![Go Report Card](https://goreportcard.com/badge/github.com/graingo/mconv)](https://goreportcard.com/report/github.com/graingo/mconv)
[![codecov](https://codecov.io/github/graingo/mconv/branch/master/graph/badge.svg?token=WmK3x7HV4k)](https://codecov.io/github/graingo/mconv)
[![Go Reference](https://pkg.go.dev/badge/github.com/graingo/mconv.svg)](https://pkg.go.dev/github.com/graingo/mconv)
[![License](https://img.shields.io/github/license/graingo/mconv.svg)](https://github.com/graingo/mconv/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/graingo/mconv.svg)](https://github.com/graingo/mconv/releases)

[中文](README.md) | [English](README_EN.md)

`mconv` is a dependency-free Go conversion library for scalar values, containers, JSON, and structs. It supports Go 1.18 and later.

## Installation

```bash
go get github.com/graingo/mconv
```

New code should import the canonical root package `github.com/graingo/mconv`.
The `basic` and `complex` subpackages remain available for existing code, while
the root package is the recommended application entry point.

## Quick start

```go
package main

import (
	"fmt"

	"github.com/graingo/mconv"
)

func main() {
	fmt.Println(mconv.ToString(42))       // 42
	fmt.Println(mconv.ToInt("42"))       // 42
	fmt.Println(mconv.ToBool("yes"))     // true
	fmt.Println(mconv.ToDuration("1m"))  // 1m0s
}
```

Every permissive entry point has an `E` variant that returns an error:

```go
value, err := mconv.ToIntE("42")
if err != nil {
	return err
}
```

Permissive functions such as `ToInt` and `ToBool` ignore conversion errors and return the target type's zero value. Prefer the `E` variants for request data, configuration, and other external input.

## Generic conversion

The root package provides unified generic entry points:

```go
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

user, err := mconv.ToE[User](map[string]any{
	"id":   "7",
	"name": "maltose",
})

numbers, err := mconv.ToSliceTE[int]([]string{"1", "2", "3"})
labels, err := mconv.ToMapTE[int, string](map[string]int{"1": 10})
```

The permissive counterparts are `To[T]`, `ToSliceT[T]`, and `ToMapT[K,V]`.

Generic map conversion works directly with source keys instead of converting them through strings, so any comparable key can be preserved:

```go
type Key struct{ ID int }

result, err := mconv.ToMapTE[Key, int](map[Key]string{
	{ID: 1}: "42",
})
```

## Struct conversion

Use `ToE[T]` in new code. Use `ToStructE` when an existing destination value should be reused:

```go
type Profile struct {
	Name      string `mconv:"display_name"`
	Age       int    `json:"age"`
	Enabled   bool   `yaml:"enabled"`
	CreatedAt time.Time
}

profile := Profile{Enabled: true}
err := mconv.ToStructE(map[string]any{
	"display_name": "Alice",
	"age":          "18",
	"CreatedAt":    "2026-08-24T12:00:00Z",
}, &profile)
```

Struct conversion rules:

- Tag priority is `mconv`, `json`, then `yaml`; ignored fields and comma options are supported.
- Exact key matches win, followed by case-insensitive matching.
- Shallower fields win over anonymous embedded fields. A collision at the same depth returns an ambiguity error.
- An anonymous field with an explicit tag remains nested instead of being promoted.
- Multi-level pointers, defined scalar types, arrays, slices, maps, and nested structs share the same conversion semantics.
- Conversion finishes in an isolated copy. A field error leaves the original destination unchanged.
- A map-key collision after conversion returns an error instead of silently overwriting data.

## Custom hooks

Custom hooks run after the built-in string-to-time and string-to-duration hooks:

```go
hook := func(from, to reflect.Type, data any) (any, error) {
	if from.Kind() == reflect.Int && to.Kind() == reflect.String {
		return fmt.Sprintf("status-%d", data.(int)), nil
	}
	return data, nil
}

value, err := mconv.ToE[string](1, hook)
```

Return the original value to continue normal conversion, a new value to pass it to later hooks or built-in conversion, or an error to stop the complete operation.

## Error handling

Error sentinels and `ConversionError` are available from the root package and work with `errors.Is` and `errors.As`:

```go
_, err := mconv.ToInt8E(128)
if errors.Is(err, mconv.ErrOverflow) {
	// handle overflow
}

var conversionErr *mconv.ConversionError
if errors.As(err, &conversionErr) {
	fmt.Println(conversionErr.TargetType)
	fmt.Println(conversionErr.Path)
}
```

Nested failures include a complete path such as `Users[0].Age`. Public sentinels include:

- `ErrUnsupportedType`
- `ErrConversionFailed`
- `ErrOverflow`
- `ErrInvalidTimeFormat`
- `ErrInvalidJSONFormat`

## JSON

```go
jsonText, err := mconv.ToJSONE(map[string]any{"name": "maltose"})

var user User
err = mconv.FromJSONE(jsonText, &user)

data, err := mconv.ToMapFromJSONE(jsonText)
```

## Cache policy

The library automatically caches struct field plans; no configuration is required.

String and time value caches are disabled by default. Benchmarks show that scalar-to-string conversion is cheaper than a synchronized cache lookup. The time cache helps only when a small set of strings is parsed repeatedly. Enable it after confirming that the workload has low-cardinality time input:

```go
mconv.SetTimeCacheSize(100)
defer mconv.SetTimeCacheSize(0)
```

`SetStringCacheSize` remains available for compatibility and experiments. `SetTypeInfoCacheSize`, `SetConversionCacheSize`, `ClearTypeInfoCache`, and `ClearConversionCache` are deprecated no-ops retained for source compatibility.

## Performance

The following results were measured on an Apple M2 with `go test -run '^$' -bench . -benchmem`. Results vary by Go version and machine:

```text
BenchmarkToString-8             ~21 ns/op       3 B/op      1 allocs/op
BenchmarkToSliceT-8            ~160 ns/op     136 B/op      7 allocs/op
BenchmarkToMapT-8              ~404 ns/op     488 B/op     12 allocs/op
BenchmarkStructConversion-8    ~318 ns/op      48 B/op      1 allocs/op
```

A repeated time string takes about `75 ns/op` with the time cache and `98 ns/op` without it. High-cardinality input takes about `340 ns/op` with the cache and `109 ns/op` without it. Choose based on the real input distribution.

## Development

```bash
go test -race ./...
go vet ./...
go test -run '^$' -bench . -benchmem
```

CI verifies both Go 1.18 and the current stable release. On stable Go it also
runs short fuzz smoke tests, enforces `basic` coverage and allocation budgets,
and compares the public API with the latest v1 tag. Scheduled benchmarks retain
downloadable results for cross-commit `benchstat` comparisons.

## License

MIT
