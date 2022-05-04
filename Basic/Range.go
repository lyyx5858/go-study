package main

import "fmt"

func main() {
	var a [3]int = [3]int{7, 8, 9}
	b := []int{1, 2, 3}

	var c []int
	c = make([]int, 4)
	fmt.Printf("len:%d,cap:%d\n", len(c), cap(c))
	fmt.Printf("%p\n", c)
	fmt.Printf("%p\n", &c)

	c = append(c, 7, 8, 9)
	fmt.Printf("len:%d,cap:%d\n", len(c), cap(c))
	fmt.Printf("%p\n", c)
	fmt.Printf("%p\n", &c)

	//c = append(c, 5)

	//c := make([]int, 5)

	//fmt.Println("c is:", c[1])

	fmt.Println(a[0])
	fmt.Println(len(a))
	fmt.Println(b)

	for i, v := range a {
		fmt.Println("%d %d", i, v)
	}
}
