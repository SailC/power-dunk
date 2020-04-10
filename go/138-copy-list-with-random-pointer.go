package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copyMap := make(map[*Node]*Node)
	var dfsCopy func(*Node) *Node
	dfsCopy = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if _, ok := copyMap[node]; !ok {
			copyNode := &Node{
				node.Val,
				nil,
				nil,
			}
			copyMap[node] = copyNode
			copyNode.Next = dfsCopy(node.Next)
			copyNode.Random = dfsCopy(node.Random)
		}
		copyNode := copyMap[node]
		return copyNode
	}
	return dfsCopy(head)
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	que := []*Node{head}
	copyMap := map[*Node]*Node{
		head: &Node{
			head.Val, nil, nil,
		}}
	var node *Node
	for len(que) > 0 {
		node, que = que[0], que[1:]
		copyNode := copyMap[node]
		if node.Next != nil {
			if _, ok := copyMap[node.Next]; !ok {
				copyNext := &Node{
					node.Next.Val,
					nil,
					nil,
				}
				copyMap[node.Next] = copyNext
				que = append(que, node.Next)
			}
			copyNode.Next = copyMap[node.Next]
		}
		if node.Random != nil {
			if _, ok := copyMap[node.Random]; !ok {
				copyRandom := &Node{
					node.Random.Val,
					nil,
					nil,
				}
				copyMap[node.Random] = copyRandom
				que = append(que, node.Random)
			}
			copyNode.Random = copyMap[node.Random]
		}
	}
	return copyMap[head]
}
