package Problem_501_600

import . "LeetCode_Go/DataStructure"

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{root.Val}

	for i := 0; i < len(root.Children); i++ {
		ans = append(ans, preorder(root.Children[i])...)
	}

	return ans
}
