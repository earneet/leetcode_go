package Problem_101_200

import . "LeetCode_Go/DataStructure"

func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	level := make([]*TreeNode, 0)
	if root != nil {
		level = append(level, root)
	}

	for len(level) > 0 {
		levelAns := make([]int, 0)
		steps := len(level)
		for i := 0; i < steps; i++ {
			node := level[0]
			level = level[1:]
			levelAns = append(levelAns, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}

		ans = append(ans, levelAns)
	}
	return ans
}
