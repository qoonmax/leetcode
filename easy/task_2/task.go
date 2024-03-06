package main

import "fmt"

//https://leetcode.com/problems/roman-to-integer/description/

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var result int

	rMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var prev, curr int
	for i := len(s) - 1; i >= 0; i-- {
		curr = rMap[string(s[i])]

		if curr >= prev {
			result += curr
		} else {
			result -= curr
		}

		prev = curr
	}

	return result
}
