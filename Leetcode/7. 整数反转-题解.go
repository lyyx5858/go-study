package main

import (
	"fmt"
	"math"
)

func main() {
	x := 12345
	y := reverse(x)
	fmt.Println(y)
}

func reverse(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || ans > math.MaxInt32/10 { //此处是按题目要求的判断
			return 0
		}
		ans = ans*10 + x%10
		x /= 10
	}
	return ans
}

//作者：dfzhou6
//链接：https://leetcode.cn/problems/reverse-integer/solution/jian-dan-yi-dong-go-by-dfzhou6-nhf3/
