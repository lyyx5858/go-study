package main

import "fmt"

func main() {
	s := "cbaabcabc"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	s1 := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if len(s1) > 0 && s[i] == s1[len(s1)-1] {
			s1 = s1[0 : len(s1)-1]
		} else {
			s1 = append(s1, s[i])
		}
	}
	return string(s1)
}
