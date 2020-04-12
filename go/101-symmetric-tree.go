package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var preOrderTraversal func(*TreeNode, *TreeNode) bool
	preOrderTraversal = func(left *TreeNode, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == nil && right == nil
		}
		return left.Val == right.Val &&
			preOrderTraversal(left.Left, right.Right) &&
			preOrderTraversal(left.Right, right.Left)
	}

	return preOrderTraversal(root.Left, root.Right)
}
