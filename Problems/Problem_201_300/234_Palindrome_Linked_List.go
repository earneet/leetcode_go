package Problem_201_300

import . "LeetCode_Go/DataStructure"

func isPalindrome(head *ListNode) bool {
	fast := head
	slow := head
	stack := make([]int, 0)
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		fast = fast.Next.Next
		slow = slow.Next
	}
	if slow == head {
		return true
	}
	for slow != nil {
		if slow.Val == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		slow = slow.Next
	}
	return len(stack) == 0
}
