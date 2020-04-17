package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWaterArea := 0
	for left < right {
		h := height[left]
		if height[right] < height[left] {
			h = height[right]
		}
		area := (right - left) * h
		if area > maxWaterArea {
			maxWaterArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxWaterArea
}
