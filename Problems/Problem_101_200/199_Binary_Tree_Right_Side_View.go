package Problem_101_200

import . "LeetCode_Go/DataStructure"

func rightDfs(node *TreeNode, depth int, maxDepth *int) []int {
	ans := make([]int, 0)
	if node == nil {
		return ans
	}
	if depth >= *maxDepth {
		ans = append(ans, node.Val)
		*maxDepth += 1
	}
	ans = append(ans, rightDfs(node.Right, depth+1, maxDepth)...)
	ans = append(ans, rightDfs(node.Left, depth+1, maxDepth)...)
	return ans
}

func rightSideView(root *TreeNode) []int {
	maxDepth := 0
	return rightDfs(root, 0, &maxDepth)
}
