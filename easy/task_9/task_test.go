package main

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{1, 3, 5, 6}, target: 5, expected: 2},
		{nums: []int{1, 3, 5, 6}, target: 2, expected: 1},
		{nums: []int{1, 3, 5, 6}, target: 7, expected: 4},
		{nums: []int{1, 3, 5, 6}, target: 0, expected: 0},
		{nums: []int{1, 3}, target: 0, expected: 0},
		{nums: []int{1}, target: 2, expected: 1},
		{nums: []int{1}, target: 0, expected: 0},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %v, %d", testCase.nums, testCase.target), func(t *testing.T) {
			result := searchInsert(testCase.nums, testCase.target)
			if result != testCase.expected {
				t.Errorf("For input %v, %d, expected %d but got %d", testCase.nums, testCase.target, testCase.expected, result)
			}
		})
	}
}
