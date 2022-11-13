package main

import "fmt"

//请将Leetcode格式的二叉树构建.go这个文件 包含进来，否则编译出错

func main() {

	//v := []int{5, 4, 1, -200, 1, -200, 4, 2, -200, 2, -200}
	//v := []int{1, 2, 2, 3, 4, 4, 3}
	v := []int{1, 2, 2, 2, -200, 2, -200}

	t := buildTree(v)
	r := isSymmetric(t)
	fmt.Println(r)

}

func isSymmetric(root *TreeNode) bool {
	f := check(root, root)
	return f
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left) {
		return true
	}

	return false
}
