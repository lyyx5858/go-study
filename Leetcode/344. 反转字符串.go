package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	s1 := string(s)
	s2 := []byte(s1)
	fmt.Println(s, s1, s2)

	reverseString(s2)
}

func reverseString(s []byte) {
	fmt.Println(s)
	l := len(s) >> 1 //找出中间位置
	for i := 0; i <= l-1; i++ {
		//tmp := s[i]
		//s[i] = s[len(s)-1-i]
		//s[len(s)-1-i] = tmp

		//上面三句与下面一句作用相同：a,b=b,a 就是交换a，b值
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	fmt.Println(s)

}

func reverseString01(s []byte) {
	left := 0           //左指针
	right := len(s) - 1 //右指针
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
