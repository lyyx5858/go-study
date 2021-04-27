package main

import (
	"fmt"
)

type Tester int //)Tester目前是一个新定义的type

func (a Tester) Less(b Tester) bool { //Less是一个method,它的值是个bool型
	return a < b
}

func main() {
	var a Tester = 1

	if a.Less(2) {
		fmt.Println(a, "Less 2")
	}

}
