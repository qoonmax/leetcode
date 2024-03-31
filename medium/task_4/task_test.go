package main

import (
	"fmt"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert := testify.New(t)

	testCases := []struct {
		input    int
		expected string
	}{
		{input: 3, expected: "III"},
		{input: 58, expected: "LVIII"},
		{input: 1994, expected: "MCMXCIV"},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case: \"%d\"", testCase.input), func(t *testing.T) {
			result := intToRoman(testCase.input)
			assert.Equal(testCase.expected, result, "they should be equal")
		})
	}
}
