package main

import "fmt"

//https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
}

func lengthOfLongestSubstring(s string) int {
	var left, right, maxCount int
	m := make(map[uint8]int)

	for right < len(s) {
		if _, ok := m[s[right]]; ok {
			if left < m[s[right]]+1 {
				left = m[s[right]] + 1
			}
		}

		m[s[right]] = right
		if maxCount < right-left+1 {
			maxCount = right - left + 1
		}

		right++
	}

	return maxCount
}
