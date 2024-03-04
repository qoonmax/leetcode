package main

import (
	"fmt"
)

//https://leetcode.com/problems/remove-element/description/

func main() {
	//fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
	//fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	var counter int
	for i, k := len(nums)-1, len(nums)-1; i >= 0; i-- {
		if nums[i] == val {
			nums[i], nums[k] = nums[k], nums[i]
			k--
			counter++
		}
	}

	fmt.Println(nums)
	return len(nums) - counter
}
