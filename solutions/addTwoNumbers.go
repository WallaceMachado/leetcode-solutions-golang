package solutions

/*

2. Add Two Numbers
Medium

21552

4225

Add to List

Share
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		n1   = l1
		n2   = l2
		diff int
		head = new(ListNode)
		curr = head
	)

	for n1 != nil || n2 != nil {
		var (
			v1 int
			v2 int
		)

		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		}

		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		}
		sum := diff + v1 + v2
		diff = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}
	if diff > 0 {
		curr.Next = &ListNode{Val: diff}
	}

	return head.Next

}
