package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "efe"

	fmt.Println(partition(s))
}

var (
	res  [][]string
	path []string
)

func partition(s string) [][]string {
	path = make([]string, 0)
	res = make([][]string, 0)
	dfs(s, s)
	return res
}

func dfs(s string, orgin string) { //orgin是用来判断是否path组合是否已经完全包含原始字串s

	if len(path) > 0 && strings.Join(path, "") == orgin { //当分割完成，开始进行判断，Join函数是将path重新组合进来
		tmp := make([]string, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for j := 0; j < len(s); j++ { //外循环，也叫宽度循环

		if !isHuiwen(s[:j+1]) { //剪枝，只有回文才能进path
			continue
		}
		/* 以s=abcd为例，循环过程如下：
			第0层：a, ab, abc, abcd
		    第1层：当第0层是a时，1层是：b，bc，bcd，当第0层是ab时，1层是：c，cd
		     以此类推，当path已经完全等于s时，结束循环。
		     详情见：md文档

		*/
		path = append(path, s[:j+1])
		dfs(s[j+1:], orgin) //内循环，也叫深度循环
		path = path[:len(path)-1]
	}
}

func isHuiwen(s string) bool { //用来判断s是否是回文
	if len(s) == 0 {
		return false
	}

	l, r := 0, len(s)-1 //定义左、右指针
	for l <= r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
