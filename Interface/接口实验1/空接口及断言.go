package main

import "fmt"

type Book struct {
	auth string
}

func main() {
	book := Book{"Golang"}

	myFunc(book)
	myFunc(100)
	myFunc(3.14)
	myFunc("abc")
}

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("===================")
	fmt.Println(arg)

	//interface{} 改如何区分 此时引用的底层数据类型到底是什么？

	//给 interface{} 提供 “类型断言” 的机制
	value, ok := arg.(string) //当断言成功后，arg的值会赋给value
	if !ok {
		fmt.Println("断言失败，arg is not string type")
	} else {
		fmt.Println("断言成功，arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}
