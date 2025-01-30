package main

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	binStr := ""

	for head != nil {
		binStr += strconv.Itoa(head.Val)
		head = head.Next
	}

	intValue, _ := strconv.ParseInt(binStr, 2, 64)
	return int(intValue)
}

func main() {
	println(getDecimalValue(&ListNode{1, &ListNode{0, &ListNode{1, nil}}}))
}
