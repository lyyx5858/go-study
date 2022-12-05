package main

import "fmt"

func main() {
	v1 := [4]int{1, 2, 3, 4} //v1是个数组，方括号里有值
	fmt.Println(v1)

	var s1 []int
	fmt.Println(s1)

	s2 := []int{0, 1, 2, 3, 5}
	fmt.Printf("\n s2 content: %p  s2 address:%p", s2, &s2) //s2本质是一个指向底层数组的指针，当然s2这个指针本身也有个地址，就是&s2

	s2 = append(s2, 6, 7)
	fmt.Printf("\n s2 content: %p  s2 address:%p same as above", s2, &s2) //当追加元素时，s2里面存的指针指向了新的地址，但是s2本身的地址不会变的，所以此处打印的&s2与上面相同。

	s2 = s2[:2] //表示截取元素0，1，但不包括2,方括号的左边是闭区间，右边是开区间
	fmt.Printf("\n s2 content: %p  s2 address:%p \n", s2, &s2)
	fmt.Println(s2)

}
