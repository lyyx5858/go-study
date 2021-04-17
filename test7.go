package main

import (
	"fmt"
)

//匿名函数

func main() {

	func(date int) {
		fmt.Println("hello", date)
	}(100)

}
