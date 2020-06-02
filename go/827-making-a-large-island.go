package main

import "fmt"

type IslandGroup struct {
	area  int
	label int
}

func largestIsland(grid [][]int) int {
	rNum := len(grid)
	if rNum == 0 {
		return 0
	}
	cNum := len(grid[0])
	islandGroups := make([][]IslandGroup, rNum)
	for i := 0; i < rNum; i++ {
		islandGroups[i] = make([]IslandGroup, cNum)
	}

	visited := make([][]bool, rNum)
	for i := 0; i < rNum; i++ {
		visited[i] = make([]bool, cNum)
	}

	numGroup := 0
	var dfs func(int, int) int
	dfs = func(x int, y int) int {
		if x < 0 || x >= rNum || y < 0 || y >= cNum {
			return 0
		}
		if visited[x][y] || grid[x][y] == 0 {
			return 0
		}
		visited[x][y] = true
		sum := 1 + dfs(x-1, y) + dfs(x+1, y) + dfs(x, y-1) + dfs(x, y+1)
		islandGroups[x][y] = IslandGroup{sum, numGroup}
		numGroup++
		return sum
	}

	for i := 0; i < rNum; i++ {
		for j := 0; j < cNum; j++ {
			area := dfs(i, j)
			islandGroups[i][j] = IslandGroup{
				area,
				numGroup,
			}
			numGroup++
		}
	}
	maxArea := 0
	for i := 0; i < rNum; i++ {
		for j := 0; j < cNum; j++ {
			var area int
			if grid[i][j] == 0 {
				area = getPotentialArea(i, j, rNum, cNum, islandGroups)
			} else {
				area = islandGroups[i][j].area
			}
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func getPotentialArea(i int, j int, rNum int, cNum int, groups [][]IslandGroup) int {
	areaSum := 1
	labelToArea := make(map[int]int)
	indices := [][]int{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
	for _, index := range indices {
		x, y := index[0], index[1]
		if x < 0 || x >= rNum || y < 0 || y >= cNum {
			continue
		}
		label, area := groups[i][j].label, groups[i][j].area
		labelToArea[label] = area
	}
	for _, area := range labelToArea {
		areaSum += area
	}

	return areaSum
}

func main() {
	grid := [][]int{{1, 0}, {0, 1}}
	area := largestIsland(grid)
	fmt.Println(area)
}
