package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
			continue
		}
		current = current.Next
	}

	return head
}

func main() {
	fmt.Println(deleteDuplicates(&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}}))
}
