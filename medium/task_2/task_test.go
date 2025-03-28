package main

import (
	"fmt"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := testify.New(t)

	testCases := []struct {
		input    string
		expected int
	}{
		{input: "abcabcbb", expected: 3},
		{input: "bbbb", expected: 1},
		{input: "pwwkew", expected: 3},
		{input: "", expected: 0},
		{input: "d", expected: 1},
		{input: "tmmzuxt", expected: 5},
		{input: "dvdf", expected: 3},
		{input: "aabaab!bb", expected: 3},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case: \"%s\"", testCase.input), func(t *testing.T) {
			result := lengthOfLongestSubstring(testCase.input)
			assert.Equal(testCase.expected, result, "they should be equal")
		})
	}
}
