package main

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "III", expected: 3},
		{input: "LVIII", expected: 58},
		{input: "MCMXCIV", expected: 1994},
		{input: "I", expected: 1},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %s", testCase.input), func(t *testing.T) {
			result := romanToInt(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %s, expected %d but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
