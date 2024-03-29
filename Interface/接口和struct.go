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

type jiegou1 struct {
	interface1
}

type jiegou2 struct {
	int
}

func test(i interface1) { //i的数据类型是interface1，
	fmt.Println(i)
	i.method1() //接口i直接调用方法method1
	i.method2()
}

func main() {

	var l1 Liu //l1是数据类型Liu的实例
	var w1 Wang
	var x1 int

	var t interface1

	t = jiegou1{l1}

	x1 = 1
	w1 = "好好学习"
	l1 = Liu(x1) //当底层数据类型一样时，可以进行强制类型转换。因为x1和l1都是int.

	test(l1)
	test(w1)
	j1 := jiegou1{l1}
	j1.method1()

}

//==================================
func (l Liu) method1() {
	fmt.Println("L's method1", l)
}
func (l Liu) method2() {
	fmt.Println("L's method2", l)
}

//----------------------------------
func (w Wang) method1() {
	fmt.Println("W's method1", w)
}
func (w Wang) method2() {
	fmt.Println("W's method2", w)
}
