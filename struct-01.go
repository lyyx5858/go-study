package main

import "fmt"

type s1 struct { //结构体
	age int
}

type s3 struct { //嵌套结构体
	s1         //等效为：t1 s1,此处省略了t1
	s2   [3]s1 //结构体数组
	name string
}

func main() {

	a1 := s1{18}

	a2 := [3]s1{s1{11},
		a1,
		s1{20}}

	a3 := s3{s1{30}, a2, "liuyan"}

	fmt.Println(a1)
	fmt.Printf("%T,%v", a2, a2)
	fmt.Println()
	fmt.Println(a3)

	fmt.Println(a3.s1.age)

}
