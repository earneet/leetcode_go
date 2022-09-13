package Problem_501_600

import . "LeetCode_Go/DataStructure"

func max(a ...int) int {
	_m := a[0]
	for _, v := range a {
		if _m < v {
			_m = v
		}
	}
	return _m
}

func deepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(deepth(node.Left), deepth(node.Right))
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	deepLeft := deepth(root.Left)
	deepRight := deepth(root.Right)
	return max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right), deepLeft+deepRight)
}
