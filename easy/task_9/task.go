package main

import (
	"fmt"
)

//https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
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
			if r == m {
				l = r
			} else {
				l = m + 1
			}
		} else if nums[m] > target {
			if l == m {
				r = l
			} else {
				r = m - 1
			}
		} else {
			x = m
		}

		m = (l + r) / 2

		if l == r {
			if nums[m] > target {
				x = m
			} else if nums[m] < target {
				x = m + 1
			}
		}
	}

	return x
}
