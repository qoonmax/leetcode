package main

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{input: "{}", expected: true},
		{input: "(", expected: false},
		{input: "()", expected: true},
		{input: "(])", expected: false},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "()]", expected: false},
		{input: "[()]", expected: true},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %s", testCase.input), func(t *testing.T) {
			result := isValid(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %s, expected %t but got %t", testCase.input, testCase.expected, result)
			}
		})
	}
}
