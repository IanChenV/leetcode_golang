package leetcode

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	dump := &ListNode{Val: 0}
	head := dump
	carry := 0
	n1, n2 := 0
	current := head
	for l1 != nil || l2 != nil {
		n1 = 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		n2 = 0
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
		current = current.Next
		carry = (n1 + n2 + carry) / 10
	}
	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}
	return dump.Next
}

func Test_addTwoNumbers(t *testing.T) {

}
