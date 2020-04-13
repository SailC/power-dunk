package main

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
