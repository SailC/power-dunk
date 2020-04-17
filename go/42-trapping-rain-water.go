package main

type Unit struct {
	leftBound  int
	rightBound int
}

func trap(height []int) int {
	units := make([]Unit, len(height))
	maxHeight := 0
	for i, h := range height {
		units[i].leftBound = maxHeight
		if h > maxHeight {
			maxHeight = h
		}
	}
	maxHeight = 0
	for i := len(height) - 1; i >= 0; i-- {
		units[i].rightBound = maxHeight
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}

	waterSum := 0
	for i, unit := range units {
		upperBound := unit.leftBound
		if unit.rightBound < unit.leftBound {
			upperBound = unit.rightBound
		}
		if upperBound > height[i] {
			waterSum += upperBound - height[i]
		}
	}
	return waterSum
}
