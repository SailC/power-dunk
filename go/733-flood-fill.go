package main

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	rNum := len(image)
	if rNum == 0 {
		return image
	}
	cNum := len(image[0])
	ogColor := image[sr][sc]
	if newColor == ogColor {
		return image
	}
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i < 0 || i >= rNum || j < 0 || j >= cNum {
			return
		}
		if image[i][j] != ogColor {
			return
		}
		image[i][j] = newColor
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}
	dfs(sr, sc)
	return image
}
