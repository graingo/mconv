// Package mconv converts scalar values, containers, JSON, and structs with a
// consistent error model. The root package is the canonical public API.
//
// Functions ending in E return conversion errors and are preferred for data
// from requests, configuration, storage, or other external systems. Their
// convenience counterparts return the target type's zero value on failure.
package mconv
