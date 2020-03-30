func permute(nums []int) [][]int {
	var results [][]int
	var result []
	visited := make([]bool, len(nums))
	var dfs func(result []int, visited []bool)
	dfs = func(result []int, visited []bool) {
		if len(result) == len(nums) {
			results = append(results, result[:])
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			dfs(append(result, nums[i]), visited)
			visited[i] = false
		}
	}
	dfs(result, visited)
	return results
}
