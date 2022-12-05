package main

import "fmt"

func main() {
	s := "abcdefg"
	k := 8
	fmt.Println(reverseStr(s, k))

}

func reverseStr(s string, k int) string {
	s1 := []byte(s)

	for i := 0; i < len(s1)-1; i = i + 2*k {
		l := i             //left 指针
		r := l + k - 1     //right指针
		if r > len(s1)-1 { //当右指针大于串长时，将右指针置为串长
			r = len(s1) - 1
		}
		for l < r && r <= len(s1)-1 {
			s1[l], s1[r] = s1[r], s1[l]
			l++
			r--
		}
	}

	return string(s1)
}
