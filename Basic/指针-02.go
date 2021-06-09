package main

import (
	"fmt"
)

func main() {
	s0 := [3]int{9,8,7}
	r0 := s0
	r0[0] = 12
	fmt.Println(s0)

	//上面是数组，下面是切片
	s := []int{9, 8, 7}
	r := s //当将切片s赋值给r时，s切片和r切片隐含的指针都指向了同一个底层数组{9，8，7}，这时对r的任何改动都会反映到s
	r[0] = 12
	fmt.Println(s)
}
