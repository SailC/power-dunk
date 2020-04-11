package main

func searchRange(nums []int, target int) []int {
	lo, hi := 0, len(nums)-1
	if lo > hi {
		return []int{-1, -1}
	}
	return []int{
		searchBound(nums, target, true),
		searchBound(nums, target, false),
	}
}

func searchBound(nums []int, target int, isLeftBound bool) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target > nums[mid] {
			lo = mid + 1
		} else if target < nums[mid] {
			hi = mid - 1
		} else if isLeftBound {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	// lo == hi + 1
	if isLeftBound {
		if lo == len(nums) || target != nums[lo] {
			return -1
		}
		return lo
	} else {
		if hi < 0 || target != nums[hi] {
			return -1
		}
		return hi
	}
}
