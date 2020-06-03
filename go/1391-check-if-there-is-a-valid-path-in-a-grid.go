package main

import "fmt"

type direction struct {
	dx int
	dy int
}

var codeMap = map[int][2]direction{
	2: {{-1, 0}, {1, 0}},
	1: {{0, -1}, {0, 1}},
	3: {{0, -1}, {1, 0}},
	4: {{0, 1}, {1, 0}},
	5: {{0, -1}, {-1, 0}},
	6: {{0, 1}, {-1, 0}},
}

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	if m == 0 {
		return false
	}
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var isValidPath func(int, int) bool
	isValidPath = func(i int, j int) bool {
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if i == m-1 && j == n-1 {
			return true
		}
		code := grid[i][j]
		dirs, _ := codeMap[code]
		visited[i][j] = true
		for _, dir := range dirs {
			if !connected(i, j, dir, grid, visited) {
				continue
			}
			if isValidPath(i+dir.dx, j+dir.dy) {
				return true
			}
		}
		visited[i][j] = false
		return false
	}

	return isValidPath(0, 0)
}

func connected(i int, j int, dir direction, grid [][]int, visited [][]bool) bool {
	m, n := len(grid), len(grid[0])
	k, l := i+dir.dx, j+dir.dy
	if k < 0 || k >= m || l < 0 || l >= n {
		return false
	}
	if visited[k][l] {
		return false
	}
	code := grid[k][l]
	dirs := codeMap[code]
	for _, d := range dirs {
		if k+d.dx == i && l+d.dy == j {
			return true
		}
	}
	return false
}

func main() {
	grid := [][]int{{2, 4, 3}, {6, 5, 2}}
	if hasValidPath(grid) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
