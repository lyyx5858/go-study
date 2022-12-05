package main

import (
	"fmt"
)

func main() {

	//v := []int{1}
	//v := []int{1, 2}
	v := []int{4, 2, 7, 1, 3, 6, 9}
	//v := []int{2, 3, -200, 1}
	//v := []int{2, 1, 4, 3, -200, 5}
	//v := []int{3, 9, 20, -200, -200, 15, 7}
	t := buildTree(v)
	t = invertTree(t)
	fmt.Println(invertTree(t))
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	tmp := &TreeNode{}
	tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
