package main

import "fmt"

func main() {
	var a, x, y int
	var c func(int) (int, int) //第一个int是输入，后面两个int是输出
	//此处的func是关键字，c的类型是func(int) int,由于hanshu这个函数的输入和输出参数与c完全相同，因此可以将hanshu赋值给
	//此处的c就是所谓的“函数变量”

	a = 1
	x = 2
	y = x //变量的赋值：是将x这块内存区域的值2拷贝给y这个内存区域，从下面这句话可以看出，x and y的地址是不一样的。但是函数赋值不同！
	fmt.Println(&a, &x, &y)
	fmt.Printf("The type of a is: %T\n", a)
	fmt.Printf("the value of a is: %d\n", a)
	//
	fmt.Printf("调用前的hanshu的值: %T, %v, \n", hanshu, hanshu)
	//
	fmt.Printf("赋值前c: %T, %v, \n", c, c)
	c = hanshu //注意此处，赋值给C时，不能代()，带()后，就是将函数的返回值赋给c(系统会报错)，而不带()表示将hanshu赋给c
	fmt.Printf("赋值后c: %T, %v, \n", c, c)
	//从上面可以看出，C的值从赋值前的空变成赋值后的与hanshu相同的地址。所以说，函数变量的赋值更象指针赋值。而函数变量的赋值更象给函数起了个别名
	x, y = hanshu(a)
	fmt.Printf("调用后hanshu的类型和值 %T ,%v, \n", hanshu, hanshu)
	x, y = c(a)
	fmt.Printf("c的类型和值 %T ,%v,%v, \n", c, x, y)
}

func hanshu(b int) (int, int) {
	fmt.Printf("函数内部a的值： %T %d %v \n", b, b, &b)
	return b, b + 1
}
