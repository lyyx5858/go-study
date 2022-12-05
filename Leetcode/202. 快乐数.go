package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print(isHappy(1))

}

func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] { //因为m的键是n，值是bool，因此可以直接判断
		n, m[n] = getSum(n), true
	}

	// for ; n != 1 && !m[n]; n, m[n] = getSum(n), true { }

	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func isHappyMy(n int) bool {

	s := strconv.Itoa(n)
	l := len(s)
	c := make([]int, l)
	y := 0

	for i := 0; i <= l-1; i++ {
		c[i], _ = strconv.Atoi(string(s[i]))
		y = y + c[i]*c[i]
	}
	if _, ok := map1[y]; ok {
		return false
	}
	map1[y] = n
	if y == 1 || isHappy(y) {
		return true
	}

	return false
}

var map1 = make(map[int]int)
