package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var results []int
	var postOrderFn func(*Node)
	postOrderFn = func(node *Node) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			postOrderFn(child)
		}
		results = append(results, node.Val)
	}
	postOrderFn(root)
	return results
}
