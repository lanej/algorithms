package main

import (
	"reflect"
	"testing"
)

// http://hr.gs/deabde

// Function Description

// Complete the function arrayManipulation in the editor below. It must return an integer, the maximum value in the resulting array.

// arrayManipulation has the following parameters:

// n - the number of elements in your array
// queries - a two dimensional array of queries where each queries[i] contains three integers, a, b, and k.
// Input Format

// The first line contains two space-separated integers  and , the size of the array and the number of operations.
// Each of the next  lines contains three space-separated integers ,  and , the left index, right index and summand.

// Constraints

// Output Format

// Return the integer maximum value in the finished array.

// Sample Input

// 5 3
// 1 2 100
// 2 5 100
// 3 4 100
// Sample Output

// 200
// Explanation

// After the first update list will be 100 100 0 0 0.
// After the second update list will be 100 200 100 100 100.
// After the third update list will be 100 200 200 200 100.
// The required answer will be .

var cases = []struct {
	zine []string
	note []string
	ans  bool
}{
	{
		zine: []string{"give", "me", "one", "grand", "today", "night"},
		note: []string{"give", "one", "grand", "today"},
		ans:  true,
	},
	{
		zine: []string{"two", "times", "three", "is", "not", "four"},
		note: []string{"two", "times", "two", "is", "four"},
		ans:  false,
	},
}

func TestCheckMagazine(t *testing.T) {
	for x, i := range cases {
		ans := checkMagazine(i.zine, i.note)

		if !reflect.DeepEqual(i.ans, ans) {
			t.Errorf("t: %d | e : %v, a: %v\n", x, i.ans, ans)
		}
	}
}
