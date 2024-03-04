package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		nums         []int
		expectedNums []int
		expectedK    int
	}{
		{nums: []int{1, 1, 2}, expectedNums: []int{1, 2}, expectedK: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expectedNums: []int{0, 1, 2, 3, 4}, expectedK: 5},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case %d", i+1), func(t *testing.T) {
			k := removeDuplicates(testCase.nums)
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
