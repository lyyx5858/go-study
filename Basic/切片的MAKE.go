package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println(os.Getpid())

	var b []int                                                      //这个语句会在内存中产生一个b, 并分配地址。但是这个b里的值是空的。
	fmt.Println("Before make")                                       //声明 b是个一维的整型切片
	fmt.Printf(" b's address: %p,  b's value: %p <-底层数组地址\n", &b, b) //b本质上是个指针，b里的内容应该是个地址，而这个地址应该批向一个底层数组的地址。

	b = make([]int, 1) //这个语句会在内存中产生一个b, 并分配地址。并且这个b里是有值的。这个值就是底层数组的地址。
	// b[0] = 1         				 //这句话如果在MAKE之前会报错的。
	//因此 下句话会报错的。
	//但是在没有用MAKE语句之前， 可以看出来b本身是有个地址了，但是应该指向底层数组的地址为0。
	fmt.Println("After make") //声明 b是个一维的整型切片
	fmt.Printf(" b's address: %p,  b's value: %p <-底层数组的地址\n", &b, b)

	b[0] = 1 //这句话不会报错了
	fmt.Println(b[0])

	b[0] = 2
	fmt.Println(b[0])

	c := []int{}
	fmt.Printf(" c's address: %p,  c's value: %p %#v\n", &c, c, c)

	d := make([]int, 0, 0)
	fmt.Printf(" d's address: %p,  d's value: %p %#v\n", &d, d, d)
	d = append(d, 1)
	fmt.Printf(" d's address: %p,  d's value: %p %v\n", &d, d, d)
	d = append(d, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9)
	fmt.Printf(" d's address: %p,  d's value: %p %v\n", &d, d, d)
	fmt.Println("可以看出，d本身地址没有变，但是d里面存的地址在变")
}
