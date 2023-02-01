package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s1 := "25525511135"
	s2 := "0000"
	s3 := "101023"
	fmt.Println(restoreIpAddresses(s1))
	fmt.Println(restoreIpAddresses(s2))
	fmt.Println(restoreIpAddresses(s3))
}

var (
	res  []string
	path []string
)

func restoreIpAddresses(s string) []string {
	path = make([]string, 0)
	res = make([]string, 0)
	dfs(s, s)
	return res
}

func dfs(s string, orgin string) {

	if len(path) == 4 && strings.Join(path, "") == orgin { //当分割完成，开始进行判断，Join函数是将path重新组合进来
		str := strings.Join(path, ".")
		res = append(res, str)
		return
	}

	for j := 0; j < len(s); j++ { //外循环，也叫宽度循环
		if v, _ := strconv.Atoi(s[:j+1]); v > 255 || len(path) > 4 { //值大于255的，大于4段的剪树
			continue
		} else if len(s[:j+1]) > 1 && s[0] == 48 { //开头是0的二位数以上的剪枝，48代表“0”
			continue
		}

		path = append(path, s[:j+1])
		dfs(s[j+1:], orgin) //内循环，也叫深度循环, 前j个元素分析完了，开始分析剩下的
		path = path[:len(path)-1]
	}
}
