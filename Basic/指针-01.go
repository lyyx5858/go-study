package main

import (
	"fmt"
)

type test1 struct {
	a int
}

type test3 struct {
	b int
	c *test1 //这里的含义是，定义了一个可以存放 "test1结构体地址" 的结构体指针变量c,所谓指针变量，就是这个内存里存放的是个地址。
}

func main() {
	s := 7
	fmt.Println(s, &s)

	var p *int
	p = &s //p是个int的指针变量。
	fmt.Println(p, &p)

	var r int
	r = *p //p的地址假如是：0x123, 而0x123这个地址里存的值是：0x456, *p就是: 0x456这个地址里存储的值。
	//*p 首先p的值必须是个地址（当然p自己本身也有个地址），然后就是取p这个值：地址所指向的地址里面存放的值。
	//0x123[0x456]---->0x456[100]
	//[]号里面是值，外面是地址,则*p的结果就是100，p的值是0x456,p的地址是0x123

	fmt.Println(r, &r)
	r = 11
	fmt.Println(s, p, r)

	t1 := test1{1}      //结构体是值类型，与int一样
	t3 := test3{2, &t1} //t3的第二个成员是个结构体指针，它存放了结构体t1的地址。

	fmt.Printf("=====================\n")
	fmt.Printf("t1: %v,%p\n", t1, &t1)
	fmt.Printf("t3: %v,%p\n", t3, &t3)

}
