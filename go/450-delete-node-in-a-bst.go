package main

import "math"

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}
	// key == root.Val && root.Left != nil && root.Right != nil
	minRight := getMin(root.Right)
	root.Val = minRight
	root.Right = deleteNode(root.Right, minRight)
	return root
}

func getMin(root *TreeNode) int {
	minVal := math.MaxInt64
	for root != nil {
		if root.Val < minVal {
			minVal = root.Val
		}
		root = root.Left
	}
	return minVal
}
