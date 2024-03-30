package main

import "fmt"

//https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	var pointer, counter, maxCounter int

	m := make(map[string]struct{})
	for i := pointer; i < len(s); i++ {
		if _, ok := m[string(s[i])]; ok {
			pointer++
			i = pointer
			counter = 1
			m = map[string]struct{}{string(s[i]): struct{}{}}
		} else {
			m[string(s[i])] = struct{}{}
			counter++
			if counter > maxCounter {
				maxCounter = counter
			}
		}
	}

	return maxCounter
}
