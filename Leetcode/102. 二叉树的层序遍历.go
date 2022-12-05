package main

import "fmt"

func main() {
	v := []int{3, 9, 20, -200, -200, 15, 7}
	//v := []int{1, 2, 3, 4, 5, 6, 7}
	//v := []int{}

	t := buildTree(v)
	printTree(t)
	v1 := levelOrder(t)
	fmt.Println(v1)
}

var res [][]int
var i int

func levelOrder(root *TreeNode) [][]int {
	res = make([][]int, 0)

	run(root, 0) //最后root

	return res
}

func run(root *TreeNode, i int) {
	if root == nil {
		return
	}

	if len(res) < i+1 {
		res = append(res, nil)

	}

	res[i] = append(res[i], root.Val)

	if root.Left != nil {
		run(root.Left, i+1) //先左
	}
	if root.Right != nil {
		run(root.Right, i+1) //后右
	}
	return
}
