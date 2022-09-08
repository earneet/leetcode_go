package Problem_1_100

import . "LeetCode_Go/DataStructure"

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	if root.Left != nil {
		ans = append(ans, inorderTraversal(root.Left)...)
	}
	ans = append(ans, root.Val)
	if root.Right != nil {
		ans = append(ans, inorderTraversal(root.Right)...)
	}
	return ans
}
