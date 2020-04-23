package main

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
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
		} else if nums[l] > nums[mid] {
			if target >= nums[mid] && target <= nums[r] {
				l = mid
			} else {
				r = mid - 1
			}
		} else {
			l++
		}
	}
	if target == nums[l] || target == nums[r] {
		return true
	}
	return false
}
