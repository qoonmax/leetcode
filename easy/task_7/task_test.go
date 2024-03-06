package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		nums         []int
		val          int
		expectedNums []int
		expectedK    int
	}{
		{nums: []int{3, 2, 2, 3}, val: 3, expectedNums: []int{2, 2, 3, 3}, expectedK: 2},
		{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2, expectedNums: []int{0, 1, 0, 4, 3, 2, 2, 2}, expectedK: 5},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			k := removeElement(testCase.nums, testCase.val)
			if k != testCase.expectedK {
				t.Errorf("Expected k = %d, got k = %d", testCase.expectedK, k)
			}
			for j := 0; j < k; j++ {
				if testCase.nums[j] != testCase.expectedNums[j] {
					t.Errorf("Expected nums[%d] = %d, got nums[%d] = %d", j, testCase.expectedNums[j], j, testCase.nums[j])
				}
			}
		})
	}
}
