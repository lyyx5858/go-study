package main

import "fmt"

func main() {
	i := 0
	test(&i)
	fmt.Println(i)
}

func test(x *int) int {
	fmt.Println(*x)
	*x = *x + 1
	return *x
}
