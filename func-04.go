package main

import "fmt"

func main() {
	var a int
	var c func(int)  (int,int)   //第一个int是输入，后面两个int是输出
	//此处的func是关键字，c的类型是func(int) int,由于hanshu这个函数的输入和输出参数与c完全相同，因此可以将hanshu赋值给
	//此处的c就是所谓的“函数变量”


	a = 1
	fmt.Println(&a)
	fmt.Printf("The type of a is: %T\n", a)
	fmt.Printf("the value of a is: %d\n", a)
	fmt.Println(a)
	fmt.Printf("没调用的func1的值: %T, %v, \n", hanshu, hanshu)
	c=hanshu //注意此处，赋值给C时，不能代()，带()后，就是将的返回值赋给C(系统会报错)，而不带()表示将hanshu赋给c
	x,y:=hanshu(a)
	fmt.Printf("调用后func1的类型 %T ,%v,%v, \n", hanshu, x,y)
	x,y=c(a)
	fmt.Printf("调用后func1的类型 %T ,%v,%v, \n", c, x,y)
}



func hanshu(a int) (int, int) {
	fmt.Printf("函数内部a的值： %T %d %v \n", a,a,&a)
	return a,a+1
}
