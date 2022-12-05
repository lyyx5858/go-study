package main

import (
	"fmt"
)

//map 是引用类型，初值nil, 用之前必须初始化
//所谓引用类型就是个指针类型变量。

// map1 declare但是没有make：  0x1234[nil]
//map1 如果初始化后：           0x1234[0x4567]

//0x1234是map1的地址，

func main() {
	var map1 map[int]string //方括号里是KEY， 外面是值
	//map1[1] = "hello" //这句话报错

	var map2 = make(map[int]string)
	map2[1] = "hello"

	var map3 = map[int]string{186: "li", 139: "liu"}

	//map1-3是三种初始化方法

	fmt.Printf("%v, %p \n", map1, map1)
	fmt.Printf("%v, %p \n", map2, map2)
	fmt.Printf("%v, %p \n", map3, map3)

	value, ok := map2[1] //ok是bool类型的变量
	fmt.Println(value, ok)

	for key, value := range map2 {
		fmt.Println(key, value)
	}

}
