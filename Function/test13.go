package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func test1() (func() int, func() int) {
	var y int
	y = y + 1
	return func() int { return y * y }, func() int { return y + y }
}

func main() {
	c := squares
	b, a := test1()
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
	fmt.Println(b, a)
	fmt.Println(b(), a())
	fmt.Println(c)
}
