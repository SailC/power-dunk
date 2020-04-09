package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var results [][]int
	var que []*Node
	if root != nil {
		que = append(que, root)
	}
	for len(que) > 0 {
		result := make([]int, len(que))
		var nextQue []*Node
		for i, node := range que {
			result[i] = node.Val
			if node.Children != nil {
				nextQue = append(nextQue, node.Children...)
			}
		}
		results = append(results, result)
		que = nextQue
	}
	return results
}
