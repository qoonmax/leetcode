package main

import (
	"fmt"
)

//https://leetcode.com/problems/length-of-last-word/description/

func main() {
	fmt.Println(lengthOfLastWord("a"))
}

func lengthOfLastWord(s string) int {
	var counter int
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == " " {
			if counter > 0 {
				return counter
			}
			continue
		} else {
			counter++
		}
	}

	return counter
}
