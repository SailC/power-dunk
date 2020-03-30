func exist(board [][]byte, word string) bool {

	m := len(board)
	n := len(board[0])
	const visited = '.'

	var dfs func(i int, j int, start int) bool
	dfs = func(i int, j int, start int) bool {
		if start == len(word) {
			return true
		}
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if board[i][j] == visited || word[start] != board[i][j] {
			return false
		}
		tmp := board[i][j]
		board[i][j] = visited
		if dfs(i-1, j, start+1) ||
			dfs(i+1, j, start+1) ||
			dfs(i, j-1, start+1) ||
			dfs(i, j+1, start+1) {
			return true
		}
		board[i][j] = tmp
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
