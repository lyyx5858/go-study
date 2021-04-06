package main

import "fmt"

func main() {
	var a int
	var c func(int)   //c的类型是func(int),由于func1这个函数的参数与c相同，因此可以将func1赋值给

	a = 1
	fmt.Println(&a)
	fmt.Printf("The type of a is: %T\n", a)
	fmt.Printf("the value of a is: %d\n", a)
	fmt.Println(a)
	fmt.Printf("没调用的func1的值: %T, %v, \n", func1, func1)
	c=func1 //注意此处，赋值给C时，不能代()，带()后，就是将func1的返回值赋给C(系统会报错)，而不带()表示将func1赋给c
	fmt.Printf("调用后func1的类型 %T ,%v, \n", c, c)

}



func func1(a int){
	fmt.Printf("函数内部a的值： %T %d %v \n", a,a,&a)
}
