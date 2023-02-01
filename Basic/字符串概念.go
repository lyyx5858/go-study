package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "abcd"
	fmt.Printf("%T, %v", s, s)
	fmt.Printf("\n %T, %v", s[0], s[0])

	var s1 byte = 98
	fmt.Printf("\n %T, %v", s1, s1)

	s2 := []byte(s) //字符串与字节数组之间的相互转换
	s = string(s2)

	fmt.Printf("\n %T, %v %p", s[0], s[0], &s)       //此处可以取s的地址，但是不能取s[1]的地址，因此&s[0]是错的
	fmt.Printf("\n %T, %v %p", s2[0], s2[0], &s2[0]) //但是&s2[0]是可以的。

	fmt.Println()
	fmt.Println(string(97))         //将数字97转换成为小字字母：a
	fmt.Println([]byte("a")[0])     //将a转换为97
	fmt.Println(strconv.Itoa(97))   //将数字97转换为字符串“97”
	fmt.Println(strconv.Atoi("97")) //将字符串“97“转换为数字 97
}
