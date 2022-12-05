package main

import "fmt"

func main() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

func increaseA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}
