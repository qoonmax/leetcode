package main

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{input: []string{"leetcode", "leeto"}, expected: -1},
		{input: []string{"sadbutsad", "sad"}, expected: 0},
		{input: []string{"abb", "abaaa"}, expected: -1},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %v", testCase.input), func(t *testing.T) {
			result := strStr(testCase.input[0], testCase.input[1])
			if result != testCase.expected {
				t.Errorf("For input %v, expected %d but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
