# mconv v2 migration boundary

Version 1 remains source compatible. API compatibility CI compares every pull
request and release candidate with the latest v1 tag, so v2-only cleanup cannot
enter the v1 line accidentally.

## Canonical import

New code should import the root package:

```go
import "github.com/graingo/mconv"
```

The `basic` and `complex` packages remain available throughout v1. A future v2
will make the root package the supported application API, allowing conversion
internals to evolve without duplicating the compatibility surface.

## Planned removals

The following v1 compatibility functions are no-ops and are planned for
removal in v2:

- `SetTypeInfoCacheSize`
- `SetConversionCacheSize`
- `ClearTypeInfoCache`
- `ClearConversionCache`

No replacement is required. Struct field plans are cached automatically, and
direct scalar checks do not use a conversion cache.

## Upgrade shape

When v2 is released, it will use Go semantic import versioning:

```go
import "github.com/graingo/mconv/v2"
```

Applications can prepare now by importing only the root package and removing
calls to the four deprecated functions above.

---

# mconv v2 迁移边界

v1 将继续保持源码兼容。API 兼容性 CI 会把每个 Pull Request 和发布候选版本
与最新 v1 标签比较，确保只属于 v2 的清理不会意外进入 v1。

新代码统一使用根包 `github.com/graingo/mconv`。`basic`、`complex` 在 v1
期间继续可用；v2 将根包作为应用层唯一受支持入口，让内部实现可以独立演进。

v2 计划删除 `SetTypeInfoCacheSize`、`SetConversionCacheSize`、
`ClearTypeInfoCache`、`ClearConversionCache`。这些函数在 v1 中已经是空操作，
无需替代实现。升级前移除对应调用即可。
