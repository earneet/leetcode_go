package Problem_101_200

import . "LeetCode_Go/DataStructure"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func treeDeepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(treeDeepth(node.Left), treeDeepth(node.Right))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := treeDeepth(root.Left) - treeDeepth(root.Right)
	if diff < -1 || diff > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}
