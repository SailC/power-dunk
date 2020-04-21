package main

import (
	"math/bits"
	"sort"
	"unsafe"
)

func majorityElement(nums []int) int {
	cntMap := map[int]int{}
	for _, num := range nums {
		cntMap[num]++
	}
	for num, cnt := range cntMap {
		if cnt > len(nums)/2 {
			return num
		}
	}
	return -1
}

func majorityElement(nums []int) int {
	var majority int
	for i := 0; i < bits.UintSize; i++ {
		mask := 1 << i
		cnt := 0
		for _, num := range nums {
			if num&mask > 0 {
				cnt++
			}
		}
		if cnt > len(nums)/2 {
			majority |= mask
		}
	}
	return majority
}

func majorityElement(nums []int) int {
	var majority int
	var cnt int
	for _, num := range nums {
		if cnt == 0 {
			majority = num
		}
		if num == majority {
			cnt++
		} else {
			cnt--
		}
	}
	return majority
}

func majorityElement(nums []int) int {
	return major(nums, 0, len(nums)-1)
}

func major(nums []int, l int, r int) int {
	if l == r {
		return nums[l]
	}
	m := (l + r) / 2
	leftMajor := major(nums, l, m)
	rightMajor := major(nums, m+1, r)
	if leftMajor == rightMajor {
		return leftMajor
	}
	leftMajorCnt := 0
	for i := l; i <= m; i++ {
		if leftMajor == nums[i] {
			leftMajorCnt++
		}
	}
	rightMajorCnt := 0
	for i := m + 1; i <= r; i++ {
		if rightMajor == nums[i] {
			rightMajorCnt++
		}
	}
	if leftMajorCnt > rightMajorCnt {
		return leftMajor
	}
	return rightMajor
}
