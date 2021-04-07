package main

import (
	"fmt"
)

func main() {

	var n1, n2 sayname
	var zhangsan = Student{"张三", 20}
	var lilaoshi = Teacher{"李老师"}

	fmt.Println("用方法来实现：")
	fmt.Println(zhangsan.see_name)   //注意此处不带括号，是打印的是see_name这个方法的内存地址，
	fmt.Println(lilaoshi.see_name)   //这个地址与上面一行地址不同，表明不同结构的方法虽然名字一样，其实是不同的实体
	fmt.Println(lilaoshi.see_name()) //注意方法带括号和不带括号的区别

	//fmt.Printf("%T   %v\n",other,other)

	fmt.Println("===============\n下面用接口实现:")
	fmt.Printf("n1, n2的类型：%T  %T\n", n1, n2)
	fmt.Printf("n1, n2的值：%v  %v\n", n1, n2)
	fmt.Printf("n1, n2的值：%v  %v\n", &n1, &n2)

	//fmt.Printf("other的值：%v  类型%T\n", other, other)

	n1 = zhangsan //只要且必须 结构体zhangsan有n1所有的方法，它就可以满足n1,也就是可以给n1赋值
	n2 = lilaoshi

	var other = Otherpreson{
		n1,
		"柳林馆",
	}

	fmt.Printf("赋值以后，n2的：类型：%T       值：%v\n", n2, n2)
	fmt.Println(n2.see_name())

	fmt.Println(fayan(lilaoshi)) //fayan是个以接口为参数的函数
	fmt.Println(fayan(n2))       //这句与上行的区别是？

	fmt.Println(fayan(other.n))

}

//所有实现see_name方法的类型都叫做实现sayname接口，注意：实现接口的不一定全是结构体
type sayname interface { // name 是个接口
	see_name() string //see_name是个方法, 这个方法不光定义了名字，还有输入参数及类型，输出参数及类开型，
	//有相同输入入及输出的方法，也叫相同的方法
}

type saymorename interface {
	sayname
}

// Student 是个结构
type Student struct {
	n   string
	age int
}

// Techer 也是个结构
type Teacher struct {
	n string
}

type Otherpreson struct {
	n    sayname
	addr string
}

// 这是学生的see_name是个方法
func (s Student) see_name() string {
	return "我是学生，我的姓名是：" + s.n
}

//下面这个方法是不被允许的，虽然类型与上面不同，但名字也不能相同
//func (s Student) see_name() int{
//	return s.age
//}

//这是老师的see_name方法
func (t Teacher) see_name() string {
	return "我是老师，我的姓名是：" + t.n
}

// fayan是个函数，形参是接口name,函数返回一个字符串
func fayan(n sayname) string {
	return n.see_name()
}
