package main

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "Hello World", expected: 5},
		{input: "   fly me   to   the moon  ", expected: 4},
		{input: "luffy is still joyboy", expected: 6},
		{input: "a", expected: 1},
		{input: "", expected: 0},
		{input: " ", expected: 0},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %s", testCase.input), func(t *testing.T) {
			result := lengthOfLastWord(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %s, expected %d but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
