package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	testCases := []struct {
		root     *TreeNode
		expected []int
	}{
		{
			root: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			expected: []int{1, 2, 3},
		},
		{
			root: &TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with root: %v", testCase.root), func(t *testing.T) {
			result := preorderTraversal(testCase.root)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("For root %v, expected %v but got %v", testCase.root, testCase.expected, result)
			}
		})
	}
}
