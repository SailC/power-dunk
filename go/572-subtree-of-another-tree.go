package main

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return t == nil
	}
	if isSameTree(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	}
	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
