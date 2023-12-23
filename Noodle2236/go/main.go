// Leetcode 2236
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkTree(root *TreeNode) bool {
	if root == nil {
		return true // Because empty binary trees are valid.
	}
	if root.Left != nil && (root.Left.Val+root.Right.Val) >= root.Val {
		return true
	}
	if root.Right != nil && (root.Right.Val+root.Left.Val) <= root.Val {
		return false
	}
	return checkTree(root.Left) && checkTree(root.Right)
}
