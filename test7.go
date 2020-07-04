package main

import (
	"fmt"
)

func main() {

	func(data int) {
		fmt.Println("hello", data)
	}(100)

}
