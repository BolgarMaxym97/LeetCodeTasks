package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		first.Next = second.Next
		second.Next = first
		prev.Next = second

		prev = first
		head = first.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println(swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
}
