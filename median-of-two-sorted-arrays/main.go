package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(A []int, B []int) float64 {
	var i, j, m, n, maxOfLeft, minOfRight int

	m, n = len(A), len(B)

	if m > n {
		A, B, m, n = B, A, n, m
	}

	imin, imax, halfLen := 0, m, (m+n+1)/2

	for imin <= imax {
		i = (imin + imax) / 2
		j = halfLen - i
		fmt.Printf("i=%d,j=%d,imin=%d,imax=%d,A=%v,B=%v\n", i, j, imin, imax, A, B)
		if i < m && B[j-1] > A[i] {
			imin = i + 1
		} else if i > 0 && A[i-1] > B[j] {
			imax = i - 1
		} else {
			if i == 0 {
				maxOfLeft = B[j-1]
			} else if j == 0 {
				maxOfLeft = A[i-1]
			} else {
				maxOfLeft = int(math.Max(float64(A[i-1]), float64(B[j-1])))
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			if i == m {
				minOfRight = B[j]
			} else if j == n {
				minOfRight = A[i]
			} else {
				minOfRight = int(math.Min(float64(A[i]), float64(B[j])))
			}

			return float64(maxOfLeft+minOfRight) / 2.0
		}
	}

	return 0.0
}
