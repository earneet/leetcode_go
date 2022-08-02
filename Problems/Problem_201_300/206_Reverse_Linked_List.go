package Problem_201_300

import . "LeetCode_Go/DataStructure"

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode = nil
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
