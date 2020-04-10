package main

type Node struct {
	Val       int
	Neighbors []*Node
}

var node1 = &Node{
	Val:       1,
	Neighbors: []*Node{node2, node4},
}

var node2 = &Node{
	Val:       1,
	Neighbors: []*Node{node1, node3},
}

var node3 = &Node{
	Val:       1,
	Neighbors: []*Node{node2, node4},
}

var node4 = &Node{
	Val:       1,
	Neighbors: []*Node{node1, node3},
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	clonedMap := map[*Node]*Node{}
	var dfs func(*Node) *Node
	dfs = func(v *Node) *Node {
		if v == nil {
			return nil
		}
		if u, ok := clonedMap[v]; ok {
			return u
		}
		clonedNode := &Node{
			v.Val,
			make([]*Node, len(v.Neighbors)),
		}
		clonedMap[v] = clonedNode
		for i, neighbor := range v.Neighbors {
			clonedNode.Neighbors[i] = dfs(neighbor)
		}
		return clonedNode
	}
	return dfs(node)
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	clonedMap := make(map[*Node]*Node)
	que := []*Node{node}
	clonedMap[node] = &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	var v *Node
	for len(que) > 0 {
		v, que = que[0], que[1:]
		u := clonedMap[v]
		for i, neighbor := range v.Neighbors {
			if _, ok := clonedMap[neighbor]; !ok {
				clonedMap[neighbor] = &Node{
					Val:       neighbor.Val,
					Neighbors: make([]*Node, len(neighbor.Neighbors)),
				}
				que = append(que, neighbor)
			}
			u.Neighbors[i] = clonedMap[neighbor]
		}
	}
	return clonedMap[node]
}
