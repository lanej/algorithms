package main

import "fmt"

func quickselect(nums []int, k int) int {
	return nQuickselect(nums, 0, len(nums)-1, k)
}

func nQuickselect(nums []int, lo, hi, k int) int {
	if lo == hi {
		return nums[lo]
	}
	pivotIndex := (lo + hi) / 2
	fmt.Printf("pi=%d, pv=%d, nums=%v, lo=%d, hi=%d, k=%d\n", pivotIndex, nums[pivotIndex], nums, lo, hi, k)
	parition(nums, lo, hi, pivotIndex)
	fmt.Printf("part -- nums=%v, lo=%d, hi=%d, k=%d\n", nums, lo, hi, k)

	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return nQuickselect(nums, lo, pivotIndex-1, k)
	} else {
		return nQuickselect(nums, pivotIndex+1, hi, k)
	}
}

func parition(nums []int, lo, hi, pi int) int {
	pv := nums[pi]
	swap(nums, pi, hi)

	j := lo
	for i := lo; i < hi; i++ {
		if nums[i] < pv {
			swap(nums, j, i)
			j++
		}

	}

	swap(nums, hi, j)

	return j
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
