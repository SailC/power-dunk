package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	childIdx := 0
	candyIdx := 0
	contentNum := 0
	for childIdx < len(g) && candyIdx < len(s) {
		if g[childIdx] <= s[candyIdx] {
			contentNum++
			childIdx++
		}
		candyIdx++
	}
	return contentNum
}
