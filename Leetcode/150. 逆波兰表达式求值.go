package main

import (
	"fmt"
	"strconv"
)

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	//tokens := []string{"2"}
	//tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	v := evalRPN(tokens)
	fmt.Println(v)

}

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	v := 0
	if len(tokens) == 1 {
		v, _ = strconv.Atoi(tokens[0])
		return v
	}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			stack = append(stack, tokens[i]) //入栈
		} else {
			a, _ := strconv.Atoi(stack[len(stack)-2])
			b, _ := strconv.Atoi(stack[len(stack)-1])
			if tokens[i] == "+" {
				v = a + b
			}
			if tokens[i] == "-" {
				v = a - b
			}
			if tokens[i] == "*" {
				v = a * b
			}
			if tokens[i] == "/" {
				v = a / b
			}
			stack = stack[0 : len(stack)-2]        //出栈
			stack = append(stack, strconv.Itoa(v)) //入栈

		}
	}

	return v
}
