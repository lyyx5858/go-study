package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "PAYPALISHIRING"
	r := convert(s, 3)
	fmt.Println(r)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	n := 2*numRows - 2
	//假如numRows=3, 从下可以看出，循环是以2*3-2=4为单位
	// 0 1 2 1 / 0 1 2 1 /0 1 2 1 / 0 1....

	for i, char := range s {
		x := i % n                        //%为求余数
		rows[min(x, n-x)] += string(char) //x与n-x之间小者为构建新字符顺序
	}
	// rows[0]="PAHN"
	// rows[1]="APLSIIG"
	// rows[2]="YIR"

	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//作者：zoffer
//链接：https://leetcode.cn/problems/zigzag-conversion/solution/ji-jian-jie-fa-by-ijzqardmbd/
