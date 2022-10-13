package Problem_201_300

import . "LeetCode_Go/DataStructure"

func deleteNode(node *ListNode) {
	var pre *ListNode = nil
	for node != nil && node.Next != nil {
		node.Val = node.Next.Val
		pre = node
		node = node.Next
	}
	pre.Next = nil
}
