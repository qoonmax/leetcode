package main

import (
	"fmt"
)

//https://leetcode.com/problems/search-insert-position/

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
}

func searchInsert(nums []int, target int) int {
	var (
		l = 0
		r = len(nums) - 1
		m = (l + r) / 2
		x = -1
	)

	for x < 0 {
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m
		}

		m = (l + r) / 2

		if l == r || l > r {
			if nums[m] > target {
				x = m
			} else if nums[m] < target {
				x = m + 1
			}
		}
	}

	return x
}
