package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	input  [][]int32
	output []int32
}{
	{
		input: [][]int32{
			{1, 5},
			{1, 6},
			{3, 2},
			{1, 10},
			{1, 10},
			{1, 6},
			{2, 5},
			{3, 2},
		},
		output: []int32{0, 1},
	},
}

func TestFreqQueries(t *testing.T) {
	for x, i := range cases {
		fmt.Printf("input: %v\n", i.input)
		ans := freqQuery(i.input)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
