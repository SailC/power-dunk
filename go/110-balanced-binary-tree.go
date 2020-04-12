package main

func isBalanced(root *TreeNode) bool {
	isBst, _ := isBstWithHeight(root)
	return isBst
}

func isBstWithHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	isLeftBst, leftHeight := isBstWithHeight(root.Left)
	if !isLeftBst {
		return false, 0
	}
	isRightBst, rightHeight := isBstWithHeight(root.Right)
	if !isRightBst {
		return false, 0
	}
	if leftHeight < rightHeight {
		if rightHeight-leftHeight > 1 {
			return false, 0
		}
		return true, rightHeight + 1
	} else {
		if leftHeight-rightHeight > 1 {
			return false, 0
		}
		return true, leftHeight + 1
	}
}
