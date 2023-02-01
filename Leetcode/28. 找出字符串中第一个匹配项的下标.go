package main

import "fmt"

func main() {
	haystack := "aabaabaafa"
	needle := "baa"
	fmt.Println(strStr(haystack, needle))
}

// 方法二: 前缀表无减一或者右移

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	j := 0
	next := make([]int, n)
	getNext(next, needle)
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1] // 回退到j的前一位
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

// getNext 构造前缀表next
// params:
//		  next 前缀表数组
//		  s 模式串
func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); { //j是左指针，i是右指针
		if s[i] == s[j] { //如果相同，同时前进，更新next
			next[i] = j + 1
			j++
			i++
		} else if j == 0 { //如果不同，j已经回退至最左，则next=0
			next[i] = 0
			i++
		} else {
			j = next[j-1] //j回退至next数组j-1里所定义的的位置
		}
	}
}
