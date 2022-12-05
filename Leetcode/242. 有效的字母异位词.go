package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(isAnagram0(s, t))

}

func isAnagram0(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m1 := make(map[rune]int) //rune=int32
	for _, char := range s {
		m1[char]++
	}
	fmt.Println(m1)

	for _, char := range t {
		m1[char]--
		if m1[char] < 0 {
			return false
		}
	}

	return true
}
