package main

func numIslands(grid [][]byte) int {
	cnt := 0
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var dfsFlush func(int, int)
	dfsFlush = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfsFlush(i, j+1)
		dfsFlush(i, j-1)
		dfsFlush(i+1, j)
		dfsFlush(i-1, j)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfsFlush(i, j)
			}
		}
	}

	return cnt
}
