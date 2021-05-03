package main

import (
	"fmt"
)

type liu int //注意此处，liu其实是int的一个别名而已，但是local type就可以有方法了(如下面的method1)。也就是说，方法不一定非得是结构独有的。

type wang struct {
	liu //结构wang的匿名字段
	test liu //注意与上行的区别
}

type zhang func(int)(int)

func main() {

	var x liu
	x = 2
	fmt.Printf("%T %v \n", x, x)
	x.method1(1)

	var w1 wang

	w1.method1(2) //方法的继承
	w1.liu.method1(3) //注意此行与上行用区别

	w1.test.method1(4) //注意此行与上行用区别
	w1.method2(5) //注意此行与上行用区别
	w1.method2(6).method1(7) //注意此行是先执行method2，然后返回的liu型值，然后再执行liu的方法method1

	z1:=func(a int)int{return a+a}
	z2:=zhang(z1)
	z2.method(7)


}


func (l liu) method1(x int) {
	fmt.Println("liu's method",x)
}

func (w wang)method2(x int) liu {
	fmt.Println("wang's method",x)
	var l3 liu
	l3=3
	return l3  //此处，直接写成return 5也是可以工作的。

}

func (z zhang)method(x int){
	fmt.Println("Zhang's Method",x)

}