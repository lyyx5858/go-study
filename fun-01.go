package main

import "fmt"

func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", t1)
	fmt.Printf("%T\n", oper)


	res1 := add
	res1 = sub
	res1 = t1
	res1 = t2  //上面的t1可以赋值给res1的原因是，t1和res1有着相同的输入和输出，而t2就不能赋值给res1

	fmt.Println(res1)

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	res3 := oper(10, 20, sub)
	fmt.Println(res3)

	fun1 := func(a, b int) int {
		return a * b
	}

	res4 := oper(2, 3, fun1)
	fmt.Println(res4)

	res5 := oper(8, 4, func(a, b int) int {
		fmt.Println("除法")
		return a / b
	})

	fmt.Println(res5)

}
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b) //此处利用res调用了函数fun,这个函数的参数是与它一起传进来的a， b
	return res
}

//当一个函数具有相同的输入和输出时，就称这两个函数具备相同的签名。
//如果将函数看做变量的话，那么此时这两个变量就是同样类型的。

func sub(a, b int) int {
	return a - b
}

func add(a, b int) int {
	return a + b
}

func t1(a, b int) int {
	return b
}

func t2(a int) int {
	return a
}
