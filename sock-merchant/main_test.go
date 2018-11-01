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
		n:      9,
		input:  []int32{10, 20, 20, 10, 10, 30, 50, 10, 20},
		output: 3,
	},
}

func TestAnagrams(t *testing.T) {
	for x, i := range cases {
		ans := sockMerchant(i.n, i.input)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
