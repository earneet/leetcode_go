package Problem_201_300

import . "LeetCode_Go/DataStructure"

func traceNode(node, target *TreeNode, stack *([]*TreeNode)) bool {
	if node == nil {
		return false
	}
	*stack = append(*stack, node)
	if node == target {
		return true
	}

	if traceNode(node.Left, target, stack) || traceNode(node.Right, target, stack) {
		return true
	}
	*stack = (*stack)[:len(*stack)-1]
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := make([]*TreeNode, 0)
	qPath := make([]*TreeNode, 0)
	traceNode(root, p, &pPath)
	traceNode(root, q, &qPath)
	size := min(len(pPath), len(qPath))

	for i := 0; i < size; i++ {
		if pPath[i] != qPath[i] {
			return pPath[i-1]
		}
	}
	return pPath[size-1]
}
