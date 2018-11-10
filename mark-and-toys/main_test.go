package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	prices []int32
	k      int32
	output int32
}{
	{
		prices: []int32{1, 12, 5, 111, 200, 1000, 10},
		k:      50,
		output: 4,
	},
}

func TestFreqQueries(t *testing.T) {
	for x, i := range cases {
		fmt.Printf("input: %v, k: %v\n", i.prices, i.k)
		output := maximumToys(i.prices, i.k)

		if !reflect.DeepEqual(i.output, output) {
			t.Errorf("x: %d: last -  Expected: %v, Actual: %v\n", x, i.output, output)
		}
	}
}
