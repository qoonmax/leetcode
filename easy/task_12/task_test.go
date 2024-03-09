package main

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	testCases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		},
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 1}},
			expected: &ListNode{Val: 1},
		},
		{
			input:    &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			expected: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %v", listToString(testCase.input)), func(t *testing.T) {
			result := deleteDuplicates(testCase.input)
			if !isSameList(result, testCase.expected) {
				t.Errorf("For input testCase(%v), expected %v but got %v", listToString(testCase.input), listToString(testCase.expected), listToString(result))
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
