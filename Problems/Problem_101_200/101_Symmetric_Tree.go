package Problem_101_200

import . "LeetCode_Go/DataStructure"

func revertTree(node *TreeNode) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	if node.Left != nil {
		revertTree(node.Left)
	}
	if node.Right != nil {
		revertTree(node.Right)
	}
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSymmetric(root *TreeNode) bool {
	revertTree(root.Right)
	return isSameTree(root.Left, root.Right)
}
