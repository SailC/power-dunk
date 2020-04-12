package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minLevel := -1
	var preOrderTraversal func(*TreeNode, int)
	preOrderTraversal = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			if minLevel == -1 || level < minLevel {
				minLevel = level
			}
		}
		preOrderTraversal(node.Left, level+1)
		preOrderTraversal(node.Right, level+1)
	}
	preOrderTraversal(root, 1)
	return minLevel
}
