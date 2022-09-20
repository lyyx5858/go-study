package main

import "fmt"

func main() {

	a := new([]int) //初始化一个切片
	fmt.Printf("%T, %v, %v, %p  \n", a, a, &a, &a)

	b := make([]int, 1) //初始化一个长度为1的切片
	fmt.Printf("%T, %v, %v, %p   \n", b, b, &b, &b)

	var x int
	fmt.Printf("%T, %v %v \n", x, x, &x)
	x = 11

	var y *int
	fmt.Printf("%T, %v  %v \n", y, y, &y)
	y = &x
	fmt.Printf("%T, %v  %v \n", y, y, &y)
	fmt.Printf("%v \n", *y)

}
