package main

// ListNode is the defination of the node of a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	l, r := dummy, head
	for r != nil {
		r = r.Next
		n--
		if n == 0 {
			break
		}
	}
	for r != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return dummy.Next
}
