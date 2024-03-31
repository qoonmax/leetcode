package main

import "fmt"

//https://leetcode.com/problems/container-with-most-water/description/

func main() {
	fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
	var rString string

	var a = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var b = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for num > 0 {
		for i := 0; i < len(a); {
			if num >= a[i] {
				rString += b[i]
				num -= a[i]
			} else {
				i++
			}
		}
	}

	return rString
}
