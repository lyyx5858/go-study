package main

import "fmt"

func main() {
	fmt.Printf("%v\n", add)
	fmt.Printf("%v\n", sub)
	fmt.Printf("%T\n", oper)

	res1 := add(1, 2)

	fmt.Println(res1)

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	res3 := oper(10, 20, sub)
	fmt.Println(res3)

	fun1:=func(a,b int)int{
		return a*b
	}

	res4:=oper(2,3,fun1)
	fmt.Println(res4)

	res5:=oper(8,4,func(a,b int)int {
		fmt.Println("除法")
		return a / b
	})

	fmt.Println(res5)

}
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}



