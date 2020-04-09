package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var results []int
	var inorderFn func(*TreeNode)
	inorderFn = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorderFn(node.Left)
		results = append(results, node.Val)
		inorderFn(node.Right)
	}
	inorderFn(root)
	return results
}

func inorderTraversal(root *TreeNode) []int {
	var results []int

	var stack []*TreeNode
	for node := root; node != nil; node = node.Left {
		stack = append(stack, node)
	}

	for len(stack) > 0 {
		node, stack := stack[len(stack)-1], stack[:len(stack)-1]
		results = append(results, node.Val)
		for node = node.Right; node != nil; node = node.Left {
			stack = append(stack, node)
		}
	}

	return results
}
