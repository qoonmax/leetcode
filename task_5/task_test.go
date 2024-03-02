package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		input    []*ListNode
		expected *ListNode
	}{
		{input: []*ListNode{
			{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
			{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		}, expected: &ListNode{
			Val: 1, Next: &ListNode{
				Val: 1, Next: &ListNode{
					Val: 2, Next: &ListNode{
						Val: 3, Next: &ListNode{
							Val: 3, Next: &ListNode{
								Val: 4}}}}}}},
		{input: []*ListNode{
			{Val: 0},
			{Val: 0},
		}, expected: &ListNode{Val: 0, Next: &ListNode{Val: 0}}},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with case: %d", i), func(t *testing.T) {
			result := mergeTwoLists(testCase.input[0], testCase.input[1])
			if !isSameList(result, testCase.expected) {
				t.Errorf("For input testCase(%d), expected %s but got %s", i, listToString(testCase.expected), listToString(result))
			}
		})
	}
}

func isSameList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func listToString(l *ListNode) string {
	var result string
	for l != nil {
		result += fmt.Sprintf("%d -> ", l.Val)
		l = l.Next
	}
	result += "nil"
	return result
}
