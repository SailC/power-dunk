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
		if nums[l] == nums[mid] {
			if l == mid {
				if nums[l] < nums[r] {
					return nums[l]
				}
				return nums[r]
			} else {
				l++
			}
		} else if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}
