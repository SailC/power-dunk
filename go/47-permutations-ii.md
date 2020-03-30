- 全排列II
- https://leetcode-cn.com/problems/permutations-ii/

给定一个可包含重复数字的序列，返回所有不重复的全排列。
示例:
输入: [1,1,2]
输出:
```
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```

思路：
- 和全排列I基本一样的dfs
- 但是需要判定重复
- 判定重复的关键是对序列进行排序
- 在生成临时解的过程中，进行剪枝。
- 以[1a, 1b, 2] 这个序列作为例子
  - 在dfs的第一层生成临时解 [1a] , [1b], [2] 的时候，可以对[1b]这个临时解进行剪枝。因为以[1b]为第一个元素的所有临时解都可以由 [1a]生成。
  - 如何剪枝？当准备将一个元素放入临时解的时候，
  - 看看这个元素(nums[i])是否和之前那个元素(nums[i-1])的值一样
  - 以及之前那个元素(nums[i-1])是否不存在于临时解中。nums[i-1]不存在于临时解中说明该元素是作为这一层的临时解的当前元素。我们不允许两个相同的元素出现在临时解的相同位置。
  - 如果nums[i - 1]存在于临时解中，说明nums[i - 1]是在上一层中被加入到临时解中，我们允许两个相同的元素被放入一个临时解的不同位置。

```go
func permuteUnique(nums []int) [][]int
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
```
