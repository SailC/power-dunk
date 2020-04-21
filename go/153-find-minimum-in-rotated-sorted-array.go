package main

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] < nums[r] {
			return nums[l]
		}
		mid := (l + r) / 2
		if nums[mid] > nums[l] {
			l = mid + 1
		} else if nums[mid] < nums[l] {
			r = mid
		} else {
			if nums[l] < nums[r] {
				return nums[l]
			} else {
				return nums[r]
			}
		}
	}
	return nums[l]
}
