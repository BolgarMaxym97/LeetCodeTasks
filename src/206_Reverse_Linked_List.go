package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}

	return prev
}

func main() {
	fmt.Println(reverseList(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}))
}
