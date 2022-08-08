package Problem_1_100

import . "LeetCode_Go/DataStructure"

func isValidBSTSub(node *TreeNode, g, l int) bool {
	if node == nil {
		return true
	}

	if node.Val > g && node.Val < l {
		return isValidBSTSub(node.Right, node.Val, l) && isValidBSTSub(node.Left, g, node.Val)
	}
	return false
}

func isValidBST(root *TreeNode) bool {
	return isValidBSTSub(root, -0xFFFFFFFF, 0xFFFFFFFF)
}
