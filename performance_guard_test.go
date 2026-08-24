package mconv_test

import (
	"testing"

	"github.com/graingo/mconv"
)

type allocationUser struct {
	ID   int    `mconv:"id"`
	Name string `mconv:"name"`
}

var (
	allocationSliceSink  []string
	allocationMapSink    map[string]string
	allocationStructSink allocationUser
)

// TestAllocationBudgets protects stable allocation characteristics while the
// scheduled benchmark workflow records latency trends. The small allowance
// above today's baseline accommodates supported Go runtime versions.
func TestAllocationBudgets(t *testing.T) {
	var (
		sliceSource  = []int{1, 2, 3}
		mapSource    = map[string]int{"a": 1, "b": 2, "c": 3}
		structSource = map[string]interface{}{"id": 7, "name": "Maltose"}
	)

	tests := []struct {
		name           string
		maximumAllocs  float64
		conversionFunc func()
	}{
		{
			name:          "generic slice",
			maximumAllocs: 10,
			conversionFunc: func() {
				allocationSliceSink = mconv.ToSliceT[string](sliceSource)
			},
		},
		{
			name:          "generic map",
			maximumAllocs: 16,
			conversionFunc: func() {
				allocationMapSink = mconv.ToMapT[string, string](mapSource)
			},
		},
		{
			name:          "struct",
			maximumAllocs: 4,
			conversionFunc: func() {
				var target allocationUser
				if err := mconv.ToStructE(structSource, &target); err != nil {
					panic(err)
				}
				allocationStructSink = target
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			allocations := testing.AllocsPerRun(1000, test.conversionFunc)
			if allocations > test.maximumAllocs {
				t.Fatalf("allocations %.1f exceed budget %.1f", allocations, test.maximumAllocs)
			}
		})
	}
}
