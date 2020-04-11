package main

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	if lo > hi {
		return -1
	}
	for lo < hi {
		mid := (lo + hi) / 2
		if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if target <= nums[lo] {
		return lo
	}
	return lo + 1
}
