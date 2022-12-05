package main

import "fmt"

func main() {

	s := " "
	fmt.Println(replaceSpace(s))

}

func replaceSpace(s string) string {
	s1 := []byte(s)
	sp := []byte("%20")
	for i := 0; i <= len(s1)-1; i++ {
		if s1[i] == ' ' {
			s1 = append(s1[:i], append(sp, s1[i+1:]...)...)
		}
	}

	return string(s1)
}
