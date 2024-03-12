package main

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{2, 2, 1}, expected: 1},
		{input: []int{4, 1, 2, 1, 2}, expected: 4},
		{input: []int{1}, expected: 1},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %v", testCase.input), func(t *testing.T) {
			result := singleNumber(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %v, expected %d but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
