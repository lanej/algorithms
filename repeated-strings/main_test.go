package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	input  string
	n      int64
	output int64
}{
	{
		input:  "aba",
		n:      10,
		output: 7,
	},
}

func TestRepeatedString(t *testing.T) {
	for x, i := range cases {
		ans := repeatedString(i.input, i.n)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
