package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// Key takeaway:
// If the linked lists have different lengths, we need to compensate for the difference.
// First, traverse your own list, then switch to the other list.
// This ensures they "sync up" and meet at the intersection point.
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	currA, currB := headA, headB

	for currA != currB {
		if currA == nil {
			currA = headB
		} else {
			currA = currA.Next
		}

		if currB == nil {
			currB = headA
		} else {
			currB = currB.Next
		}
	}

	return currA
}

func main() {
	intersection := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	listA := &ListNode{4, &ListNode{1, intersection}}
	listB := &ListNode{5, &ListNode{0, &ListNode{1, intersection}}}

	fmt.Println(getIntersectionNode(listA, listB))
}
