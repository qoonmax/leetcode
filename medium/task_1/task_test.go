package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}

	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1}}}}}

	result := addTwoNumbers(l1, l2)

	if !isSameList(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func isSameList(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
