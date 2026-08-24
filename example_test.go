package mconv_test

import (
	"errors"
	"fmt"

	"github.com/graingo/mconv"
)

func Example() {
	fmt.Println(mconv.ToString(42))
	fmt.Println(mconv.ToBool("yes"))
	fmt.Println(mconv.ToInt("invalid"))
	// Output:
	// 42
	// true
	// 0
}

func ExampleToE() {
	type user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	converted, err := mconv.ToE[user](map[string]interface{}{
		"id":   "7",
		"name": "Maltose",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d %s\n", converted.ID, converted.Name)
	// Output:
	// 7 Maltose
}

func ExampleConversionError() {
	_, err := mconv.ToInt8E(128)
	fmt.Println(errors.Is(err, mconv.ErrOverflow))

	var conversionErr *mconv.ConversionError
	if errors.As(err, &conversionErr) {
		fmt.Println(conversionErr.TargetType)
	}
	// Output:
	// true
	// int8
}
