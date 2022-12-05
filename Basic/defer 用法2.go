package main

import "fmt"

func main() {
	fmt.Println(test10())
	fmt.Println(test20())
}

func test1() (res int) {
	res = 1
	defer func() {
		res++
	}()
	return 0
}

func test10() (res int) { //test1=test10
	res = 1
	res = 0 //赋值
	func() {
		res++ //defer
	}()
	return res //返回
}

func test2() (res int) {
	tmp := 1
	defer func() {
		tmp++
	}()
	return tmp
}

func test20() (res int) { //test2=test20
	tmp := 1
	res = tmp
	func() {
		tmp++
	}()
	return res
}
