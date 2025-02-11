package main

import (
	"fmt"
	"math/big"
)

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

func listToStr(node *ListNode) string {
	var result string
	for node != nil {
		result += fmt.Sprint(node.Val)
		node = node.Next
	}

	return result
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode

	l1R, l2R := reverseList(l1), reverseList(l2)
	n1, n2, res := new(big.Int), new(big.Int), new(big.Int)
	firstNum, _ := n1.SetString(listToStr(l1R), 10)
	secondNum, _ := n2.SetString(listToStr(l2R), 10)
	resNum := res.Add(firstNum, secondNum).String()

	for _, ch := range resNum {
		result = &ListNode{int(ch - '0'), result}
	}

	return result
}

func main() {
	list := &ListNode{
		1, &ListNode{
			0, &ListNode{
				0, &ListNode{
					0, &ListNode{
						0, &ListNode{
							0, &ListNode{
								0, &ListNode{
									0, &ListNode{
										0, &ListNode{
											0, &ListNode{
												0, &ListNode{
													0, &ListNode{
														0, &ListNode{
															0, &ListNode{
																0, &ListNode{
																	0, &ListNode{
																		0, &ListNode{
																			0, &ListNode{
																				0, &ListNode{
																					0, &ListNode{
																						0, &ListNode{
																							0, &ListNode{
																								0, &ListNode{
																									0, &ListNode{
																										0, &ListNode{
																											0, &ListNode{
																												0, &ListNode{
																													0, &ListNode{
																														0, &ListNode{
																															1, nil,
																														},
																													},
																												},
																											},
																										},
																									},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	fmt.Println(addTwoNumbers(
		list,
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
	))
}
