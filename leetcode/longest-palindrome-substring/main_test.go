package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	input  string
	output string
}{
	{
		input:  "aba",
		output: "aba",
	},
	{
		input:  "abcd",
		output: "a",
	},
	{
		input:  "abbd",
		output: "bb",
	},
}

func TestLongestPalindrome(t *testing.T) {
	for x, i := range cases {
		ans := longestPalindrome(i.input)

		if !reflect.DeepEqual(i.output, ans) {
			t.Errorf("x: %d, Expected: %v, Actual: %v\n", x, i.output, ans)
		}
	}
}
