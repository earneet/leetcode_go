package Problem_101_200

import . "LeetCode_Go/DataStructure"

func search(node *TreeNode, target int, stack []int, ans *[][]int) {
	if node == nil {
		return
	}
	stack = append(stack, node.Val)
	if target == node.Val && node.Left == nil && node.Right == nil {
		tmp := make([]int, len(stack))
		copy(tmp, stack)
		*ans = append(*ans, tmp)
	}
	search(node.Left, target-node.Val, stack, ans)
	search(node.Right, target-node.Val, stack, ans)
	stack = stack[:len(stack)-1]
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	search(root, targetSum, []int{}, &ans)
	return ans
}
