package main

import "fmt"

//https://leetcode.com/problems/binary-root-preorder-traversal/

func main() {
	//fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{1, 1, 6}, 3))
	//fmt.Println(merge([]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3))
	//fmt.Println(merge([]int{0, 0, 3, 0}, 3, []int{-1}, 1))
	fmt.Println(merge([]int{0, 0, 3, 0, 0, 0, 0, 0, 0}, 3, []int{-1, 1, 1, 1, 2, 3}, 6))
}

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var (
		a = m - 1
		b = len(nums1) - 1
		c = n - 1
	)

	if a < 0 {
		a = 0
	}

	if b < 0 {
		b = 0
	}

	for c >= 0 {
		fmt.Println(nums1)
		fmt.Println(a, b, c)
		if nums2[c] >= nums1[a] || b == c {
			nums1[b] = nums2[c]
			c--
			b--
		} else {
			nums1[b] = nums1[a]
			nums1[a] = 0
			if a >= 1 {
				a--
			}
			b--
		}
	}

	return nums1
}
