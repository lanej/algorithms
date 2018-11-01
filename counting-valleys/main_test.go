package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	n      int32
	input  string
	output int32
}{
	{
		n:      8,
		input:  "UDDDUDUU",
		output: 1,
	},
	{
		n:      12,
		input:  "DDUUDDUDUUUD",
		output: 2,
	},
}

func TestCountValleys(t *testing.T) {
	for x, i := range cases {
		ans := countingValleys(i.n, i.input)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
