package main

import "fmt"

type s struct {
	Host string
}

type t12 func(s1 *s)

func main() {

	host := "test"
	s3 := &s{}

	f12 := f1()
	f12(s3) //此时，f1()的调用已经结了，但是由于f1返回函数f12的存在，其实f1还是存在的，并且在这句话中，将s3进行赋值。

	fmt.Println(s3.Host)
	func() {
		fmt.Println(host) //此处，这个匿名函数内部可以访问外部的host变量
	}()

}

func f1() t12 {
	host := "内部变量"
	f := func(s1 *s) {
		s1.Host = host //此处，这个匿名函数内部可以访问外部的host变量
	}
	return f
}
