package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	n      int32
	input  []int32
	output int32
}{
	{
		n:      7,
		input:  []int32{0, 0, 1, 0, 0, 1, 0},
		output: 4,
	},
	{
		n:      6,
		input:  []int32{0, 0, 0, 0, 1, 0},
		output: 3,
	},
	{
		n:      6,
		input:  []int32{0, 0, 0, 1, 0, 0},
		output: 3,
	},
}

func TestJumpingOnClouds(t *testing.T) {
	for x, i := range cases {
		ans := jumpingOnClouds(i.input)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
