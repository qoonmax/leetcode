package main

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, expected: 6},
		{input: []int{4, 2, 0, 3, 2, 5}, expected: 9},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %v", testCase.input), func(t *testing.T) {
			result := trap(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %v, expected %d but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
