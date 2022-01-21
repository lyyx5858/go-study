package main

import (
	"fmt"
)

type Liu int //Liu 是个本地定义的数据类型，其实就是int的别名
type Wang string

type interface1 interface { //所有实现了方法method1和method2的变量都可以用interface1接口。
	method1() //注意，必须实现所有的方法,只实现一个不行。
	method2()
}

func test(i interface1) { //i的数据类型是interface1，
	fmt.Println(i)
	i.method1() //接口i直接调用方法method1
	i.method2()
}

func main() {

	var l1 Liu //l1是数据类型Liu的实例
	var w1 Wang

	l1 = 1
	w1 = "好好学习"

	test(l1) //test函数的形参是interface 1,而interface 1有两个方法
	//由于类型Liu的数据类型已经实现了method1，因此所有数据类型Liu都可以用interface1接口来调用test1函数
	test(w1)
}

func (l Liu) method1() {
	fmt.Println("L's method1", l)
}
func (l Liu) method2() {
	fmt.Println("L's method2", l)
}

//=========================================================================
func (w Wang) method1() {
	fmt.Println("W's method1", w)
}
func (w Wang) method2() {
	fmt.Println("W's method2", w)
}
