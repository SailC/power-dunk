package main

func preorder(root *Node) []int {
	var results []int
	var preOrderFn func(*Node)
	preOrderFn = func(node *Node) {
		if node == nil {
			return
		}
		results = append(results, node.Val)
		for _, child := range node.Children {
			preOrderFn(child)
		}
	}
	preOrderFn(root)
	return results
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var results []int
	var node *Node
	stack := []*Node{root}
	for len(stack) > 0 {
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		results = append(results, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			if node.Children[i] != nil {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return results
}
