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
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}

	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var overDigit int
	var currList = l2
	for l1 != nil || overDigit != 0 {
		currList.Val += l1.Val + overDigit
		overDigit = 0

		if currList.Val > 9 {
			overDigit = 1
			currList.Val %= 10
		}

		if l1.Next == nil && currList.Next != nil || l1.Next == nil && overDigit != 0 {
			l1.Next = &ListNode{Val: 0}
		}
		l1 = l1.Next
		if currList.Next == nil && l1 != nil {
			currList.Next = &ListNode{Val: 0}
		}
		currList = currList.Next
	}

	return l2
}
