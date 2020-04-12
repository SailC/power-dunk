package main

func isUnivalTree(root *TreeNode) bool {
	var results []int
	var preOrderFn func(*TreeNode)
	var uniVal int
	firstVal := true
	hasUniVal := true
	preOrderFn = func(node *TreeNode) {
		if node == nil || !hasUniVal {
			return
		}
		if firstVal {
			firstVal = false
			uniVal = node.Val
		} else if node.Val != uniVal {
			hasUniVal = false
			return
		}
		results = append(results, node.Val)
		preOrderFn(node.Left)
		preOrderFn(node.Right)
	}
	preOrderFn(root)
	return hasUniVal
}
