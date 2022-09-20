package main

import (
	"fmt"
)

type Liu int //Liu 是个本地定义的数据类型，其实就是int的别名
func (l Liu) method1() {
	fmt.Println("牛排")
}

type Wu int // 是个本地定义的数据类型，其实就是int的别名
func (w Wu) method1() {
	fmt.Println("青菜")
}

type interFunc func(int)

func (if1 interFunc) method1() {
	fmt.Println("转换")
}

type interface11 interface { //interface1实现了方法method1
	method1()
}

type f31 func(interface11) //f3是一个函数类型，这个函数只有一个输入参数：接口 interface11

//hi, baby, if you have this method, you are my type.

func main() {
	var l1 Liu //l1是数据类型Liu
	var w1 Wu  //w1是数据类型Wu

	var f1 func(interface11) //f1是个函数类型
	//var f2 func(i interface1)  //注意与上行的区别

	var f3real f31 //注意与f1, f31和f3real的区别与联系，f1是一个实体，类型是函数， f31，是一个虚例，只是定义而已，
	//而f3real是对f31的实现，

	fmt.Printf("i41   %v\n", f3real)
	fmt.Printf("-------%T \n", l1.method1)
	fmt.Printf("-------%v \n", w1.method1)
	fmt.Printf("---cook----%v \n", cook)

	fmt.Printf("----l1---%v \n", &l1)
	fmt.Printf("-----w1--%v \n", &w1)

	i3 := w1
	f1 = cook //其实这句话的实质就是将cook的地址赋给f1
	f4 := f1
	f3real = cook
	f3real(i3)

	fmt.Printf("----f1---%v %v \n", f1, f4)
	fmt.Printf("----i3---%v \n", &i3)
	fmt.Printf("----i3---%T \n", i3)
	fmt.Printf("----i3---%v \n", i3.method1)

	cook(l1)
	cook(w1)
	cook(i3)
	f1(i3)

	fmt.Println(interFunc(x11).method1) //x11只是一个普通函数，但是被interFunc类型转换后，它就变成了一个接口，有了方法。
	cook(interFunc(x11))                //强制转换后，x11这个函数可以当做一个接口做为cook的形参了。
}

func cook(i interface11) { // 函数test的形参是接口 i，i的数据类型是interface1, 由于该接口已经实现了方法method1, 因此可以直接调用
	i.method1() //接口i直接调用方法method1
}

func tt(i interface11) {

}

func f3(interface1, a int) int {
	return a
}
func f4(s func(int)) interface11 {
	return interFunc(s)
}

func x11(int) {

}
