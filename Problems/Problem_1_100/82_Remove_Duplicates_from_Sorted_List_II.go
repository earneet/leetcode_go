package Problem_1_100

import . "LeetCode_Go/DataStructure"

func deleteDuplicates(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	pre := fakeHead
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
		} else if pre.Next != cur {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return fakeHead.Next
}
