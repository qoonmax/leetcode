package main

import (
	"fmt"
)

//https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	//head := &ListNode{Val: 1}

	fmt.Println(deleteDuplicates(head))
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			if curr.Next.Next == nil {
				curr.Next = nil
			} else {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}

	return head
}
