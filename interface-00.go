package main

import (
	"fmt"
)

type Liu int //Liu 是个本地定义的数据类型，其实就是int的别名
type Wu int // 是个本地定义的数据类型，其实就是int的别名

type interface1 interface { //interface1实现了方法method1
	method1()
}

//hi, baby, if you have this method, you are my type.

func main() {
	var l1 Liu //l1是数据类型Liu
	var w1 Wu //w1是数据类型Wu
	var i3 interface1 //l3是接口
	var f1 func(i interface1) //f1是个函数类型

	fmt.Printf("-------%T \n", l1.method1)
	fmt.Printf("-------%v \n", w1.method1)
	fmt.Printf("---cook----%v \n", cook)
	fmt.Printf("----i3---%v \n", &i3)
	fmt.Printf("----l1---%v \n", &l1)
	fmt.Printf("-----w1--%v \n", &w1)

	i3=w1
	f1=cook //其实这句话的实质就是将cook的地址赋给f1
	f2:=f1

	fmt.Printf("----f1---%v %v \n", f1,f2)
	fmt.Printf("----i3---%v \n", &i3)
	fmt.Printf("----i3---%T \n", i3)
	fmt.Printf("----i3---%v \n", i3.method1)

	cook(l1)
	cook(w1)
	cook(i3)
	f1(i3)
}

func (l Liu) method1() {
	fmt.Println("牛排")
}
func (w Wu) method1() {
	fmt.Println("青菜")
}

func cook(i interface1) { // 函数test的形参是接口 i，i的数据类型是interface1, 由于该接口已经实现了方法method1, 因此可以直接调用
	i.method1() //接口i直接调用方法method1
}
