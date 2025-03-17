package basic

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
	"time"

	"github.com/mingzaily/mconv/internal"
)

// ToStringE converts any type to string with error
func ToStringE(value interface{}) (string, error) {
	if value == nil {
		return "", nil
	}

	if cachedValue, ok := internal.GetStringFromCache(value); ok {
		return cachedValue, nil
	}

	var result string
	switch v := value.(type) {
	case string:
		result = v
	case bool:
		result = strconv.FormatBool(v)
	case float64:
		result = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		result = strconv.FormatFloat(float64(v), 'f', -1, 32)
	case int:
		result = strconv.Itoa(v)
	case int64:
		result = strconv.FormatInt(v, 10)
	case int32:
		result = strconv.FormatInt(int64(v), 10)
	case int16:
		result = strconv.FormatInt(int64(v), 10)
	case int8:
		result = strconv.FormatInt(int64(v), 10)
	case uint:
		result = strconv.FormatUint(uint64(v), 10)
	case uint64:
		result = strconv.FormatUint(v, 10)
	case uint32:
		result = strconv.FormatUint(uint64(v), 10)
	case uint16:
		result = strconv.FormatUint(uint64(v), 10)
	case uint8:
		result = strconv.FormatUint(uint64(v), 10)
	case complex64:
		result = fmt.Sprintf("%v", v)
	case complex128:
		result = fmt.Sprintf("%v", v)
	case []byte:
		result = string(v)
	case template.HTML:
		result = string(v)
	case template.URL:
		result = string(v)
	case template.JS:
		result = string(v)
	case template.CSS:
		result = string(v)
	case template.HTMLAttr:
		result = string(v)
	case json.Number:
		result = string(v)
	case time.Time:
		result = v.Format(time.RFC3339)
	case fmt.Stringer:
		result = v.String()
	default:
		result = fmt.Sprintf("%v", value)
	}

	internal.AddStringToCache(value, result)

	return result, nil
}

// ToString converts any type to string
func ToString(value interface{}) string {
	result, _ := ToStringE(value)
	return result
}
