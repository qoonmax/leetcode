package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(generateSlice(5))
}

func generateSlice(n int) []int {
	sl := make([]int, n)
	m := make(map[int]struct{})

	for i := 0; i < n; i++ {
		x := rand.Int()
		for _, ok := m[x]; ok; {
			x = rand.Int()
		}
		sl[i] = x
	}

	return sl
}
