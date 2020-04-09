package main

func postorderTraversal(root *TreeNode) []int {
	var results []int
	var postOrderFn func(*TreeNode)
	postOrderFn = func(root *TreeNode) {
		if root == nil {
			return
		}
		postOrderFn(root.Left)
		postOrderFn(root.Right)
		results = append(results, root.Val)
	}
	postOrderFn(root)
	return results
}

func postorderTraversal(root *TreeNode) []int {
	var results []int
	var stack []*TreeNode
	var node *TreeNode
	for node = root; node != nil; node = node.Left {
		stack = append(stack, node)
	}

	var lastPopNode *TreeNode
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		if node.Right == nil || node.Right == lastPopNode {
			stack = stack[:len(stack)-1]
			lastPopNode = node
			results = append(results, node.Val)
		} else {
			for node = node.Right; node != nil; node = node.Left {
				stack = append(stack, node)
			}
		}
	}

	return results
}
