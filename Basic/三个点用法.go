package main

import "fmt"

func main() {

	testFunc(1, 2, 3, 4, 5, 6)

	s0 := []int{1, 2, 3, 4, 5}
	s1 := []int{6, 7, 8}
	s0 = append(s0, s1...)   //s1后面必须加...,不然会报错
	s0 = append(s0, 6, 7, 8) //这句与上句结果相同
	fmt.Println(s0)

	t0 := []string{"a", "b", "c"}
	t1 := []string{"d", "e", "f"}
	t2 := append(t0, "def")
	fmt.Println(t2)
	t1 = append(t0, t1...)
	fmt.Println(t1)

}

func testFunc(i ...int) { //...表示参数个数可变，但都必须是int

	fmt.Println(i)
	fmt.Println(i[0])
	fmt.Println(len(i))
}
