# mconv

[![Go Report Card](https://goreportcard.com/badge/github.com/graingo/mconv)](https://goreportcard.com/report/github.com/graingo/mconv)
[![codecov](https://codecov.io/github/graingo/mconv/branch/master/graph/badge.svg?token=WmK3x7HV4k)](https://codecov.io/github/graingo/mconv)
[![Go Reference](https://pkg.go.dev/badge/github.com/graingo/mconv.svg)](https://pkg.go.dev/github.com/graingo/mconv)
[![License](https://img.shields.io/github/license/graingo/mconv.svg)](https://github.com/graingo/mconv/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/graingo/mconv.svg)](https://github.com/graingo/mconv/releases)

[中文](README.md) | [English](README_EN.md)

`mconv` 是一个零依赖的 Go 类型转换库，提供基础类型、容器、JSON 和结构体转换。最低支持 Go 1.18。

## 安装

```bash
go get github.com/graingo/mconv
```

新代码统一导入根包 `github.com/graingo/mconv`。`basic` 和 `complex` 子包在
v1 中继续保持兼容，但不作为新的应用代码入口。v2 边界见
[V2_MIGRATION.md](V2_MIGRATION.md)。

## 快速开始

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

每个宽松入口都有返回错误的 `E` 版本：

```go
value, err := mconv.ToIntE("42")
if err != nil {
	return err
}
```

`ToInt`、`ToBool` 等宽松入口会忽略转换错误并返回目标类型零值。输入来自请求、配置或外部系统时，优先使用 `E` 版本。

## 泛型转换

根包提供统一的泛型入口：

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

忽略错误的对应入口为 `To[T]`、`ToSliceT[T]` 和 `ToMapT[K,V]`。

泛型 map 直接转换原始键，不经过字符串中转，因此可以保留任意可比较键：

```go
type Key struct{ ID int }

result, err := mconv.ToMapTE[Key, int](map[Key]string{
	{ID: 1}: "42",
})
```

## 结构体转换

`ToE[T]` 适合新代码，`ToStructE` 适合需要复用已有目标对象的场景：

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

结构体转换规则：

- 标签优先级为 `mconv`、`json`、`yaml`，支持 `-` 和逗号选项。
- 精确键匹配优先，其次进行大小写不敏感匹配。
- 外层字段优先于匿名嵌入字段；同一深度的同名字段会返回歧义错误。
- 匿名字段带显式标签时保持嵌套，不进行字段提升。
- 多级指针、用户自定义基础类型、数组、slice、map 和嵌套结构体使用相同转换语义。
- 转换先在隔离副本中完成；任何字段失败时，原目标对象保持不变。
- map 键转换发生碰撞时返回错误，避免静默覆盖数据。

## 自定义 Hook

Hook 在默认的字符串时间、字符串时长转换之后执行：

```go
hook := func(from, to reflect.Type, data any) (any, error) {
	if from.Kind() == reflect.Int && to.Kind() == reflect.String {
		return fmt.Sprintf("status-%d", data.(int)), nil
	}
	return data, nil
}

value, err := mconv.ToE[string](1, hook)
```

Hook 返回原值表示继续转换，返回新值表示交给后续 Hook 或内置转换处理，返回错误会终止整个转换。

## 错误处理

错误哨兵和 `ConversionError` 位于根包，可配合 `errors.Is`、`errors.As`：

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

嵌套错误包含完整路径，例如 `Users[0].Age`。公开哨兵包括：

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

## 缓存策略

结构体字段计划由库自动缓存，无需配置。

字符串和时间值缓存默认关闭。基准显示字符串转换本身比加锁查询缓存更快；时间缓存只在少量热点字符串被高频重复解析时有收益。确认业务输入具有低基数特征后，可以主动开启时间缓存：

```go
mconv.SetTimeCacheSize(100)
defer mconv.SetTimeCacheSize(0)
```

`SetStringCacheSize` 为兼容和实验场景保留。`SetTypeInfoCacheSize`、`SetConversionCacheSize`、`ClearTypeInfoCache`、`ClearConversionCache` 已废弃并保留为空操作，后续主版本会删除。

## 性能

以下数据来自 Apple M2，命令为 `go test -run '^$' -bench . -benchmem`。结果会随 Go 版本和机器变化：

```text
BenchmarkToString-8             ~21 ns/op       3 B/op      1 allocs/op
BenchmarkToSliceT-8            ~160 ns/op     136 B/op      7 allocs/op
BenchmarkToMapT-8              ~404 ns/op     488 B/op     12 allocs/op
BenchmarkStructConversion-8    ~318 ns/op      48 B/op      1 allocs/op
```

时间缓存的典型权衡：重复值约 `75 ns/op`，关闭缓存约 `98 ns/op`；高基数输入开启缓存约 `340 ns/op`，关闭缓存约 `109 ns/op`。请用实际输入分布决定是否开启。

## 开发与验证

```bash
go test -race ./...
go vet ./...
go test -run '^$' -bench . -benchmem
```

CI 同时验证 Go 1.18 和当前稳定版，在稳定版运行 fuzz 冒烟测试，并检查
`basic` 覆盖率、分配预算和相对最新 v1 标签的公共 API 兼容性。定时
Benchmark 会保存可下载的性能结果，用于跨提交执行 `benchstat` 对比。

## 许可证

MIT
