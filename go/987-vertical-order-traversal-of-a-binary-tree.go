package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Index struct {
	X int
	Y int
}

func verticalTraversal(root *TreeNode) [][]int {
	var results [][]int
	resultMap := make(map[Index][]int)
	var traverse func(*TreeNode, int, int)
	minX, maxX := 0, 0
	minY, maxY := 0, 0
	traverse = func(node *TreeNode, x int, y int) {
		if node == nil {
			return
		}
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
		index := Index{x, y}
		if result, ok := resultMap[index]; !ok {
			resultMap[index] = []int{node.Val}
		} else {
			resultMap[index] = append(result, node.Val)
		}
		traverse(node.Left, x-1, y+1)
		traverse(node.Right, x+1, y+1)
	}
	traverse(root, 0, 0)
	results = make([][]int, maxX-minX+1)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			index := Index{i, j}
			result := resultMap[index]
			sort.Ints(result)
			results[i-minX] = append(results[i-minX], result...)
		}
	}
	return results
}
