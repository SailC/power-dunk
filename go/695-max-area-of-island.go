package main

func maxAreaOfIsland(grid [][]int) int {
	rNum := len(grid)
	if rNum == 0 {
		return 0
	}
	cNum := len(grid[0])
	maxArea := 0
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i < 0 || i >= rNum || j < 0 || j >= cNum {
			return 0
		}
		if grid[i][j] == 0 {
			return 0
		}
		grid[i][j] = 0
		return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
	}
	for i := 0; i < rNum; i++ {
		for j := 0; j < cNum; j++ {
			if grid[i][j] == 0 {
				continue
			}
			area := dfs(i, j)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}
