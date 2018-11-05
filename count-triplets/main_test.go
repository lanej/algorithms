package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	input  []int64
	r      int64
	output int64
}{
	{
		input:  []int64{1, 2, 2, 4},
		r:      2,
		output: 2,
	},
	{
		input:  []int64{1, 3, 9, 9, 27, 81},
		r:      6,
		output: 6,
	},
	{
		input:  []int64{1, 5, 5, 25, 125},
		r:      5,
		output: 4,
	},
	{
		r: 3,
		// 1 2 3
		input:  []int64{1, 1, 1},
		output: 1,
	},
	{
		r: 4,
		// 1 2 3
		// 1 2 4
		// 1 3 4
		// 2 3 4
		input:  []int64{1, 1, 1, 1},
		output: 4,
	},
	{
		r: 5,
		// 1 2 3
		// 1 2 4
		// 1 2 5
		// 1 3 4
		// 1 3 5
		// 1 4 5
		// 2 3 4
		// 2 3 5
		// 2 4 5
		// 3 4 5
		input:  []int64{1, 1, 1, 1, 1},
		output: 9,
	},
	{
		r: 6,
		// 1 2 3
		// 1 2 4
		// 1 2 5
		// 1 2 6
		// 1 3 4
		// 1 3 5
		// 1 3 6
		// 1 4 5
		// 1 4 6
		// 1 5 6
		// 2 3 4
		// 2 3 5
		// 2 3 6
		// 2 4 5
		// 2 4 6
		// 2 5 6
		// 3 4 5
		// 3 4 6
		// 3 5 6
		// 4 5 6
		input:  []int64{1, 1, 1, 1, 1, 1},
		output: 20,
	},
	{
		r:      100,
		input:  []int64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		output: 161700,
	},
}

func TestCountTriplets(t *testing.T) {
	for x, i := range cases {
		fmt.Printf("input: %v\n", i.input)
		ans := countTriplets(i.input, i.r)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
