package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	expenditure []int32
	d           int32
	output      int32
}{
	{
		expenditure: []int32{2, 3, 4, 2, 3, 6, 8, 4, 5},
		d:           5,
		output:      2,
	},
	{
		expenditure: []int32{1, 2, 3, 4, 4},
		d:           4,
		output:      0,
	},
}

func TestFreqQueries(t *testing.T) {
	for x, i := range cases {
		fmt.Printf("input: %v, k: %v\n", i.expenditure, i.d)
		output := activityNotifications(i.expenditure, i.d)

		if !reflect.DeepEqual(i.output, output) {
			t.Errorf("x: %d: last -  Expected: %v, Actual: %v\n", x, i.output, output)
		}
	}
}
