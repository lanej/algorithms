package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	capacity int
	input    [][]int
	ops      []string
	output   []int
}{
	{
		capacity: 2,
		input:    [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
		ops:      []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
		output:   []int{0, 0, 1, 0, -1, 0, -1, 3, 4},
	},
	{
		capacity: 10,
		ops:      []string{"put", "put", "put", "put", "put", "get", "put", "get", "get", "put", "get", "put", "put", "put", "get", "put", "get", "get", "get", "get", "put", "put", "get", "get", "get", "put", "put", "get", "put", "get", "put", "get", "get", "get", "put", "put", "put", "get", "put", "get", "get", "put", "put", "get", "put", "put", "put", "put", "get", "put", "put", "get", "put", "put", "get", "put", "put", "put", "put", "put", "get", "put", "put", "get", "put", "get", "get", "get", "put", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "get", "get", "get", "put", "put", "put", "get", "put", "put", "put", "get", "put", "put", "put", "get", "get", "get", "put", "put", "put", "put", "get", "put", "put", "put", "put", "put", "put", "put"},
		input:    [][]int{{10, 13}, {3, 17}, {6, 11}, {10, 5}, {9, 10}, {13}, {2, 19}, {2}, {3}, {5, 25}, {8}, {9, 22}, {5, 5}, {1, 30}, {11}, {9, 12}, {7}, {5}, {8}, {9}, {4, 30}, {9, 3}, {9}, {10}, {10}, {6, 14}, {3, 1}, {3}, {10, 11}, {8}, {2, 14}, {1}, {5}, {4}, {11, 4}, {12, 24}, {5, 18}, {13}, {7, 23}, {8}, {12}, {3, 27}, {2, 12}, {5}, {2, 9}, {13, 4}, {8, 18}, {1, 7}, {6}, {9, 29}, {8, 21}, {5}, {6, 30}, {1, 12}, {10}, {4, 15}, {7, 22}, {11, 26}, {8, 17}, {9, 29}, {5}, {3, 4}, {11, 30}, {12}, {4, 29}, {3}, {9}, {6}, {3, 4}, {1}, {10}, {3, 29}, {10, 28}, {1, 20}, {11, 13}, {3}, {3, 12}, {3, 8}, {10, 9}, {3, 26}, {8}, {7}, {5}, {13, 17}, {2, 27}, {11, 15}, {12}, {9, 19}, {2, 15}, {3, 16}, {1}, {12, 17}, {9, 1}, {6, 19}, {4}, {5}, {5}, {8, 1}, {11, 7}, {5, 2}, {9, 28}, {1}, {2, 2}, {7, 4}, {4, 22}, {7, 24}, {9, 26}, {13, 28}, {11, 26}},
		output:   []int{0, 0, 0, 0, 0, -1, 0, 19, 17, 0, -1, 0, 0, 0, -1, 0, -1, 5, -1, 12, 0, 0, 3, 5, 5, 0, 0, 1, 0, -1, 0, 30, 5, 30, 0, 0, 0, -1, 0, -1, 24, 0, 0, 18, 0, 0, 0, 0, -1, 0, 0, 18, 0, 0, -1, 0, 0, 0, 0, 0, 18, 0, 0, -1, 0, 4, 29, 30, 0, 12, -1, 0, 0, 0, 0, 29, 0, 0, 0, 0, 17, 22, 18, 0, 0, 0, -1, 0, 0, 0, 20, 0, 0, 0, -1, 18, 18, 0, 0, 0, 0, 20, 0, 0, 0, 0, 0, 0, 0},
	},
}

func TestLRUCache(t *testing.T) {
	for x, tc := range cases {
		cache := Constructor(tc.capacity)
		output := make([]int, len(tc.input))

		for j, op := range tc.ops {
			input := tc.input[j]
			switch op {
			case "get":
				output[j] = cache.Get(input[0])
			case "put":
				cache.Put(input[0], input[1])
			}

		}

		if !reflect.DeepEqual(tc.output, output) {
			t.Errorf("x: %d\nExpect: %v\nActual: %v\n", x, tc.output, output)
		}
	}
}
