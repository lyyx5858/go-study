package main

import (
	"fmt"
)

type type01 func(x int, y int) int

func test(a1 int, b1 int) (c1 int) {
	c1 = a1 + b1
	return c1
}

func test2(a2 int, b2 type01) int {
	return a2 + b2(1, 1)
}

func bb(a3 int, b3 int) int {
	return a3 + b3
}

func main() {

	p := test2(1, bb)

	fmt.Print("\n", test(1, 2))
	fmt.Print("\n", p)
}
