package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	s := " -42"
	//s := "   +0 123"
	fmt.Println(myAtoi(s))

}

func myAtoi(s string) int {
	ans := 0
	t := make([]string, len(s))
	f := 1
	var c1 int32 = -1
	for _, c := range s {
		if c == int32(' ') && (c1 == int32(' ') || c1 == -1) {
			c1 = c
			continue
		} else if c == int32('-') && (c1 == int32(' ') || c1 == -1) {
			f = -1
			c1 = c
		} else if c == int32('+') && (c1 == int32(' ') || c1 == -1) {
			f = 1
			c1 = c
			continue
		} else if c >= 48 && c <= 57 {
			c1 = c
			t = append(t, string(c))
		} else {
			break
		}
	}

	//fmt.Println(t)
	s1 := strings.Join(t, "")
	//fmt.Println(s1)
	ans, _ = strconv.Atoi(s1)
	ans = ans * f
	if ans <= math.MinInt32 { //此处是按题目要求的判断
		ans = math.MinInt32
	} else if ans >= math.MaxInt32 {
		ans = math.MaxInt32
	}
	return ans
}
