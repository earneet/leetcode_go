package Problem_501_600

import . "LeetCode_Go/DataStructure"

func equals(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil && subRoot != nil || root != nil && subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}

	return equals(root.Left, subRoot.Left) && equals(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	if equals(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
