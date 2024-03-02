package main

import (
	"fmt"
	"strconv"
)

//https://leetcode.com/problems/palindrome-number/description/

func main() {
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	lenStrX := len(strX)
	iterations := lenStrX / 2

	if lenStrX%2 != 0 {
		iterations++
	}

	for i := 0; i < iterations; i++ {
		if strX[i] != strX[lenStrX-i-1] {
			return false
		}
	}

	return true
}
