package Problem_1_100

import . "LeetCode_Go/DataStructure"

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}

	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
