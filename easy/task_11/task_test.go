package main

import (
	"fmt"
	"math"
	"testing"
)

func TestStrStr(t *testing.T) {
	testCases := []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 2, 4}},
		{input: []int{1, 2, 9}, expected: []int{1, 3, 0}},
		{input: []int{9}, expected: []int{1, 0}},
		{input: []int{9, 9}, expected: []int{1, 0, 0}},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Testing with input: %v", testCase.input), func(t *testing.T) {
			result := plusOne(testCase.input)
			m := int(math.Max(float64(len(testCase.input)), float64(len(testCase.expected))))
			for i := 0; i <= m-1; i++ {
				if testCase.expected[i] != result[i] {
					t.Errorf("For input %v, expected %v but got %v", testCase.input, testCase.expected, result)
				}
			}
		})
	}
}
