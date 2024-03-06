package main

import (
	"fmt"
)

//https://leetcode.com/problems/add-two-numbers/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 6}}

	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var resultNode = &ListNode{Val: 0}
	var overDigit int

	currentNode := resultNode
	for l1 != nil || l2 != nil || overDigit != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += overDigit
		overDigit = 0

		if sum > 9 {
			overDigit = 1
			sum %= 10
		}

		currentNode.Next = &ListNode{Val: sum}
		currentNode = currentNode.Next
	}

	return resultNode.Next
}
