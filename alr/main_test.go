package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	size      int32
	rotations int32
	input     []int32
	output    []int32
}{
	{
		size:      20,
		rotations: 10,
		input:     []int32{41, 73, 89, 7, 10, 1, 59, 58, 84, 77, 77, 97, 58, 1, 86, 58, 26, 10, 86, 51},
		output:    []int32{77, 97, 58, 1, 86, 58, 26, 10, 86, 51, 41, 73, 89, 7, 10, 1, 59, 58, 84, 77},
	},
	{
		size:      5,
		rotations: 4,
		input:     []int32{1, 2, 3, 4, 5},
		output:    []int32{5, 1, 2, 3, 4},
	},
}

func TestRotateLeft(t *testing.T) {
	for _, i := range cases {
		if answer := rotLeft(i.input, i.rotations); !reflect.DeepEqual(i.output, answer) {
			t.Errorf("Expected: %v, Actual: %v\n", i.output, answer)
		}
	}
}
