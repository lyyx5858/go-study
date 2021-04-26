package main

import (
	"fmt"
)

type liu int //注意此处，liu其实是int的一个别名而已，但是local type就可以有方法了(如下面的method1)。也就是说，方法不一定非得是结构独有的。

type wang struct {
	liu //结构wang的匿名字段
}

func main() {

	var x liu
	x = 2
	fmt.Printf("%T %v \n", x, x)
	x.method1(1)

	var w1 wang
	w1.method1(2) //方法的继承
	w1.liu.method1(3) //注意此行与上行用区别

}

func (l liu) method1(x int) {
	fmt.Println("this is method",x)
}
