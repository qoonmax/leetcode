package main

import "fmt"

//https://leetcode.com/problems/single-number/

func main() {
	fmt.Println(singleNumber([]int{1, 1, 2, 3, 2}))
}

func singleNumber(nums []int) int {
	result := 0

	for _, n := range nums {
		result ^= n
	}

	return result
}
