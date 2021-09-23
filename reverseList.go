package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func PrintList(pHead *ListNode) {
	cur := pHead
	for cur != nil {
		fmt.Println(*cur)
		cur = cur.Next
	}
}
func main() {
	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 2
	head.Next = node1
	var node2 = new(ListNode)
	node2.Val = 3
	node1.Next = node2
	PrintList(head)
	head2 := ReverseList(head)
	fmt.Println("print reverse List....")
	PrintList(head2)
}
