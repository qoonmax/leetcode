package main

import (
	"fmt"
)

//https://leetcode.com/problems/plus-one/

func main() {
	fmt.Println(plusOne([]int{9, 9}))
}

func plusOne(digits []int) []int {
	var d = 1

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += d
		d = 0
		if digits[i] > 9 {
			digits[i] = 0
			d = 1
		}
		if i == 0 && d == 1 {
			digits = append([]int{1}, digits...)
		}
	}

	return digits
}
