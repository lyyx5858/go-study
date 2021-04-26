package main

import (
	"fmt"
	"math"
)

type square struct { //正方形
	side float64
}

func (s square) area() float64 { //正方形面积
	return s.side * s.side
}

type circle struct { //圆
	radius float64
}

func (c circle) area() float64 { //圆形面积
	return math.Pi * c.radius * c.radius
}

type shape interface { //定义interface
	area() float64
}

func printArea(x shape) { //函数PrintArea的形参是一个接口
	fmt.Println("My area is:", x.area())
}

func main() {
	s1 := square{10}
	c1 := circle{4}
	printArea(s1) //因为结构s1实现了area方法，因此它成为接口的一个实例
	printArea(c1)
	//这种以接口为参数的函数，做到了与具体方法的解耦，以后如果打印一个三角形的面积，只要在三角形的方法中，实现area()就行了，而不需要改动printArea函数了。
}
