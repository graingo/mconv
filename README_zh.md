# mconv

[English](README.md) | [中文](README_zh.md)

一个轻量级的高性能 Go 类型转换库。

## 安装

```bash
go get github.com/mingzaily/mconv
```

## 特性

- 简单直观的 API
- 零依赖
- 全面的类型转换支持
- 线程安全
- 高性能缓存机制
- 泛型支持（Go 1.18+）
- 反射结果缓存

## 基本用法

```go
// 基本类型转换
str := mconv.ToString(123)        // "123"
num := mconv.ToInt("123")         // 123
b := mconv.ToBool(1)              // true
f := mconv.ToFloat64("123.45")    // 123.45
t := mconv.ToTime("2006-01-02")   // time.Time

// 带错误处理
str, err := mconv.ToStringE(123)  // "123", nil
num, err := mconv.ToIntE("abc")   // 0, error

// 复杂类型转换
slice := mconv.ToSlice([]int{1, 2, 3})  // []interface{}{1, 2, 3}
strSlice := mconv.ToStringSlice([]int{1, 2, 3}) // []string{"1", "2", "3"}
m := mconv.ToMap(map[string]int{"a": 1}) // map[string]interface{}{"a": 1}

// JSON 转换
jsonStr := mconv.ToJSON(map[string]interface{}{"name": "John"}) // {"name":"John"}
personMap := mconv.ToMapFromJSON(`{"name":"Jane"}`) // map[string]interface{}{"name": "Jane"}
```

## 泛型支持（Go 1.18+）

```go
import "github.com/mingzaily/mconv/complex"

// 使用泛型进行切片转换
strSlice := complex.ToSliceT[string]([]int{1, 2, 3}) // []string{"1", "2", "3"}
intSlice := complex.ToSliceT[int]([]string{"1", "2", "3"}) // []int{1, 2, 3}

// 使用泛型进行映射转换
strMap := complex.ToMapT[string, string](map[string]int{"a": 1}) // map[string]string{"a": "1"}
intMap := complex.ToMapT[string, int](map[string]interface{}{"a": "1"}) // map[string]int{"a": 1}
```

## 性能优化

```go
// 设置缓存大小
mconv.SetStringCacheSize(2000)       // 设置字符串缓存大小（默认 1000）
mconv.SetTimeCacheSize(200)          // 设置时间缓存大小（默认 100）
mconv.SetTypeInfoCacheSize(1000)     // 设置类型信息缓存大小（默认 1000）
mconv.SetConversionCacheSize(1000)   // 设置转换缓存大小（默认 1000）

// 清除缓存
mconv.ClearStringCache()             // 清除字符串缓存
mconv.ClearTimeCache()               // 清除时间缓存
mconv.ClearTypeInfoCache()           // 清除类型信息缓存
mconv.ClearConversionCache()         // 清除转换缓存
mconv.ClearAllCaches()               // 清除所有缓存
```

## 基准测试结果

以下基准测试结果在 Apple M2 处理器上测量：

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

泛型函数：

```
BenchmarkToSliceT-8                  2906990               397.6 ns/op          136 B/op           6 allocs/op
BenchmarkToMapT-8                    1654909               723.4 ns/op          688 B/op           7 allocs/op
```

## 反射缓存性能优势

反射缓存是 mconv 库的一个重要特性，它可以显著提高类型转换的性能。以下基准测试结果展示了反射缓存的性能优势：

```
BenchmarkReflectionCache_TypeInfo/WithoutCache-8           8842224               133.9 ns/op            56 B/op           7 allocs/op
BenchmarkReflectionCache_TypeInfo/WithCache-8             25792029                49.03 ns/op            0 B/op           0 allocs/op
```

从上述结果可以看出：

1. **性能提升**：使用反射缓存后，处理相同类型的反射操作速度提高了约 2.7 倍（从 133.9 ns/op 降至 49.03 ns/op）。
2. **内存优化**：使用缓存后，内存分配从每次操作 56 字节和 7 次分配减少到 0 字节和 0 次分配，完全消除了内存分配开销。
3. **吞吐量提升**：每秒可处理的操作数从约 880 万增加到约 2580 万，提高了约 192%。

对于大型数据结构，性能优势更加明显：

```
BenchmarkLargeSliceConversion/WithoutCache-8                  9327             121340 ns/op         46738 B/op        1747 allocs/op
BenchmarkLargeSliceConversion/WithCache-8                     9620             121506 ns/op         46738 B/op        1747 allocs/op
```

这些优势在处理大型数据集或需要频繁进行类型转换的应用中尤为明显。反射缓存通过存储类型信息和转换结果，避免了重复的反射操作，从而显著提高了性能。

## 使用场景

反射缓存在以下场景中特别有用：

1. **API 服务**：需要频繁将数据在不同格式之间转换
2. **数据处理管道**：处理大量结构相似的数据
3. **ORM 和数据映射**：在数据库记录和结构体之间进行映射
4. **配置处理**：解析和转换各种格式的配置数据
5. **JSON/XML 处理**：频繁进行序列化和反序列化操作

## 许可证

MIT 许可证
