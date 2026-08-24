package complex_test

import (
	"encoding/json"
	"testing"

	"github.com/graingo/mconv"
)

func FuzzDecodedJSONConversions(f *testing.F) {
	for _, seed := range [][]byte{
		[]byte(`null`),
		[]byte(`{"name":"maltose","age":"7"}`),
		[]byte(`[1,"2",true,null]`),
		[]byte(`{"nested":{"enabled":1},"items":[{"id":"1"}]}`),
	} {
		f.Add(seed)
	}

	type Item struct {
		ID int `json:"id"`
	}
	type Nested struct {
		Enabled bool `json:"enabled"`
	}
	type Payload struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Nested Nested `json:"nested"`
		Items  []Item `json:"items"`
	}

	f.Fuzz(func(t *testing.T, data []byte) {
		var source interface{}
		if err := json.Unmarshal(data, &source); err != nil {
			return
		}

		_, _ = mconv.ToE[Payload](source)
		_, _ = mconv.ToSliceTE[string](source)
		_, _ = mconv.ToMapTE[string, interface{}](source)
	})
}
