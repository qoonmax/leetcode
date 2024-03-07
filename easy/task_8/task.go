package main

import (
	"fmt"
)

//https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/

func main() {
	fmt.Println(strStr("safbutsad", "sad"))
}

func strStr(haystack string, needle string) int {
	for i, _ := range haystack[:len(haystack)-len(needle)+1] {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
