package main

import "fmt"

//闭包

func main() {

	res1 := increment //可以理解为将increment这个函数的地址给res1
	fmt.Println(res1) //res1其实就是increment

	res2 := increment() //此处和res1不一样，
	//此处的含义是将increment调用返回值赋给res2,
	//而increment的返回值就是匿名函数fun的地址

	fmt.Println(res2) //此处为匿名函数的地址
	fmt.Println(res2())
	fmt.Println(res2())

	res3 := increment()
	fmt.Println(res3)   //res3再次调用increment的时候，res3又指向了匿名函数fun，但是其中的i重新创建了。
	fmt.Println(res3()) //此处打印的值是新的i

}

func increment() func() int {
	i := 0
	f := func() int { //这个匿名函数里没有声明i,但却可以使用i,非但可以使用，而且可以带出去
		i++
		return i
	}
	return f
}
