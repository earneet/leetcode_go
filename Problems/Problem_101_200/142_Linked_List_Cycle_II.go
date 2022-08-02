package Problem_101_200

import . "LeetCode_Go/DataStructure"

func detectCycle(head *ListNode) *ListNode {
	quick := head
	slow := head
	findCycle := false
	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			findCycle = true
			break
		}
	}

	if !findCycle {
		return nil
	}

	slow = head
	for quick != slow {
		quick = quick.Next
		slow = slow.Next
	}
	return slow
}
