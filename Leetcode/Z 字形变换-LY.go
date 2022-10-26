package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	s := "PAYPALISHIRING"
	r := convert(s, 1)
	fmt.Println(r)
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	result := make([]string, 2)
	l := len(s)
	t := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		t[i] = make([]string, numRows)
	}

	jstart := -(numRows - 2) //-1
	jmax := numRows - 1      //2

	j := jstart
	x := 0

	for i := 0; i <= l-1; i++ {
		fmt.Printf("%v ", x)
		t[x] = append(t[x], s[i:i+1])
		x = jmax - int(math.Abs(float64(j)))

		if x == 0 && j < jmax {
			j++
			continue
		}
		if j < jmax {
			j++
		} else {
			j = jstart
		}

	}
	for i := 0; i < numRows; i++ {
		result = append(result, t[i]...)
	}
	fmt.Println("\n", result)
	return strings.Join(result, "")

}
