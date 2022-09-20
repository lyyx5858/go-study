package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "bcaabcbb"

	fmt.Println(lengthOfLongestSubstring(s))

}

func lengthOfLongestSubstring(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right] //窗口字串
	for right = 0; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1 //左窗口移到：查到重复字符的位置(index)的右边1位
		}
		s1 = s[left : right+1] //右窗口向右移1位
		if len(s1) > Length {  //Length记录窗口字串的最大值
			Length = len(s1)
		}
	}

	return Length
}
