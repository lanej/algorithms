package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	edges    [][]int
	start    int
	end      int
	expected []int
}{
	{
		edges: [][]int{
			[]int{1, 2, 1},
			[]int{2, 3, 1},
			[]int{3, 7, 1},
			[]int{7, 9, 1},
			[]int{2, 5, 2},
			[]int{5, 6, 3},
			[]int{6, 8, 5},
			[]int{8, 9, 4},
		},
		start:    1,
		end:      9,
		expected: []int{1, 2, 3, 7, 9},
	},
}

func TestDjikstra(t *testing.T) {
	for ti, tc := range cases {
		actual := djikstra(tc.start, tc.end, tc.edges)

		if !reflect.DeepEqual(tc.expected, actual) {
			t.Errorf("%d: %v != %v\n", ti, tc.expected, actual)
		}
	}
}
