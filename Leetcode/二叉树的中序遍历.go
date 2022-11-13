package main

import (
	"fmt"
)

//请将Leetcode格式的二叉树构建.go这个文件 包含进来，否则编译出错

func main() {
	v := []int{1, -200, 2, -200, 3, -200}
	//v := []int{1, 2, 3, 4, 5, 6, 7}

	t := buildTree(v)
	v1 := inorderTraversal(t)
	fmt.Println(v1)
}

var res []int

func inorderTraversal(root *TreeNode) []int {
	res = []int{}
	run(root) //最后root
	return res
}

func run(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		run(root.Left) //先左
	}
	res = append(res, root.Val)

	if root.Right != nil {
		run(root.Right) //后右
	}
	return
}
