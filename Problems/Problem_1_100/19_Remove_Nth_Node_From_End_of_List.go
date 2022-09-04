package Problem_1_100

import . "LeetCode_Go/DataStructure"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := ListNode{Next: head}
	fastPointer := head
	for i := 0; i < n && fastPointer != nil; i++ {
		fastPointer = fastPointer.Next
	}

	if fastPointer == nil {
		return fakeHead.Next.Next
	}

	fastPointer = fastPointer.Next
	slowPointer := head
	for fastPointer != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next
	}
	slowPointer.Next = slowPointer.Next.Next
	return fakeHead.Next
}
