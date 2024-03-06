package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/longest-common-prefix/description/

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	var result string

	sort.Strings(strs)

	for i := 0; i < len(strs[0]); i++ {
		if strs[0] == "" {
			return ""
		}
		result += string(strs[0][i])
		for _, word := range strs {
			if result == word[:i+1] {
				continue
			} else {
				result = result[:i]
				return result
			}
		}
	}

	return result
}
