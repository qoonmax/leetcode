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
		for {
			x := rand.Int()
			if _, ok := m[x]; !ok {
				m[x] = struct{}{}
				sl[i] = x
			}
			break
		}
	}

	return sl
}
