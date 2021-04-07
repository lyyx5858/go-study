package main

import "fmt"

func main() {

	var a [3]int = [3]int{7, 8, 9}
	b := []int{1, 2, 3}

	fmt.Println(a[0])
	fmt.Println(len(a))
	fmt.Println(b)

	for i, v := range a {
		fmt.Println("%d %d", i, v)
	}
}
