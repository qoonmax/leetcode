package main

import (
	"fmt"
	testify "github.com/stretchr/testify/assert"
	"testing"
)

func TestSortColors(t *testing.T) {
	assert := testify.New(t)

	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{2, 0, 2, 1, 1, 1, 0}, expected: []int{0, 0, 1, 1, 1, 2, 2}},
		{input: []int{2, 0, 1}, expected: []int{0, 1, 2}},
		{input: []int{0}, expected: []int{0}},
		{input: []int{1, 2, 0}, expected: []int{0, 1, 2}},
		{input: []int{1, 0, 1}, expected: []int{0, 1, 1}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test case: \"%v\"", testCase.input), func(t *testing.T) {
			result := sortColors(testCase.input)
			assert.Equal(testCase.expected, result, "they should be equal")
		})
	}
}
