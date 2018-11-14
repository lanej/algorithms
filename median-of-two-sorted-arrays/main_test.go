package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	A      []int
	B      []int
	output float64
}{
	{
		A:      []int{1, 2, 3, 4, 5},
		B:      []int{6, 7, 8, 9, 10},
		output: 5.5,
	},
	{
		A:      []int{6, 7, 8, 9, 10},
		B:      []int{1, 2, 3, 4, 5},
		output: 5.5,
	},
	{
		A:      []int{1, 3, 5, 6, 9},
		B:      []int{2, 4, 6, 8, 10},
		output: 5.5,
	},
	{
		A:      []int{6, 7, 8, 9, 10},
		B:      []int{15},
		output: 8.5,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for x, i := range cases {
		ans := findMedianSortedArrays(i.A, i.B)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
