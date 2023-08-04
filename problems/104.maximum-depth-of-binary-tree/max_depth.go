package leetcode0104

import (
	"github.com/lushenle/interview-150/structures"
)

type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
