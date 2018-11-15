package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	nums   []int
	k      int
	output int
}{
	{
		nums:   []int{1, 2, 3, 4, 5},
		k:      3,
		output: 4,
	},
	{
		nums:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		k:      5,
		output: 5,
	},
	{
		nums:   []int{1},
		k:      0,
		output: 1,
	},
	{
		nums:   []int{10, 5, 15},
		k:      1,
		output: 10,
	},
}

func TestQuickselect(t *testing.T) {
	for x, i := range cases {
		ans := quickselect(i.nums, i.k)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
