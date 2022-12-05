package main

import (
	"fmt"
)

func main() {
	s := "abaaacd"
	p := "aba*cd"

	//s := "aa"
	//p := "aa*"

	b := isMatch(s, p)
	fmt.Println(b)
}

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1) //dp[i][j]表示s的前i个字符与p的前j个字符是否match
	// match就是dp的值是ture，不match就是false
	for i, _ := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true //当s与p都是空串是，是true

	for j := 1; j <= n; j++ { //除了左上角，dp的第一列肯定是0，
		// 但dp的第一行是就可能出现1的。比如：p="a*b*c*"
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	printdp(s, p, dp)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1] //当s与p各自的最后一个字符相同时,或者p的最后一个字串是.时
				//则s与p是否相同，取决于各自前一位的字串是否相同
			} else if p[j-1] == '*' {
				if s[i-1] != p[j-2] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2] //由于*可以消除前面字符（重复0次），则当前 dp值 就取决于p往前数两位的字符（一个是*号本身，一个是被*消除的字符 ） 与s是否相同
				} else {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				}
			}
		}

	}

	printdp(s, p, dp)

	return dp[m][n]

}

// printdp是用来打印dp这个二维数组用于测试
func printdp(s, p string, dp [][]bool) {
	m, n := len(s), len(p)
	for i := 0; i <= n; i++ {
		if i == 0 {
			fmt.Printf("            ")
		} else {
			fmt.Printf("%c ", p[i-1])
		}
	}

	fmt.Printf("\n")
	for i := 0; i <= m; i++ {
		if i == 0 {
			fmt.Printf("s[%d], %s  %d \n", i, " ", booltoint(dp[i]))
		} else {
			fmt.Printf("s[%d], %c  %d \n", i, s[i-1], booltoint(dp[i]))

		}
	}

	fmt.Println("===========================")
}

//将ture转换为1，将false转换为0
func booltoint(b1 []bool) []int {

	m := len(b1)
	x := make([]int, m)
	for i := 0; i < m; i++ {
		if b1[i] == true {
			x[i] = 1
		} else {
			x[i] = 0
		}
	}
	return x
}
