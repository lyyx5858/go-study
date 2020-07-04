package main

import(
	"fmt"
)
// Student 是个结构
type Student struct{
	name string
}
// Techer 也是个结构
type Teacher struct{
	name string
}
// see_name是个方法
func (s Student) see_name() (string){
	return "我是学生，我的姓名是：" + s.name
}

func (t Teacher) see_name() (string){
	return "我是老师，我的姓名是：" + t.name
}

type name interface{ 		// name 是个接口
	see_name() string   	//see_name是个方法
}


// fayan是个函数，
func fayan(p name) (string) {
	return p.see_name()
}

func main(){

	var zhangsan = Student{"张三"}
	var lilaoshi = Teacher{"李老师"}


	fmt.Println(zhangsan.see_name())  //这是用方法来实现see_name
	fmt.Println(lilaoshi.see_name())

//===========n1 是name类型=============================================================

	var n1, n2 name
	fmt.Printf("%T\n",n1)
	n1 =zhangsan //只要且必须 zhangsan有n1所有的方法，它就可以满足n1,也就是可以给n1赋值
	n2=lilaoshi
	fmt.Printf("%T%v\n",n2,n2)
	fmt.Println(n1.see_name())
	fmt.Println(fayan(n2))


}