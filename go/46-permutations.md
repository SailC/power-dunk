- 全排列
- https://leetcode-cn.com/problems/permutations/

给定一个 没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
```
[
 [1,2,3],
 [1,3,2],
 [2,1,3],
 [2,3,1],
 [3,1,2],
 [3,2,1]
]
```

思路：
- 既然这个序列没有重复的数组，那么 n! 个结果每个都不相同， 我们可以放心地用dfs生产所有可能的排列。
- dfs生成临时解：
  - 从序列的第一个元素开始扫描，并核对visited数组(visited[i] = true 表示序列的第i个元素已经加入了临时解), 如果visited[i] == true, 那么说明该元素已经被加入了临时解，可以跳过。
  - 如果visited[i] == false, 则将在当前层中将扫描到的不在临时解的元素加入临时解，并进入下一层dfs
  - 临界条件是当临时解的大小等于序列大小的时候，将临时解加入到结果中。
   - 此处注意要复制一下临时解，避免将同一个数组重复的加入结果中。
- T(n) = n ! + (n - 1)! +… + 1  = n!

```go
func permute(nums []int) [][]int
   var results [][]int
   var result []int
   visited := make([]bool, len(nums))
   var dfs func(result []int, visited []bool)
   dfs = func(result []int, visited []bool) {
      if len(result) == len(nums) {
         result = append(result[:0:0], result...)
         results = append(results, result)
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
```
