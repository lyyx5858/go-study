package main

import "fmt"

func main() {
	s := "the   sky is blue"
	fmt.Println(reverseWords01(s))
}

func reverseWords(s string) string {
	s1 := append([]byte(" "), []byte(s)...) //在头部加一个空格，以方便分析
	r := len(s1) - 1                        //右指针
	l := r                                  //左指针和右指针同时在串的最右边
	var s2 []byte

	//右指针遇到非空就停，而左指针继续左移，直至另外一个空格
	for l >= 0 {
		if s1[l] == 32 && s1[r] == 32 { //当左、右指针都是空格时，同时向左挪动
			l--
			r--
		} else if s1[l] != 32 { //当左指针不是空格时，只挪动左指针
			l--
		} else { //此时，左指针在空格处，右指针在非空格处
			s2 = append(s2, s1[l+1:r+1]...)
			s2 = append(s2, ' ')
			r = l //右指针归位到左指针处，重新开始分析
		}

	}
	s2 = s2[0 : len(s2)-1] //删除最后一个空格

	return string(s2)
}

func reverseWords01(s string) (res string) {
	s = " " + s + " "
	l, r := len(s)-1, len(s)-1
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] == ' ' {
			l, r = i, l
			if r > l+1 {
				res = res + s[l+1:r] + " "
			}
		}
	}
	return res[:len(res)-1]
}
