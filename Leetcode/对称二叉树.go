package main

import (
	"fmt"
)

//请将Leetcode格式的二叉树构建.go这个文件 包含进来，否则编译出错

func main() {

	v := []int{5, 4, 1, -200, 1, -200, 4, 2, -200, 2, -200}
	//v := []int{1, 2, 2, 3, 4, 4, 3}
	//v := []int{1, 2, 2, 2, -200, 2, -200}

	t := buildTree(v)
	r := isSymmetric(t)
	fmt.Println(r)

}

func isSymmetric(root *TreeNode) bool {
	left := inorderTraversal_modify(root.Left)
	right := inorderTraversal_modify(root.Right)
	fmt.Println(left)
	fmt.Println(right)

	if len(left) == 0 && len(right) == 0 {
		return true
	}

	if len(left) != len(right) || len(left) == 0 || len(right) == 0 {
		return false
	}

	if root.Left.Val != root.Right.Val {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] == right[len(left)-i-1] {
			continue
		} else {
			return false
		}

	}

	return true
}

func inorderTraversal_modify(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode) //定义一个函数变量:inorder
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)

		if node.Left == nil && node.Right == nil {
			res = append(res, -200)
		}
		res = append(res, node.Val)

		inorder(node.Right)
		if node.Left == nil && node.Right == nil {
			res = append(res, -200)
		}
	}
	inorder(root)
	return
}
