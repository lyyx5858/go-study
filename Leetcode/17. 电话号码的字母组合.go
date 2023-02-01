package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
}

var (
	m    []string
	path []byte
	res  []string
)

func letterCombinations(digits string) []string {
	m = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	path, res = make([]byte, 0), make([]string, 0)
	if digits == "" {
		return res
	}
	dfs(digits, 0)
	return res
}
func dfs(digits string, start int) {
	if len(path) == len(digits) { //终止条件，字符串长度等于digits的长度
		tmp := string(path)
		res = append(res, tmp)
		return
	}
	d := int(digits[start] - '0') // 将index指向的数字转为int（确定下一个数字）
	str := m[d-2]                 // 取数字对应的字符集（注意和map中的对应）
	for j := 0; j < len(str); j++ {
		//当start=0,此处for循环是循环：str=a,b,c
		//当start=1,此处for循环是循环：str=d,e,f,
		//也可以这样理解：start=0是外循环，start=1是内循环

		path = append(path, str[j])
		dfs(digits, start+1)
		path = path[:len(path)-1]
	}
}
