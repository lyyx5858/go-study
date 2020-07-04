package main

import (
"fmt"
)

type Liu int //Liu 是个本地定义的数据类型，其实就是int的别名

type interface1 interface {  //所有实现了方法method1的变量都可以用interface1接口。
	method1()
}

func main() {

	var l1 Liu //l1是数据类型Liu
	test(l1)  //由于类型LIU的数据类型已经实现了method1，因此所有数据类型LIU都可以用interface1接口来调用test1函数。
	fmt.Printf("%T %v %v \n", l1, l1, &l1)
}

func (l Liu)method1(){
	fmt.Println("this is method", l)
}

func test(i interface1) { // 函数test的形参是接口 i，i的数据类型是interface1，而interface1的接口定义中：必须实现方法method1：见程序开并头！因此，所有实现了method1的数据类型都可调用函数test1.
	
	i.method1() //接口i直接调用方法method1

}

// 假如有个男人，性别是：男，有两个方法，分别是： 吃（）{馒头}，睡（）{硬板床}
// 假如有个女人，性别是：女，有两个方法，分别是： 吃（）{水果}，睡（）{席梦思}
// 现在定义一个接口：人。这个接口有两个方法，分别是：吃（），睡（）：
// var 人 interface {
// 吃（）
// 睡（）
// }
// 现在写一个函数，函数的名字是：打印睡和吃（ r 人）。其传入参数是接口：r, 而r的类型是interface：人，函数内容如下：
// r.吃（），r.睡（）
// 则调用这个函数时，如果传进去的接口是男人的实例，则打印的是：馒头和硬板床
// 调用这个函数时，如果传进去的接口是女人的实例，则打印的是：水果和席梦思
// 用接口的好处是：我写："打印睡和吃"这个函数时，不需要具体再区别是男人还是女人，其实这个函数将男人和女人的共同点：吃和睡抽象为"人"这个接口，然后对其进行操作。这样就不需要分别为男人和女人单独写一个函数了。 
