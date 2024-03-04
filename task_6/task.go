package main

import (
	"fmt"
)

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 4}))
}

/* First try
func removeDuplicates(nums []int) int {
	var uniq []int

	uniq = append(uniq, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		} else {
			uniq = append(uniq, nums[i])
		}
	}

	for i, u := range uniq {
		nums[i] = u
	}

	return len(uniq)
}
*/

func removeDuplicates(nums []int) int {
	var counter = 1

	if len(nums) == 0 {
		return 1
	}

	for i, j := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[j] = nums[i]
			j++
			counter++
		}
	}

	return counter
}
