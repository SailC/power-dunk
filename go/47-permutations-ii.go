// https://leetcode-cn.com/problems/permutations-ii/

func permuteUnique(nums []int) [][]int {
	var results [][]int
	var result []int
	sort.Ints(nums)
	visited := make([]bool, len(nums))
	var dfs func(result []int, visited []bool)
	dfs = func(result []int, visited []bool) {
		if len(result) == len(nums) {
            result = append(result[:0:0], result...)
            results = append(results, result)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] || (i > 0 && nums[i] == nums[i-1] && !visited[i-1]) {
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
