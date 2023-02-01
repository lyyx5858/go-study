package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	fmt.Println(findContentChildren(g, s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var x, j int
	for i := 0; i < len(s); i++ {
		if j < len(g) {
			if s[i] >= g[j] {
				x++
				j++
			}
		}
	}

	return x
}
