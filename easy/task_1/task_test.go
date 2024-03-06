package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{12321, true},
		{123321, true},
		{123456, false},
		{100001, true},
		{1221, true},
		{121, true},
		{1234554321, true},
		{-1234554321, false},
		{123450, false},
		{1234567890, false},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %d", testCase.input), func(t *testing.T) {
			result := isPalindrome(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %d, expected %t but got %t", testCase.input, testCase.expected, result)
			}
		})
	}
}
