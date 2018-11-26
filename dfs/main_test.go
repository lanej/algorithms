package main

import (
	"reflect"
	"testing"
)

var postOrderingDFSCases = []struct {
	edges  [][2]int
	output []int
}{
	{
		edges: [][2]int{
			[2]int{0, 1},
			[2]int{1, 2},
			[2]int{2, 3},
			[2]int{2, 4},
			[2]int{3, 4},
		},
		output: []int{0, 1, 2, 3, 4},
	},
}

func TestPostOrderingDFS(t *testing.T) {
	for i, tc := range postOrderingDFSCases {
		ans := postOrderingDFS(tc.edges)
		if !reflect.DeepEqual(ans, tc.output) {
			t.Errorf("#%d : %v != %v\n", i, ans, tc.output)
		}
	}
}
