package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			nums1:    []int{0, 0, 3, 0},
			m:        2,
			nums2:    []int{-1},
			n:        1,
			expected: []int{-1, 0, 0, 3},
		},
		{
			nums1:    []int{0, 0, 3, 0, 0, 0, 0, 0, 0},
			m:        3,
			nums2:    []int{-1, 1, 1, 1, 2, 3},
			n:        6,
			expected: []int{-1, 0, 0, 1, 1, 1, 2, 3, 3},
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case: nums1=%v, m=%d, nums2=%v, n=%d", testCase.nums1, testCase.m, testCase.nums2, testCase.n), func(t *testing.T) {
			result := merge(testCase.nums1, testCase.m, testCase.nums2, testCase.n)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("Expected %v, but got %v", testCase.expected, result)
			}
		})
	}
}
