package main

import "fmt"

type s1 struct { //结构体
	age    int
	string //结构体成员可以匿名，只有类型，但相同类型只能有一个，下面S3的例子也一样
}

type s3 struct { //嵌套结构体
	s1 //等效为：t1 s1,此处省略了t1
	//结构体的成员的数据类型在结构体中唯一时，可以省略名字！
	s2   [3]s1 //结构体数组
	name string
	t    *s1
}

func main() {

	a1 := s1{18, "yan"}

	a2 := [3]s1{s1{11, "wang"},
		a1,
		s1{20, "li"}}

	a3 := s3{
		s1:   s1{30, "test"},
		s2:   a2,
		name: "liuyan",
		t:    &s1{18, "wang"}}

	fmt.Println("a1=", a1)
	fmt.Printf("a2= %T,%v\n", a2, a2)
	fmt.Println()
	fmt.Println("a3=", a3)

	fmt.Println("a3.s1.age=", a3.s1.age)

}
