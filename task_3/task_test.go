package main

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{input: []string{"flower", "flow", "flight"}, expected: "fl"},
		{input: []string{"dog", "racecar", "car"}, expected: ""},
		{input: []string{"dog", "", "car"}, expected: ""},
		{input: []string{""}, expected: ""},
		{input: []string{"a"}, expected: "a"},
		{input: []string{"a", "ab"}, expected: "a"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %s", testCase.input), func(t *testing.T) {
			result := longestCommonPrefix(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %s, expected %s but got %s", testCase.input, testCase.expected, result)
			}
		})
	}
}
