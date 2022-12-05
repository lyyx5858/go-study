package main

import "fmt"

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())

}

func test1() (x int) {
	i := 0
	defer func() {
		i++
	}()
	return i //因为返回值是x，所以这句话其实隐含了一句话：x=i, (将i赋值给x，而这句赋值发生在defer之前），而defer加的是i，不是x
}

/*在go语言中，return这句话不是原子操作，而是分为三步：
第一步：将i赋值给x，因为函数定义时返回的是x，而return语句中写的是i，而在test2程序中，则没有这一步。
第二步：defer 语句
第三步：返回x值。

*/

func test2() (i int) {
	i = 0
	defer func() {
		i++
	}()
	return i //与test1没有任何改动，只不过函数定义时返回值是i，return里也是i，defer加的是i，返回的是i，因此返回了1
	//这里的写return和return i结果是一样的
}

func test3() int {
	i := 0
	defer func() {
		i++
	}()
	return i //与test1没有任何改动，只不过函数定义时返回值是匿名，但其实与test1一样流程，go会产生一个临时变量，如同临时的x
}
