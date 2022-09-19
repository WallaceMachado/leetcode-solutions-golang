package main

import "fmt"

func main() {

	var (
		head = new(ListNode)
		curr = head
	)
	fmt.Println("head ", &head)
	fmt.Println("curr ", &curr)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		n1   = l1
		n2   = l2
		diff int
		head = new(ListNode)
		curr = head
	)
	fmt.Println("head ", head)
	fmt.Println("curr ", curr)
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
		fmt.Println("diff ", diff)
		curr.Next = &ListNode{Val: sum % 10}
		fmt.Println("curr ", curr)
		curr = curr.Next
		fmt.Println("curr ", curr)
	}
	if diff > 0 {
		curr.Next = &ListNode{Val: diff}
	}

	return head.Next

}
