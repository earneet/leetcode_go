package Problem_101_200

import . "LeetCode_Go/DataStructure"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var left, right = maxDepth(root.Left), maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}
