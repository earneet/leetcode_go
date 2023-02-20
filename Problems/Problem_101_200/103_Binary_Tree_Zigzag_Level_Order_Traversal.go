package Problem_101_200

import . "LeetCode_Go/DataStructure"

func revert(array []int) {
	size := len(array)
	for i := 0; i < size/2; i++ {
		array[i], array[size-i-1] = array[size-i-1], array[i]
	}
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	need_revert := false
	for len(stack) > 0 {
		size := len(stack)
		level := make([]int, 0)
		for i := 0; i < size; i++ {
			node := stack[0]
			stack = stack[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		if need_revert {
			revert(level)
		}
		res = append(res, level)
		need_revert = !need_revert
	}
	return res
}
