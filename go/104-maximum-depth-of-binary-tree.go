package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var maxLevel int
	var preOrderTraversal func(*TreeNode, int)
	preOrderTraversal = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
		}
		preOrderTraversal(node.Left, level+1)
		preOrderTraversal(node.Right, level+1)
	}
	preOrderTraversal(root, 1)
	return maxLevel
}
