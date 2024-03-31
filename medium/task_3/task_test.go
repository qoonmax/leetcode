package main

import (
	"fmt"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := testify.New(t)

	testCases := []struct {
		input    []int
		expected int
	}{
		{input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, expected: 49},
		{input: []int{1, 1}, expected: 1},
		{input: []int{2, 3, 4, 5, 18, 17, 6}, expected: 17},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case: \"%v\"", testCase.input), func(t *testing.T) {
			result := maxArea(testCase.input)
			assert.Equal(testCase.expected, result, "they should be equal")
		})
	}
}
