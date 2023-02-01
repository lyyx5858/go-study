package main

import "fmt"

func main() {
	s := "{()}]"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'} //右括号是键，左括号是值
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i]) //压栈所有左括号
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1] //出栈
		} else {
			return false //右括号多了，或者没匹配上
		}
	}
	return len(stack) == 0 //判断是否左括号多了
}
