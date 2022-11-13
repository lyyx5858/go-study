package main

import "fmt"

func main() {

	a := [2][3]int{{1, 2, 3}, {4, 5, 6}} //表示a数组里面有2个长度为3的一维数组
	// a[0]: 1,2,3, a[0][0]表示第一个数组里的第一个元素，这是就是1
	// a[1]: 4,5,6

	fmt.Printf("\n %p", &a[0])    //a[0]里实际存放的是第1个数组的第一个元素的地址，a[1]里放的是第二个数组的第一个元素的地址
	fmt.Printf("\n %p", &a[0][0]) //地址如上一样
	fmt.Printf("\n %p \n\n", &a[1])

	//二维切片
	var b [][]int
	fmt.Printf("b's value:%p,  b's address:%p \n", b, &b)
	b = make([][]int, 2)
	fmt.Printf("b's value:%p,  b's address:%p \n\n", b, &b)

	fmt.Println("Before make b[0]")
	fmt.Printf("b[0]'s value:%p,  b[0]'s address:%p \n", b[0], &b[0])

	fmt.Println("After make b[0]")
	b[0] = make([]int, 2)
	fmt.Printf("b[0]'s value:%p,  b[0]'s address:%p \n", b[0], &b[0])

}
