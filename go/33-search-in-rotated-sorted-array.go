package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1

	for l+1 < r {
		mid := (l + r) / 2
		if nums[l] < nums[mid] {
			if target >= nums[l] && target <= nums[mid] {
				r = mid
			} else {
				l = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[r] {
				l = mid
			} else {
				r = mid - 1
			}
		}
	}
	if target == nums[l] {
		return l
	} else if target == nums[r] {
		return r
	}
	return -1
}
