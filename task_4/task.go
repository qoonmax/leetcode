package main

import "fmt"

//https://leetcode.com/problems/valid-parentheses/

func main() {
	fmt.Println(isValid("{}"))
}

func isValid(s string) bool {
	bracketMap := map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}

	var stack []string

	for _, bracket := range s {
		_, ok := bracketMap[string(bracket)]
		if ok {
			stack = append(stack, string(bracket))
		} else {
			if len(stack) > 0 && bracketMap[stack[len(stack)-1]] == string(bracket) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
