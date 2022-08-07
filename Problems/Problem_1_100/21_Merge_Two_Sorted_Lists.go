package Problem_1_100

import . "LeetCode_Go/DataStructure"

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	fakeHead := ListNode{}
	cur := &fakeHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
			cur = cur.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
			cur = cur.Next
		}
	}

	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return fakeHead.Next
}
