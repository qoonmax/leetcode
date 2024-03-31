package main

import "fmt"

//https://leetcode.com/problems/container-with-most-water/description/

func main() {
	fmt.Println(maxArea([]int{2, 3, 4, 5, 18, 17, 6}))
}

func maxArea(height []int) int {
	var left = 0
	var right = len(height) - 1
	var maxWater int

	for left != right {
		m := min(height[left], height[right])

		if maxWater < m*(right-left) {
			maxWater = m * (right - left)
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxWater
}
