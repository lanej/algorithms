package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{9, 3, 5, 1, 7},
		output: []int{1, 3, 5, 7, 9},
	},
}

func TestHeapPop(t *testing.T) {
	for i, tc := range cases {
		heap := heapify(tc.input)

		ans := []int{}

		for i := 0; i < len(tc.input); i++ {
			ans = append(ans, heap.Pop())
		}

		if !reflect.DeepEqual(ans, tc.output) {
			t.Errorf("%d: %v != %v", i, tc.output, ans)
		}
	}
}
