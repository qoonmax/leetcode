package main

import (
	"fmt"
)

//https://leetcode.com/problems/length-of-last-word/description/

func main() {
	fmt.Println(lengthOfLastWord("a"))
}

func lengthOfLastWord(s string) int {
	var x []byte
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if len(x) > 0 {
				return len(x)
			}
			continue
		} else {
			x = append(x, s[i])
		}
	}

	return len(x)
}
