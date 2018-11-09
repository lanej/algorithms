package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	input []int32
	swaps int32
	first int32
	last  int32
}{
	{
		input: []int32{1, 2, 3},
		swaps: 0,
		first: 1,
		last:  3,
	},
}

func TestFreqQueries(t *testing.T) {
	for x, i := range cases {
		fmt.Printf("input: %v\n", i.input)
		s, f, l := countSwaps(i.input)

		if !reflect.DeepEqual(i.swaps, s) {
			t.Errorf("x: %d: swaps -  Expected: %v, Actual: %v\n", x, i.swaps, s)
		}

		if !reflect.DeepEqual(i.first, f) {
			t.Errorf("x: %d: first -  Expected: %v, Actual: %v\n", x, i.first, f)
		}

		if !reflect.DeepEqual(i.last, l) {
			t.Errorf("x: %d: last -  Expected: %v, Actual: %v\n", x, i.last, l)
		}
	}
}
