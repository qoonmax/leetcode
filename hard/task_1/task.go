package main

import (
	"fmt"
)

//https://leetcode.com/problems/trapping-rain-water/description/

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

func trap(height []int) int {
	var maxHeightIndex int
	var waterCounter int

	if len(height) == 0 {
		return 0
	} else {
		maxHeightIndex = 0
	}

	for i, _ := range height {
		if height[i] > height[maxHeightIndex] {
			maxHeightIndex = i
		}
	}

	var maxLeftHeight = height[0]

	for i := 0; i < maxHeightIndex; i++ {
		if maxLeftHeight > height[i] {
			waterCounter += maxLeftHeight - height[i]
		} else if maxLeftHeight < height[i] {
			maxLeftHeight = height[i]
		}
	}

	var maxRightHeight = height[len(height)-1]

	for i := len(height) - 1; i > maxHeightIndex; i-- {
		if maxRightHeight > height[i] {
			waterCounter += maxRightHeight - height[i]
		} else if maxRightHeight < height[i] {
			maxRightHeight = height[i]
		}
	}

	return waterCounter
}
