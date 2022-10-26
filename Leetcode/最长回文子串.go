package main

import "fmt"

func main() {
	s := "abcbaeads"
	r := longestPalindrome(s)
	fmt.Println(r)

}

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1] //初始化结果(最小的回文就是单个字符)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 单个字符一定是回文，所以标记为true，dp[1][1]表示开始坐标为1，结束坐标为1的字串，其实就是单个字符。
	}
	for length := 2; length <= len(s); length++ { //长度固定，不断移动起点
		for start := 0; start < len(s)-length+1; start++ { //start从0开始，长度固定，不断移动起点
			end := start + length - 1
			if s[start] != s[end] { //首尾不同则不可能为回文
				continue
			} else if length < 3 {
				dp[start][end] = true //双字符：如果两个字符相同，则也是回文
			} else {
				dp[start][end] = dp[start+1][end-1] //等号左边的的真假决定于等号右边
				//所有长度大于3的字符串，如果首尾字符相同时，只需要判断内部字串是否是回文即可。
				//内部是，则是，内部不是，则不是。所以此处的真假就等于内部的真假。

			}

			if dp[start][end] && (end-start+1) > len(result) { //当判断是回文后，判断是否是最长的。result在不断被刷新
				result = s[start : end+1]
			}

			// s := "abcbaeads"
		}
	}
	return result
}
