package main

import (
	"fmt"
)

type Liu int //Liu 是个本地定义的数据类型，其实就是int的别名
type Wang int

type interface1 interface { //所有实现了方法method1和method2的变量都可以用interface1接口。
	method1() //注意，必须实现所有的方法,只实现一个不行。
	method2()
}

type struct1 struct {
	i1 interface1
	interface2
}

type interface2 interface {
	method1()
	method2()
	method3()
}

func main() {

	var l1 Liu //l1是数据类型Liu的实例
	test(l1)   //由于类型Liu的数据类型已经实现了method1，因此所有数据类型Liu都可以用interface1接口来调用test1函数
	fmt.Printf("%T %v %v \n", l1, l1, &l1)

	var w1 Wang
	test(w1) //注意：test函数的入参数应该是interface1类型, 但此处调用的是w1有三个方法,但是由于w1实现了inteface1的所有方法
	//因此也可以直接调用。

	//一个函数的输入参数是个接口的情况下，如果一个实例，只要它实现了这个接口的所有方法，它就可以被输入。
	//举例：interface1有两个方法，inteface2有三个方法，所有以interface1的函数都可以用interface2.
	//但用interface2的函数不用用inteface1,因为它少了一个方法。

	var s1 struct1
	s1.i1 = l1
	s1.interface2 = w1

	test(s1)

}

func (l Liu) method1() {
	fmt.Println("this is method1", l)
}
func (l Liu) method2() {
	fmt.Println("this is method2", l)
}

func (w Wang) method1() {
	fmt.Println("W: this is method1", w)
}
func (w Wang) method2() {
	fmt.Println("W: this is method2", w)
}
func (w Wang) method3() {
	fmt.Println("W: this is method3", w)
}

func test(i interface1) { // 函数test的形参是接口 i，i的数据类型是interface1，而interface1的接口定义中：必须实现方法method1：见程序开并头！因此，所有实现了method1的数据类型都可调用函数test1.

	i.method1() //接口i直接调用方法method1
	i.method2()

}

// 假如有个男人，性别是：男，有两个方法，分别是： 吃（）{馒头}，睡（）{硬板床}
// 假如有个女人，性别是：女，有两个方法，分别是： 吃（）{水果}，睡（）{席梦思}
// 现在定义一个接口：人。这个接口有两个方法，分别是：吃（），睡（）：
// var 人 interface {
// 吃（）
// 睡（）
// }

//因为男人这个结构体有已经有两个方法吃（）和睡（），因此可以称：男人实现了这个接口。也可以称：男人是这个接口的实例。 可以理解为：男人（具体）是人（抽象）的一个实例。
// 现在写一个函数，函数的名字是：打印睡和吃（ r 人）。其传入参数是接口：r, 而r的类型是interface：人，函数内容如下：
// r.吃（），r.睡（）
// 则调用这个函数时，如果传进去的接口是男人的实例，则打印的是：馒头和硬板床
// 调用这个函数时，如果传进去的接口是女人的实例，则打印的是：水果和席梦思
// 用接口的好处是：我写："打印睡和吃"这个函数时，不需要具体再区别是男人还是女人，其实这个函数将男人和女人的共同点：吃和睡抽象为"人"这个接口，然后对其进行操作。这样就不需要分别为男人和女人单独写一个函数了。
