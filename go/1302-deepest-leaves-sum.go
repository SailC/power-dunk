package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leavesSum := 0
	var maxLevel int
	var preOrderTraversal func(*TreeNode, int)
	preOrderTraversal = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
			leavesSum = node.Val
		} else if level == maxLevel {
			leavesSum += node.Val
		}
		preOrderTraversal(node.Left, level+1)
		preOrderTraversal(node.Right, level+1)
	}
	preOrderTraversal(root, 0)
	return leavesSum
}
