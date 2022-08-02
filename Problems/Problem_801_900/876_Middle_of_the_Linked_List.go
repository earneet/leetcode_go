package Problem_801_900

import . "LeetCode_Go/DataStructure"

func middleNode(head *ListNode) *ListNode {
	quick := head
	slow := head

	for quick != nil {
		quick = quick.Next
		if quick == nil {
			return slow
		}
		quick = quick.Next
		slow = slow.Next
	}
	return slow
}
