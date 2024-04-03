package main

import "fmt"

//https://leetcode.com/problems/sort-colors/description/

func main() {
	fmt.Println(sortColors([]int{2, 0, 2, 1, 1, 1, 0}))

	//2, 0, 2, 1, 1, 1, 0
	//2, 0, 2, 1, 1, 1, 0
	//2, 0, 2, 1, 1, 1, 0
	//2, 0, 2, 1, 1, 1, 0
}

func sortColors(nums []int) []int {
	ar := [3]int{}

	for i := 0; i < len(nums); i++ {
		ar[nums[i]] += 1
	}

	for i, n := 0, 0; i < len(ar); {
		if ar[i] > 0 {
			nums[n] = i
			ar[i]--
			n++
		} else {
			i++
		}
	}

	return nums
}
