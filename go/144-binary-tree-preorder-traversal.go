package main

func preorderTraversal(root *TreeNode) []int {
	var results []int
	var preOrderFn func(*TreeNode)
	preOrderFn = func(node *TreeNode) {
		if node == nil {
			return
		}
		results = append(results, node.Val)
		preOrderFn(node.Left)
		preOrderFn(node.Right)
	}
	preOrderFn(root)
	return results
}

func preorderTraversal(root *TreeNode) []int {
	var results []int
	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}
	var node *TreeNode
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		results = append(results, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return results
}
