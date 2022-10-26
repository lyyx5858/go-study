package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	x := -1230
	y := reverse(x)
	fmt.Println(y)

}

func reverse(x int) int {
	f := 1
	if x <= 0 {
		f = -1
		x = f * x
	}
	a := strconv.Itoa(x)
	l := len(a)
	fmt.Println(a)
	b := make([]string, l)
	for i, c := range a {
		b[l-i-1] = string(c)
	}

	d := strings.Join(b, "")
	e, _ := strconv.Atoi(d)
	r := e * f
	fmt.Println(r)

	r1 := math.Pow(2, 31) - 1
	r2 := -math.Pow(2, 31)
	fmt.Println(r1, r2)
	if float64(r) >= r1 || float64(r) <= r2 {
		r = 0
	}
	return r
}
