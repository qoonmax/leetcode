package main

import (
	"fmt"
)

//https://leetcode.com/problems/merge-two-sorted-lists/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	fmt.Println(mergeTwoLists(list1, list2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	currentNode := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}
		currentNode = currentNode.Next
	}

	if list1 == nil {
		currentNode.Next = list2
	} else {
		currentNode.Next = list1
	}

	return result.Next
}
