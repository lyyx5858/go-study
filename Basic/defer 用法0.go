package main

import (
	"fmt"
)

func main() {

	i := 10
	j := hello(&i)
	fmt.Println(i, j)
}

func hello(i *int) int {
	//此处的i和在main里的i不是一会事
	//在main里是int，在这里是指针
	defer func() {
		*i = 19
	}()
	return *i
}

//上面的函数的效果与下面一样
func hello1(i *int) int {
	x := *i //x=10
	func() {
		*i = 19
	}()
	return x //x=10
}
